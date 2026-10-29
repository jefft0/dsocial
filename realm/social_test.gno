package social

import "testing"

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
