import { ActivityIndicator, Platform, StyleSheet, View } from "react-native";
import React, { useEffect, useRef, useState } from "react";
import { useNavigation, useRouter } from "expo-router";
import { useFeed } from "@gno/hooks/use-feed";
import Alert from "@gno/components/alert";
import Layout from "@gno/components/layout";
import { Post } from "@gno/types";
import { FlatList } from "react-native-gesture-handler";
import useScrollToTop from "@gno/components/utils/useScrollToTopWithOffset";
import EmptyFeedList from "@gno/components/feed/empty-feed-list";
import { Tweet } from "@gno/components/feed/tweet";
import Button from "@gno/components/button";

export default function Page() {
  const pageSize = 9;
  const [startIndex, setStartIndex] = useState(0);
  const [endIndex, setEndIndex] = useState(pageSize);
  const [limit, setLimit] = useState(pageSize + 1);
  const [data, setData] = useState<Post[]>([]);
  const [error, setError] = useState<unknown | Error | undefined>(undefined);
  const [isLoading, setIsLoading] = useState(false);
  const [isEndReached, setIsEndReached] = useState(false);

  const router = useRouter();
  const feed = useFeed();
  const navigation = useNavigation();
  const ref = useRef<FlatList>(null);

  useScrollToTop(ref, Platform.select({ ios: -150, default: 0 }));

  useEffect(() => {
    const unsubscribe = navigation.addListener("focus", async () => {
      setError(undefined);
      fetchData();
    });
    return unsubscribe;
  }, [navigation]);

  const handleEndReached = async () => {
    console.log("end reached", isEndReached);
    if (!isEndReached) {
      setIsEndReached(true);
      fetchData();
    }
  };

  const fetchData = async () => {
    setIsLoading(true);
    try {
      console.log("fetching data from %d to %d", startIndex, endIndex);
      const result = await feed.fetchFeed(startIndex, endIndex);
      setLimit(result.n_posts);
      setStartIndex(endIndex);
      setEndIndex(endIndex + pageSize);

      //join the data
      console.log("current data length", data.length);
      console.log("new data length", result.data.length);
      setData([...data, ...result.data]);
      console.log("startIndex: %s, limit: %s", startIndex, limit);
      setIsEndReached(startIndex >= limit);
    } catch (error: unknown | Error | any) {
      // TODO: Check if this is the correct error message to handle and if it's the correct way to handle it
      // https://github.com/gnolang/gnonative/issues/117
      if (error.message === "[unknown] invoke bridge method error: unknown: posts for userPostsAddr do not exist") {
        setData([]);
        return;
      }
      setError(error);
      console.log(error);
    } finally {
      setIsLoading(false);
    }
  };

  const onPressPost = () => {
    router.push("/post");
  };

  const renderFooter = () => {
    if (!isLoading) return null;
    return (
      <View style={styles.footer}>
        <ActivityIndicator size="large" color="#0000ff" />
      </View>
    );
  };

  if (error) {
    return (
      <Layout.Container>
        <Layout.Body>
          <Alert severity="error" message="Error while fetching posts, please, check the logs." />
        </Layout.Body>
      </Layout.Container>
    );
  }

  return (
    <Layout.Container>
      <View style={styles.container}>
        <FlatList
          ref={ref}
          scrollToOverflowEnabled
          data={data}
          ListFooterComponent={renderFooter}
          ListEmptyComponent={<EmptyFeedList />}
          keyExtractor={(item) => `${item.id}`}
          contentContainerStyle={styles.flatListContent}
          renderItem={({ item }) => <Tweet item={item} />}
          onEndReached={handleEndReached}
          onEndReachedThreshold={0.1}
        />
        <Button.TouchableOpacity title="Post" onPress={onPressPost} style={styles.post} variant="primary" />
      </View>
    </Layout.Container>
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