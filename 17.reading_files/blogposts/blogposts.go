package blogposts

import (
	"errors"
	"io"
	"io/fs"
)

type Post struct {
	Title string
}

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("fail")
}

func NewPostsFromFS(filesystem fs.FS) ([]Post, error) {

	dir, err := fs.ReadDir(filesystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post

	for _, f := range dir {
		post, err := getPost(filesystem, f)
		if err != nil {
			return nil, err // needs clarif, should we fail if one fails?
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(filesystem fs.FS, f fs.DirEntry) (Post, error) {
	postFile, err := filesystem.Open(f.Name())
	if err != nil {
		return Post{}, err
	}

	defer postFile.Close()

	postData, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, err
	}

	post := Post{Title: string(postData)[7:]}

	return post, nil
}
