package blogrender_test

import (
	"bytes"
	"testing"

	blogrender "github.com/mircealaurentiu/go_with_tests/18.templating"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {

	var (
		aPost = blogrender.Post{
			Title:       "hello world",
			Body:        "this is a post",
			Description: "this is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrender.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

}
