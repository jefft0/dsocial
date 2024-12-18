import { ActivityIndicator, FlatList, Platform, StyleSheet, View, Alert as RNAlert, SafeAreaView } from "react-native";
import React, { useEffect, useRef, useState } from "react";
import { useNavigation, usePathname, useRouter } from "expo-router";
import { useFeed } from "@gno/hooks/use-feed";
import Layout from "@gno/components/layout";
import useScrollToTop from "@gno/components/utils/useScrollToTopWithOffset";
import Button from "@gno/components/button";
import { Post } from "@gno/types";
import { gnodTxAndRedirectToSign, selectAccount, setPostToReply, useAppDispatch, useAppSelector } from "@gno/redux";
import Alert from "@gno/components/alert";
import { FeedView } from "@gno/components/view";

export default function Page() {
  const [totalPosts, setTotalPosts] = useState(0);
  const [error, setError] = useState<unknown | Error | undefined>(undefined);
  const [isLoading, setIsLoading] = useState(true);

  const router = useRouter();
  const navigation = useNavigation();
  const feed = useFeed();
  const ref = useRef<FlatList>(null);
  const dispatch = useAppDispatch();

  const account = useAppSelector(selectAccount);
  const pathName = usePathname();

  useScrollToTop(ref, Platform.select({ ios: -150, default: 0 }));

  useEffect(() => {
    const unsubscribe = navigation.addListener("focus", async () => {
      if (!account) return;
      setError(undefined);
      setIsLoading(true);
      try {
        const total = await feed.fetchCount(account.bech32);
        setTotalPosts(total);
      } catch (error) {
        RNAlert.alert("Error while fetching posts.", " " + error);
        console.error(error);
      } finally {
        setIsLoading(false);
      }
    });
    return unsubscribe;
  }, [navigation]);

  const onPressPost = () => {
    router.navigate({ pathname: "/post" });
  };

  const onPress = async (p: Post) => {
    await dispatch(setPostToReply(p));
    router.navigate({ pathname: "/post/[post_id]" });
  };

  const onGnod = async (post: Post) => {
    if (!account) throw new Error("No active account");
    dispatch(gnodTxAndRedirectToSign({ post, callerAddressBech32: account.bech32, callbackPath: pathName })).unwrap();
  };

  if (isLoading)
    return (
      <Layout.Container>
        <Layout.Body>
          <ActivityIndicator size="large" color="#0000ff" />
        </Layout.Body>
      </Layout.Container>
    );

  if (error)
    return (
      <Layout.Container>
        <Layout.Body>
          <Alert severity="error" message="Error while fetching posts, please, check the logs." />
        </Layout.Body>
      </Layout.Container>
    );

  if (!account) {
    return (
      <Layout.Container>
        <Layout.Body>
          <Alert severity="error" message="No user found." />
        </Layout.Body>
      </Layout.Container>
    );
  }

  return (
    <SafeAreaView style={{ flex: 1, paddingTop: Platform.select({ ios: 0, default: 20 }) }}>
      <View style={styles.container}>
        <FeedView totalPosts={totalPosts} onPress={onPress} onGnod={onGnod} bech32={account.bech32} type="userFeed" />
        <Button.TouchableOpacity title="Post" onPress={onPressPost} style={styles.post} variant="primary" />
      </View>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: "flex-end",
    alignItems: "stretch",
  },
  flatListContent: {
    paddingBottom: 60, // Adjust the value to ensure it's above the app menu
  },
  footer: {
    paddingVertical: 20,
    alignItems: "center",
  },
  post: {
    position: "absolute",
    width: 60,
    height: 60,
    bottom: 40,
    right: 20,
  },
});
