package blogposts

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("fail")
}

const (
	titleSeparation       = "Title: "
	descriptionSeparation = "Description: "
	tagsSeparation        = "Tags: "
)

func NewPostsFromFS(filesystem fs.FS) ([]Post, error) {

	dir, err := fs.ReadDir(filesystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post

	for _, f := range dir {
		post, err := getPost(filesystem, f.Name())
		if err != nil {
			return nil, err // needs clarif, should we fail if one fails?
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(filesystem fs.FS, filename string) (Post, error) {
	postFile, err := filesystem.Open(filename)
	if err != nil {
		return Post{}, err
	}

	defer postFile.Close()

	return newPost(postFile)
}

func newPost(postFile io.Reader) (Post, error) {

	scanner := bufio.NewScanner(postFile)

	readLine := func(prefix string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), prefix)
	}

	return Post{
		Title:       readLine(titleSeparation),
		Description: readLine(descriptionSeparation),
		Tags:        strings.Split(readLine(tagsSeparation), ", "),
		Body:        readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // for the separation line, "---"
	buf := bytes.Buffer{}

	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	return strings.TrimSuffix(buf.String(), "\n") // remove the last new line from Fprintln
}
