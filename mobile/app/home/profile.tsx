import { Alert, ScrollView, StyleSheet, View } from "react-native";
import { router, useNavigation, usePathname } from "expo-router";
import { useEffect, useState } from "react";
import { useGnoNativeContext } from "@gnolang/gnonative";
import { avatarTxAndRedirectToSign, broadcastTxCommit, clearLinking, logedOut, reloadAvatar, selectAccount, selectQueryParamsTxJsonSigned, useAppDispatch, useAppSelector } from "@gno/redux";
import Button from "@gno/components/button";
import Layout from "@gno/components/layout";
import { LoadingModal } from "@gno/components/loading";
import { AccountBalance } from "@gno/components/settings";
import Text from "@gno/components/text";
import { useSearch } from "@gno/hooks/use-search";
import { useNotificationContext } from "@gno/provider/notification-provider";
import AvatarPicker from "@gno/components/avatar/avatar-picker";
import { ProgressViewModal } from "@gno/components/view/progress";
import { compressImage } from '@gno/utils/file-utils';
import { useUserCache } from "@gno/hooks/use-user-cache";

export default function Page() {
  const [loading, setLoading] = useState(false);
  const [modalVisible, setModalVisible] = useState(false);
  const [chainID, setChainID] = useState("");
  const [remote, setRemote] = useState("");
  const [followersCount, setFollowersCount] = useState({ n_followers: 0, n_following: 0 });
  const pathName = usePathname();

  const account = useAppSelector(selectAccount);
  const { gnonative } = useGnoNativeContext();
  const search = useSearch();
  const navigation = useNavigation();
  const dispatch = useAppDispatch();
  const push = useNotificationContext();
  const txJsonSigned = useAppSelector(selectQueryParamsTxJsonSigned);

  const userCache = useUserCache();

  useEffect(() => {
    const unsubscribe = navigation.addListener("focus", async () => {
      try {
        fetchAccountData();
      } catch (error: unknown | Error) {
        console.log(error);
      }
    });
    return unsubscribe;
  }, [navigation]);

  const onPressNotification = async () => {
    if (!account) {
      throw new Error("No active account");
    }

    setLoading(true);
    try {
      const address_bech32 = await gnonative.addressToBech32(account.address);
      await push.registerDevice(address_bech32);
      Alert.alert("Push notification", "You have successfully registered for push notification!");
    } catch (error) {
      console.log("Error on onPressNotification", JSON.stringify(error));
    } finally {
      setLoading(false);
    }
  };

  const fetchAccountData = async () => {
    const chainId = await gnonative.getChainID();
    const remote = await gnonative.getRemote();
    setChainID(chainId);
    setRemote(remote);

    if (!account) {
      throw new Error("No active account");
    }

    try {
      const followersCount = await search.GetJsonFollowersCount(account.bech32);
      console.log("followersCount", followersCount);
      setFollowersCount(followersCount);

      console.log("remote: %s chainId %s " + remote, chainId);
    } catch (error: unknown | Error) {
      console.log(error);
    }
  };

  const onRemoveAccount = async () => {
    router.navigate({ pathname: "account/remove" });
  };

  const onPressLogout = async () => {
    dispatch(clearLinking());
    dispatch(logedOut());
  };

  const onAvatarChanged = async (imagePath: string, mimeType?: string) => {
    const imageCompressed = await compressImage(imagePath)
    if (!imageCompressed || !mimeType || !imageCompressed.base64) {
      console.log("Error compressing image or missing data");
      return;
    }

    if (!account) throw new Error("No account found");
    await dispatch(avatarTxAndRedirectToSign({ mimeType, base64: imageCompressed.base64, callerAddressBech32: account.bech32, callbackPath: pathName })).unwrap();
  }

  useEffect(() => {

    (async () => {
      if (txJsonSigned) {

        console.log("txJsonSigned: ", txJsonSigned);

        const signedTx = decodeURIComponent(txJsonSigned as string)
        try {
          await dispatch(broadcastTxCommit(signedTx)).unwrap();
        } catch (error) {
          console.error("on broadcastTxCommit", error);
        }

        dispatch(clearLinking());
        userCache.invalidateCache();
        
        setTimeout(() => {
          console.log("reloading avatar");
          dispatch(reloadAvatar());
        }, 500);
      }
    })();

  }, [txJsonSigned]);

  return (
    <>
      <Layout.Container>
        <Layout.Body>
          <ScrollView >
            <View style={{ paddingBottom: 20 }}>
              <AvatarPicker onChanged={onAvatarChanged} />
            </View>
            <>
              <AccountBalance activeAccount={account} />
              <Text.Subheadline>Chain ID:</Text.Subheadline>
              <Text.Body>{chainID}</Text.Body>
              <Text.Subheadline>Remote:</Text.Subheadline>
              <Text.Body>{remote}</Text.Body>
              <Text.Subheadline>Followers:</Text.Subheadline>
              <Text.Body>{followersCount.n_followers}</Text.Body>
              <Text.Subheadline>Following:</Text.Subheadline>
              <Text.Body>{followersCount.n_following}</Text.Body>
              <View></View>
            </>
            <Layout.Footer>
              <ProgressViewModal visible={modalVisible} onRequestClose={() => setModalVisible(false)} />
              <Button.TouchableOpacity title="Logs" onPress={() => setModalVisible(true)} variant="primary" />
              <Button.TouchableOpacity
                title="Register to the notification service"
                onPress={onPressNotification}
                variant="primary"
              />
              <Button.TouchableOpacity title="Logout" onPress={onPressLogout} style={styles.logout} variant="primary-red" />
              <Button.TouchableOpacity
                title="Remove Account"
                onPress={onRemoveAccount}
                style={styles.logout}
                variant="primary-red"
              />
            </Layout.Footer>
          </ScrollView>
        </Layout.Body>
      </Layout.Container>
      <LoadingModal visible={loading} />
    </>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: "center",
    padding: 24,
  },
  main: {
    flex: 1,
    justifyContent: "center",
    maxWidth: 960,
    marginHorizontal: "auto",
  },
  logout: {
    color: "#007AFF",
    marginTop: 10,
    fontWeight: "bold",
    textAlign: "center",
  },
});
