package social

import (
	"std"
	"strconv"

	"gno.land/p/demo/avl"
	"gno.land/p/demo/avlhelpers"
	"gno.land/p/demo/ufmt"
	"gno.land/r/sys/users"
)

type UserAndPostID struct {
	UserPostAddr std.Address
	PostID       PostID
}

// Post a message to the caller's main user posts.
// The caller must already be registered with /r/gnoland/users/v1 Register.
// Return the "thread ID" of the new post.
// (This is similar to boards.CreateThread, but no message title)
func PostMessage(_ realm, body string) PostID {
	caller := std.OriginCaller()
	name := usernameOf(caller)
	if name == "" {
		panic("please register")
	}

	userPosts := getOrCreateUserPosts(caller, name)
	thread := userPosts.AddThread(body)
	return thread.id
}

// Post a reply to the user posts of userPostsAddr where threadid is the ID returned by
// the original call to PostMessage. If postid == threadid then create another top-level
// post for the threadid, otherwise post a reply to the postid "sub reply".
// The caller must already be registered with /r/gnoland/users/v1 Register.
// Return the new post ID.
// (This is similar to boards.CreateReply.)
func PostReply(_ realm, userPostsAddr std.Address, threadid, postid PostID, body string) PostID {
	caller := std.OriginCaller()
	if usernameOf(caller) == "" {
		panic("please register")
	}
	userPosts := getUserPosts(userPostsAddr)
	if userPosts == nil {
		panic("posts for userPostsAddr do not exist")
	}
	thread := userPosts.GetThread(threadid)
	if thread == nil {
		panic("threadid in user posts does not exist")
	}
	if postid == threadid {
		reply := thread.AddReply(caller, body)
		return reply.id
	} else {
		post := thread.GetReply(postid)
		if post == nil {
			panic("postid does not exist")
		}
		reply := post.AddReply(caller, body)
		return reply.id
	}
}

// Repost the message from the user posts of userPostsAddr where threadid is the ID returned by
// the original call to PostMessage. This must be a top-level thread (not a reply).
// Return the new post ID.
// (This is similar to boards.CreateRepost.)
func RepostThread(_ realm, userPostsAddr std.Address, threadid PostID, comment string) PostID {
	caller := std.OriginCaller()
	if userPostsAddr == caller {
		panic("Cannot repost a user's own message")
	}

	name := usernameOf(caller)
	if name == "" {
		panic("please register")
	}
	dstUserPosts := getOrCreateUserPosts(caller, name)

	userPosts := getUserPosts(userPostsAddr)
	if userPosts == nil {
		panic("posts for userPostsAddr do not exist")
	}
	thread := userPosts.GetThread(threadid)
	if thread == nil {
		panic("threadid in user posts does not exist")
	}
	repost := thread.AddRepostTo(caller, comment, dstUserPosts)
	return repost.id
}

// For each address/PostID in addrAndIDs, get the thread post. The Post ID must be
// for a a top-level thread (not a reply; to get reply posts, use GetThreadPosts).
// If the Post ID is not found, set the result for that Post ID to {}.
// The response is a JSON string.
func GetJsonTopPostsByID(addrAndIDs []UserAndPostID) string {
	json := "[ "
	for _, addrAndID := range addrAndIDs {
		if len(json) > 2 {
			json += ",\n  "
		}

		userPosts := getUserPosts(addrAndID.UserPostAddr)
		if userPosts == nil {
			json += "{}"
			continue
		}

		post := userPosts.GetThread(PostID(addrAndID.PostID))
		if post == nil {
			json += "{}"
			continue
		}

		postJson, err := post.MarshalJSON()
		if err != nil {
			panic("can't get post JSON")
		}
		json += string(postJson)
	}
	json += "]"

	return json
}

// Get posts in a thread for a user. A thread is the sequence of posts without replies.
// While each post has an an arbitrary id, it also has an index within the thread starting from 0.
// Limit the response to posts from startIndex up to (not including) endIndex within the thread.
// If you just want the total count, set startIndex and endIndex to 0 and see the response "n_threads".
// If threadID is 0 then return the user's top-level posts. (Like render args "user".)
// If threadID is X and replyID is 0, then return the posts (without replies) in that thread. (Like render args "user/2".)
// If threadID is X and replyID is Y, then return the posts in the thread starting with replyID. (Like render args "user/2/5".)
// The response includes reposts by this user (only if threadID is 0), but not messages of other
// users that are being followed. (See GetHomePosts.) The response is a JSON string.
func GetThreadPosts(userPostsAddr std.Address, threadID int, replyID int, startIndex int, endIndex int) string {
	userPosts := getUserPosts(userPostsAddr)
	if userPosts == nil {
		panic("posts for userPostsAddr do not exist")
	}

	if threadID == 0 {
		return getPosts(userPosts.threads, startIndex, endIndex)
	}

	thread := userPosts.GetThread(PostID(threadID))
	if thread == nil {
		panic(ufmt.Sprintf("thread does not exist with id %d", threadID))
	}

	if replyID == 0 {
		return getPosts(thread.replies, startIndex, endIndex)
	} else {
		reply := thread.GetReply(PostID(replyID))
		if reply == nil {
			panic(ufmt.Sprintf("reply does not exist with id %d in thread with id %d", replyID, threadID))
		}

		return getPosts(reply.replies, startIndex, endIndex)
	}
}

