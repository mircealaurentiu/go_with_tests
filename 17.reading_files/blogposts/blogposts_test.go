package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/mircealaurentiu/go_with_tests/17.reading_files/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello_world.md":   {Data: []byte("Title: Post 1")},
		"hello_world_2.md": {Data: []byte("Title: Post 2")},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("\ngot  %d posts \nwant %d posts", len(posts), len(fs))
	}

	got := posts[0]
	want := blogposts.Post{Title: "Post 1"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot  %+v, want %+v", got, want)
	}
}
