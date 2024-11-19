package blogrender_test

import (
	"bytes"
	"testing"

	blogrender "github.com/mircealaurentiu/go_with_tests/18.templating"
)

func TestRender(t *testing.T) {

	aPost := blogrender.Post{
		Title:       "hello world",
		Body:        "this is a post",
		Description: "this is a description",
		Tags:        []string{"go", "tdd"},
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrender.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>hello world</h1>
<p>this is a description</p>
Tags: <ul><li>go</li><li>tdd</li></ul>`

		if got != want {
			t.Errorf("\ngot  %s \nwant %s", got, want)
		}
	})

}