// Update the home posts by scanning all posts from all followed users and adding the
// followed posts since the last call to RefreshHomePosts (or since started following the user).
// Return the new count of home posts. The result is something like "(12 int)".
func RefreshHomePosts(_ realm, userPostsAddr std.Address) int {
	userPosts := getUserPosts(userPostsAddr)
	if userPosts == nil {
		panic("posts for userPostsAddr do not exist")
	}
	userPosts.refreshHomePosts()

	return userPosts.homePosts.Size()
}

// Get the number of posts which GetHomePosts or GetJsonHomePosts will return.
// The result is something like "(12 int)".
// This returns the current count of the home posts (without need to pay gas). To include the
// latest followed posts, call RefreshHomePosts.
func GetHomePostsCount(userPostsAddr std.Address) int {
	return GetHomePosts(userPostsAddr).Size()
}

// Get home posts for a user, which are the user's top-level posts plus all posts of all
// users being followed.
// The response is a map of postID -> *Post. The avl.Tree sorts by the post ID which is
// unique for every post and increases in time.
// If you just want the total count, use GetHomePostsCount.
// This returns the current state of the home posts (without need to pay gas). To include the
// latest followed posts, call RefreshHomePosts.
func GetHomePosts(userPostsAddr std.Address) *avl.Tree {
	userPosts := getUserPosts(userPostsAddr)
	if userPosts == nil {
		panic("posts for userPostsAddr do not exist")
	}
	return &userPosts.homePosts
}

// Get home posts for a user (using GetHomePosts), which are the user's top-level posts plus all
// posts of all users being followed.
// Limit the response to posts from startIndex up to (not including) endIndex within the home posts.
// If you just want the total count, use GetHomePostsCount.
// The response is a JSON string.
// This returns the current state of the home posts (without need to pay gas). To include the
// latest posts, call RefreshHomePosts.
func GetJsonHomePosts(userPostsAddr std.Address, startIndex int, endIndex int) string {
	allPosts := GetHomePosts(userPostsAddr)
	postsJson := ""
	for i := startIndex; i < endIndex && i < allPosts.Size(); i++ {
		_, postI := allPosts.GetByIndex(i)
		if postsJson != "" {
			postsJson += ",\n  "
		}

		postJson, err := postI.(*Post).MarshalJSON()
		if err != nil {
			panic("can't get post JSON")
		}
		postsJson += ufmt.Sprintf("{\"index\": %d, \"post\": %s}", int(i), string(postJson))
	}

	return ufmt.Sprintf("{\"n_posts\": %d, \"posts\": [\n  %s]}", allPosts.Size(), postsJson)
}

// Update the caller to follow the user with followedAddr. See UserPosts.Follow.
func Follow(_ realm, followedAddr std.Address) PostID {
	caller := std.OriginCaller()
	name := usernameOf(caller)
	if name == "" {
		panic("please register")
	}

	if followedAddr == caller {
		panic("you can't follow yourself")
	}

	// A user can follow someone before doing any posts, so create the UserPosts if needed.
	userPosts := getOrCreateUserPosts(caller, name)
	return userPosts.Follow(followedAddr)
}

// Update the caller to unfollow the user with followedAddr. See UserPosts.Unfollow.
func Unfollow(_ realm, followedAddr std.Address) {
	caller := std.OriginCaller()
	name := usernameOf(caller)
	if name == "" {
		panic("please register")
	}

	userPosts := getUserPosts(caller)
	if userPosts == nil {
		// We don't expect this, but just do nothing.
		return
	}

	userPosts.Unfollow(followedAddr)
}

// Add the reaction by the caller to the post of userPostsAddr, where threadid is the ID
// returned by the original call to PostMessage. If postid == threadid then add the reaction
// to a top-level post for the threadid, otherwise add the reaction to the postid "sub reply".
// (This function's arguments are similar to PostReply.)
// The caller must already be registered with /r/gnoland/users/v1 Register.
// Return a boolean indicating whether the userAddr was added. See Post.AddReaction.
func AddReaction(_ realm, userPostsAddr std.Address, threadid, postid PostID, reaction Reaction) bool {
	caller := std.OriginCaller()
	if usernameOf(caller) == "" {
		panic("please register")
	}
	userPosts := getUserPosts(userPostsAddr)
	if userPosts == nil {
		panic("posts for userPostsAddr do not exist")
	}
	thread := userPosts.GetThread(threadid)
	if thread == nil {
		panic("threadid in user posts does not exist")
	}
	if postid == threadid {
		return thread.AddReaction(caller, reaction)
	} else {
		post := thread.GetReply(postid)
		if post == nil {
			panic("postid does not exist")
		}
		return post.AddReaction(caller, reaction)
	}
}

