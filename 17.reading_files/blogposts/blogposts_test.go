package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/mircealaurentiu/go_with_tests/17.reading_files/blogposts"
)

func TestNewBlogPosts(t *testing.T) {

	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tag1, tag2
---
Hello World
1`
		secondBody = `Title: Post 2
Description: Description 2
Tags: tag3, tag4
---
Hello
WORLD
2
!`
	)

	fs := fstest.MapFS{
		"hello_world.md":   {Data: []byte(firstBody)},
		"hello_world_2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("\ngot  %d posts \nwant %d posts", len(posts), len(fs))
	}

	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tag1", "tag2"},
		Body: `Hello World
1`,
	})
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot  %+v \nwant %+v", got, want)
	}
}
