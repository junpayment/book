package models

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// BookRepository ...
type BookRepository struct {
	Source Source
}

// NewBookRepository ...
func NewBookRepository(s Source) *BookRepository {
	return &BookRepository{s}
}

// BookSource ... corresponds to Source interface
type BookSource struct {
	URL   string
	Query string
}

// NewBookSource ...
func NewBookSource() *BookSource {
	return &BookSource{
		URL: "https://www.googleapis.com/books/v1/volumes",
	}
}

// Body ...
func (b *BookSource) Body() ([]byte, error) {
	url := fmt.Sprintf("%s?q=%s", b.URL, b.Query)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	cli := http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
