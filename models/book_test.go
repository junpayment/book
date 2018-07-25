package models

import (
  "testing"
)

type Mock struct {
  echo []byte
}

func (m *Mock) Body() ([]byte, error) {
  return m.echo, nil
}

func TestGetBookItems(t *testing.T) {
  m := &Mock{}
  m.echo = []byte("{\"kind\":\"books#volumes\",\"totalItems\":1566,\"items\":[{\"kind\":\"books#volume\",\"id\":\"uPQvAAAACAAJ\",\"etag\":\"hGQtYQLXzZI\",\"selfLink\":\"https://www.googleapis.com/books/v1/volumes/uPQvAAAACAAJ\",\"volumeInfo\":{\"title\":\"片想い\",\"authors\":[\"東野圭吾\"],\"publishedDate\":\"2004-08\",\"description\":\"十年ぶりに再会した美月は、男の姿をしていた。彼女から、殺人を告白された哲朗は、美月の親友である妻とともに、彼女をかくまうが...。十年という歳月は、かつての仲間たちを、そして自分を、変えてしまったのだろうか。過ぎ去った青春の日々を裏切るまいとする仲間たちを描いた、傑作長篇ミステリー。\",\"industryIdentifiers\":[{\"type\":\"ISBN_10\",\"identifier\":\"4167110091\"},{\"type\":\"ISBN_13\",\"identifier\":\"9784167110093\"}],\"readingModes\":{\"text\":false,\"image\":false},\"pageCount\":622,\"printType\":\"BOOK\",\"categories\":[\"Detective and mystery stories\"],\"maturityRating\":\"NOT_MATURE\",\"allowAnonLogging\":false,\"contentVersion\":\"preview-1.0.0\",\"imageLinks\":{\"smallThumbnail\":\"http://books.google.com/books/content?id=uPQvAAAACAAJ&printsec=frontcover&img=1&zoom=5&source=gbs_api\",\"thumbnail\":\"http://books.google.com/books/content?id=uPQvAAAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api\"},\"language\":\"ja\",\"previewLink\":\"http://books.google.co.jp/books?id=uPQvAAAACAAJ&dq=%E6%9D%B1%E9%87%8E%E5%9C%AD%E5%90%BE&hl=&cd=1&source=gbs_api\",\"infoLink\":\"http://books.google.co.jp/books?id=uPQvAAAACAAJ&dq=%E6%9D%B1%E9%87%8E%E5%9C%AD%E5%90%BE&hl=&source=gbs_api\",\"canonicalVolumeLink\":\"https://books.google.com/books/about/%E7%89%87%E6%83%B3%E3%81%84.html?hl=&id=uPQvAAAACAAJ\"},\"saleInfo\":{\"country\":\"JP\",\"saleability\":\"NOT_FOR_SALE\",\"isEbook\":false},\"accessInfo\":{\"country\":\"JP\",\"viewability\":\"NO_PAGES\",\"embeddable\":false,\"publicDomain\":false,\"textToSpeechPermission\":\"ALLOWED\",\"epub\":{\"isAvailable\":false},\"pdf\":{\"isAvailable\":false},\"webReaderLink\":\"http://play.google.com/books/reader?id=uPQvAAAACAAJ&hl=&printsec=frontcover&source=gbs_api\",\"accessViewStatus\":\"NONE\",\"quoteSharingAllowed\":false},\"searchInfo\":{\"textSnippet\":\"十年ぶりに再会した美月は、男の姿をしていた。彼女から、殺人を告白された哲朗は、美月の親友である妻とともに、彼女をかくまうが.. ...\"}}]}")

  r := NewBookRepository(m)

  _, err := r.GetBookItems()
  if err != nil {
    t.Fatal(err)
  }
}

func TestBookSource_Body(t *testing.T) {
  bs := &BookSource{
    "https://www.googleapis.com/books/v1/volumes",
    "東野圭吾",
  }

  _, err := bs.Body()
  if err != nil {
    t.Fatal(err)
  }
}

func TestBookRepository_GetBookItems(t *testing.T) {
  bs := &BookSource{
    "https://www.googleapis.com/books/v1/volumes",
    "東野圭吾",
  }

  r := NewBookRepository(bs)
  _, err := r.GetBookItems()

  if err != nil {
    t.Fatal(err)
  }
}
