package blogposts_test

import (
	"testing"
	"testing/fstest"

	"github.com/mircealaurentiu/go_with_tests/17.reading_files/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello_world.md":   {Data: []byte("hi")},
		"hello_world_2.md": {Data: []byte("hola")},
	}

	posts := blogposts.NewPostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("\ngot  %d posts \nwant %d posts", len(posts), len(fs))
	}
}
