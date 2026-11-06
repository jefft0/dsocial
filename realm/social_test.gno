package social

import (
	"testing"

	"gno.land/p/nt/bptree/v0"
)

func TestPostLookups(t *testing.T) {
	addr := address("g1jg8mtutu9khhfwc4nxmuhcpftf0pajdhfvsqf5")
	if getUserPosts(addr) != nil {
		t.Fatal("unexpected posts for an unknown user")
	}
	userPosts := getOrCreateUserPosts(addr, "alice")
	if getUserPosts(addr) != userPosts {
		t.Fatal("stored user posts were not found")
	}
	if userPosts.GetThread(1) != nil {
		t.Fatal("unexpected thread in an empty user feed")
	}
	thread := userPosts.addThread("Hello")
	if userPosts.GetThread(thread.id) != thread {
		t.Fatal("stored thread was not found")
	}
	if thread.GetReply(thread.id+1) != nil {
		t.Fatal("unexpected reply in an empty thread")
	}
	reply := thread.addReply(addr, "Reply")
	if thread.GetReply(reply.id) != reply {
		t.Fatal("stored reply was not found")
	}
}

func TestParseHashtags(t *testing.T) {
	tests := []struct {
		body string
		want []string
	}{
		{"", nil},
		{"no tags here", nil},
		{"#gno", []string{"gno"}},
		{"Hello #Gno and #dSocial!", []string{"gno", "dsocial"}},
		{"#a #A #a", []string{"a"}},
		{"(#paren) [#bracket], #comma,", []string{"paren", "bracket", "comma"}},
		{"line1\n#next_line", []string{"next_line"}},
		{"#café #日本 #v2", []string{"café", "日本", "v2"}},
		{"#foo#bar", []string{"foo"}},
		{"#don't #Bob’s '#quoted'", []string{"don", "bob", "quoted"}},
		{"issue #1 and #_ and # alone and #", nil},
		{"a#b &#39; /#anchor ##double", nil},
		{"# Heading\n## Sub", nil},
	}
	for _, tc := range tests {
		got := parseHashtags(tc.body)
		if len(got) != len(tc.want) {
			t.Fatalf("parseHashtags(%q) = %v, want %v", tc.body, got, tc.want)
		}
		for i := range got {
			if got[i] != tc.want[i] {
				t.Fatalf("parseHashtags(%q) = %v, want %v", tc.body, got, tc.want)
			}
		}
	}

	userPosts := newUserPosts("/r/berty/social:alice", address("alice"))
	thread := userPosts.addThread("Hello #gno")
	reply := thread.addReply(address("bob"), "Reply #dsocial")
	if len(thread.hashtags) != 1 || thread.hashtags[0] != "gno" {
		t.Fatalf("thread hashtags = %v, want [gno]", thread.hashtags)
	}
	if len(reply.hashtags) != 1 || reply.hashtags[0] != "dsocial" {
		t.Fatalf("reply hashtags = %v, want [dsocial]", reply.hashtags)
	}
}

func TestPostsByHashtag(t *testing.T) {
	alice := newUserPosts("/r/berty/social:alice", address("alice"))
	bob := newUserPosts("/r/berty/social:bob", address("bob"))
	thread := alice.addThread("Thread #Topic #other")
	reply := thread.addReply(address("bob"), "Reply #topic")
	repost := thread.addRepostTo(address("bob"), "Repost #topic", bob)
	alice.addThread("No tags")

	checkPosts := func(hashtag string, want []*Post) {
		postsI := gPostsByHashtag.Get(hashtag)
		if postsI == nil {
			t.Fatalf("no posts for hashtag %q", hashtag)
		}
		posts := postsI.(*bptree.BPTree)
		if posts.Size() != len(want) {
			t.Fatalf("hashtag %q has %d posts, want %d", hashtag, posts.Size(), len(want))
		}
		for i, post := range want {
			key, postI := posts.GetByIndex(i)
			if key != postIDKey(post.id) || postI.(*Post) != post {
				t.Fatalf("hashtag %q post %d is %s, want %s", hashtag, i, key, postIDKey(post.id))
			}
		}
	}
	checkPosts("topic", []*Post{thread, reply, repost})
	checkPosts("other", []*Post{thread})
	if gPostsByHashtag.Get("Topic") != nil || gPostsByHashtag.Get("missing") != nil {
		t.Fatal("unexpected posts for a hashtag")
	}
}

func TestReactionLookups(t *testing.T) {
	userPosts := newUserPosts("/r/berty/social:alice", address("alice"))
	thread := userPosts.addThread("Hello")
	if thread.GetReactionCount(Gnod) != 0 {
		t.Fatal("unexpected reactions on a new post")
	}
	if !thread.addReaction(address("bob"), Gnod) {
		t.Fatal("first reaction was not added")
	}
	if thread.addReaction(address("bob"), Gnod) {
		t.Fatal("duplicate reaction was added")
	}
	if thread.GetReactionCount(Gnod) != 1 {
		t.Fatal("stored reaction was not counted")
	}
	if !thread.removeReaction(address("bob"), Gnod) || thread.GetReactionCount(Gnod) != 0 {
		t.Fatal("reaction was not removed")
	}
}
