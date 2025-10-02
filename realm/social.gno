package social

import (
	"gno.land/p/nt/avl"
)

var (
	gUserPostsByAddress avl.Tree // user's address -> *UserPosts
	gUserAddressByName  avl.Tree // user's username -> address
	postsCtr            uint64   // increments Post.id globally
)
