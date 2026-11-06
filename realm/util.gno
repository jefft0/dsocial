package social

import (
	"strconv"
	"strings"
	"unicode"

	"gno.land/p/nt/bptree/v0"
	"gno.land/p/nt/ufmt/v0"
	"gno.land/r/sys/users"
)

//----------------------------------------
// private utility methods

// Get the UserPosts for the user.
func getUserPosts(userAddr address) *UserPosts {
	userPosts := gUserPostsByAddress.Get(userAddr.String())
	if userPosts == nil {
		return nil
	}

	return userPosts.(*UserPosts)
}

// Get the UserPosts for the userAddr. If not found, add a new UserPosts to
// gUserPostsByAddress and update gUserAddressByName with the username.
// (The caller usually has already called usernameOf to get the username, but if
// it is "" then this will get it.)
func getOrCreateUserPosts(userAddr address, username string) *UserPosts {
	userPosts := getUserPosts(userAddr)
	if userPosts != nil {
		return userPosts
	}

	if username == "" {
		username = usernameOf(userAddr)
	}
	if username == "" {
		username = userAddr.String()
	}

	userPosts = newUserPosts(gRealmPath+":"+username, userAddr)
	gUserPostsByAddress.Set(userAddr.String(), userPosts)
	gUserAddressByName.Set(username, userAddr)

	return userPosts
}

func padZero(u64 uint64, length int) string {
	str := strconv.Itoa(int(u64))
	if len(str) >= length {
		return str
	} else {
		return strings.Repeat("0", length-len(str)) + str
	}
}

func postIDKey(pid PostID) string {
	return padZero(uint64(pid), 10)
}

func reactionKey(reaction Reaction) string {
	return strconv.Itoa(int(reaction))
}

// If reactions has an value for the given reaction, then return it.
// Otherwise, add the reaction key to reactions, set the value to an empty bptree.BPTree and return it.
func getOrCreateReactionValue(reactions *bptree.BPTree, reaction Reaction) *bptree.BPTree {
	key := reactionKey(reaction)
	valueI := reactions.Get(key)
	if valueI != nil {
		return valueI.(*bptree.BPTree)
	} else {
		value := bptree.NewBPTree32()
		reactions.Set(key, value)
		return value
	}
}

// listByteStringKeysByPrefix returns up to maxResults keys from tree that start
// with the given prefix, treating keys as byte strings (not Unicode runes).
// Inlined from gno.land/p/jefft0/avlhelpers which is not yet deployed.
func listByteStringKeysByPrefix(tree bptree.ITree, prefix string, maxResults int) []string {
	result := []string{}
	end := ""
	n := len(prefix)
	for n > 0 {
		if ascii := int(prefix[n-1]); ascii < 0xff {
			end = prefix[0:n-1] + string(ascii+1)
			break
		}
		n--
	}
	tree.Iterate(prefix, end, func(key string, value any) bool {
		result = append(result, key)
		return len(result) >= maxResults
	})
	return result
}

func indentBody(indent string, body string) string {
	lines := strings.Split(body, "\n")
	res := ""
	for i, line := range lines {
		if i > 0 {
			res += "\n"
		}
		res += indent + line
	}
	return res
}

// NOTE: length must be greater than 3.
func summaryOf(str string, length int) string {
	lines := strings.SplitN(str, "\n", 2)
	line := lines[0]
	if len(line) > length {
		line = line[:(length-3)] + "..."
	} else if len(lines) > 1 {
		// len(line) <= 80
		line = line + "..."
	}
	return line
}

// Return the unique hashtags in body, lowercase and without the '#', in order of first
// appearance. A hashtag is a '#' followed by letters (with combining marks), digits or '_', with
// at least one letter (so "#1" is not a hashtag). The '#' must be at the start of the body or
// follow a character which is not one of these or '&', '/' or '#' (to skip "a#b", "&#39;",
// "/#anchor" and "##a").
// Return nil if there are no hashtags.
func parseHashtags(body string) []string {
	var hashtags []string
	start := -1 // The byte index after the '#' of the hashtag being scanned, or -1.
	prev := rune(0)
	for i, r := range body {
		if start >= 0 && !isHashtagRune(r) {
			hashtags = addHashtag(hashtags, body[start:i])
			start = -1
		}
		if r == '#' && !isHashtagRune(prev) && !strings.ContainsRune("&/#", prev) {
			start = i + 1
		}
		prev = r
	}
	if start >= 0 {
		hashtags = addHashtag(hashtags, body[start:])
	}
	return hashtags
}

func isHashtagRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsMark(r) || unicode.IsDigit(r) || r == '_'
}

// If tag has a letter, then append it in lowercase to hashtags if not already there.
func addHashtag(hashtags []string, tag string) []string {
	if strings.IndexFunc(tag, unicode.IsLetter) < 0 {
		return hashtags
	}
	tag = strings.ToLower(tag)
	for _, hashtag := range hashtags {
		if hashtag == tag {
			return hashtags
		}
	}
	return append(hashtags, tag)
}

// Add post to gPostsByHashtag under each of its hashtags.
func addPostToHashtags(post *Post) {
	for _, hashtag := range post.hashtags {
		postsI := gPostsByHashtag.Get(hashtag)
		var posts *bptree.BPTree
		if postsI == nil {
			posts = bptree.NewBPTree32()
			gPostsByHashtag.Set(hashtag, posts)
		} else {
			posts = postsI.(*bptree.BPTree)
		}
		posts.Set(postIDKey(post.id), post)
	}
}

func displayAddressMD(addr address) string {
	user := users.ResolveAddress(addr)
	if user == nil {
		return "[" + addr.String() + "](/r/gnoland/users/v1:" + addr.String() + ")"
	} else {
		return "[@" + user.Name() + "](" + gRealmPath + ":" + user.Name() + ")"
	}
}

func usernameOf(addr address) string {
	user := users.ResolveAddress(addr)
	if user == nil {
		return ""
	} else {
		return user.Name()
	}
}

// Return the User info as a JSON string.
// (This is a temporary utility until gno.land supports returning structured data directly.)
func marshalJsonUser(user *users.UserData) string {
	return ufmt.Sprintf(
		"{\"address\": \"%s\", \"name\": \"%s\", \"deleted\": %t}",
		user.Addr().String(), user.Name(), user.IsDeleted())
}