// Remove the reaction by the caller to the post of userPostsAddr, where threadid is the ID
// returned by the original call to PostMessage. If postid == threadid then remove the reaction
// from a top-level post for the threadid, otherwise remove the reaction from the postid "sub reply".
// (This function's arguments are similar to PostReply.)
// The caller must already be registered with /r/gnoland/users/v1 Register.
// Return a boolean indicating whether the userAddr was removed. See Post.RemoveReaction.
func RemoveReaction(_ realm, userPostsAddr std.Address, threadid, postid PostID, reaction Reaction) bool {
	caller := std.OriginCaller()
	if usernameOf(caller) == "" {
		panic("please register")
	}
	userPosts := getUserPosts(userPostsAddr)
	if userPosts == nil {
		panic("posts for userPostsAddr do not exist")
	}
	thread := userPosts.GetThread(threadid)
	if thread == nil {
		panic("threadid in user posts does not exist")
	}
	if postid == threadid {
		return thread.RemoveReaction(caller, reaction)
	} else {
		post := thread.GetReply(postid)
		if post == nil {
			panic("postid does not exist")
		}
		return post.RemoveReaction(caller, reaction)
	}
}

// Call users.ResolveAddress and return the result as JSON, or "" if not found.
// (This is a temporary utility until gno.land supports returning structured data directly.)
func GetJsonUserByAddress(addr std.Address) string {
	user := users.ResolveAddress(addr)
	if user == nil {
		return ""
	}

	return marshalJsonUser(user)
}

// Call users.ResolveName and return the result as JSON, or "" if not found.
// (This is a temporary utility until gno.land supports returning structured data directly.)
func GetJsonUserByName(name string) string {
	user, _ := users.ResolveName(name)
	if user == nil {
		return ""
	}

	return marshalJsonUser(user)
}

// Get the UserPosts info for the user with the given address, including
// url, n_threads, n_followers and n_following. If the user address is not
// found, return "". The name of this function has "Info" because it just returns
// the number of items, not the items themselves. To get the items, see
// GetJsonFollowers, etc.
// The response is a JSON string.
func GetJsonUserPostsInfo(address std.Address) string {
	userPosts := getUserPosts(address)
	if userPosts == nil {
		return ""
	}

	json, err := userPosts.MarshalJSON()
	if err != nil {
		panic("can't get UserPosts JSON")
	}

	return string(json)
}

// Get the UserPosts for the user with the given address, and return
// the list of followers. If the user address is not found, return "".
// Limit the response to entries from startIndex up to (not including) endIndex.
// The response is a JSON string.
func GetJsonFollowers(address std.Address, startIndex int, endIndex int) string {
	userPosts := getUserPosts(address)
	if userPosts == nil {
		return ""
	}

	json := ufmt.Sprintf("{\"n_followers\": %d, \"followers\": [\n  ", userPosts.followers.Size())
	for i := startIndex; i < endIndex && i < userPosts.followers.Size(); i++ {
		addr, _ := userPosts.followers.GetByIndex(i)

		if i > startIndex {
			json += ",\n  "
		}
		json += ufmt.Sprintf(`{"address": "%s"}`, addr)
	}
	json += "]}"

	return json
}

// Get the UserPosts for the user with the given address, and return
// the list of other users that this user is following.
// If the user address is not found, return "".
// Limit the response to entries from startIndex up to (not including) endIndex.
// The response is a JSON string.
func GetJsonFollowing(address std.Address, startIndex int, endIndex int) string {
	userPosts := getUserPosts(address)
	if userPosts == nil {
		return ""
	}

	json := ufmt.Sprintf("{\"n_following\": %d, \"following\": [\n  ", userPosts.following.Size())
	for i := startIndex; i < endIndex && i < userPosts.following.Size(); i++ {
		addr, infoI := userPosts.following.GetByIndex(i)

		if i > startIndex {
			json += ",\n  "
		}
		startedAt, err := infoI.(*FollowingInfo).startedFollowingAt.MarshalJSON()
		if err != nil {
			panic("can't get startedFollowingAt JSON")
		}
		json += ufmt.Sprintf(`{"address": "%s", "started_following_at": %s}`,
			addr, string(startedAt))
	}
	json += "]}"

	return json
}

// Get a list of user names starting from the given prefix. Limit the
// number of results to maxResults.
func ListUsersByPrefix(prefix string, maxResults int) []string {
	return avlhelpers.ListByteStringKeysByPrefix(&gUserAddressByName, prefix, maxResults)
}

// Get a list of user names starting from the given prefix. Limit the
// number of results to maxResults.
// The response is a JSON string.
func ListJsonUsersByPrefix(prefix string, maxResults int) string {
	names := ListUsersByPrefix(prefix, maxResults)

	json := "["
	for i, name := range names {
		if i > 0 {
			json += ", "
		}
		json += strconv.Quote(name)
	}
	json += "]"
	return json
}
