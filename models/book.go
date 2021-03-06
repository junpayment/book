package models

import (
	"encoding/json"
)

// Book ...
type Book struct {
	Kind       string  `json:"kind"`
	TotalItems int     `json:"totalItems"`
	Items      []*Item `json:"items"`
}

// Item ...
type Item struct {
	Kind       string      `json:"kind"`
	ID         string      `json:"id"`
	Etag       string      `json:"etag"`
	SelfLink   string      `json:"selfLink"`
	VolumeInfo *VolumeInfo `json:"volumeInfo"`
	SaleInfo   *SaleInfo   `json:"saleInfo"`
	AccessInfo *AccessInfo `json:"accessInfo"`
	SearchInfo *SearchInfo `json:"searchInfo"`
}

// VolumeInfo ...
type VolumeInfo struct {
	Title               string                `json:"title"`
	Authors             []string              `json:"authors"`
	PublishedDate       string                `json:"publishedDate"`
	Description         string                `json:"description"`
	IndustryIdentifiers []*IndustryIdentifier `json:"industryIdentifiers"`
	ReadingModes        *ReadingModes         `json:"readingModes"`
	PageCount           int                   `json:"pageCount"`
	PrintType           string                `json:"printType"`
	Categories          []string              `json:"categories"`
	MaturityRating      string                `json:"maturityRating"`
	AllowAnonLogging    bool                  `json:"allowAnonLogging"`
	ContentVersion      string                `json:"contentVersion"`
	ImageLinks          *ImageLinks           `json:"imageLinks"`
	Language            string                `json:"language"`
	PreviewLink         string                `json:"previewLink"`
	InfoLink            string                `json:"infoLink"`
	CanonicalVolumeLink string                `json:"canonicalVolumeLink"`
}

// IndustryIdentifier ...
type IndustryIdentifier struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
}

// ReadingModes ...
type ReadingModes struct {
	Text  bool `json:"text"`
	Image bool `json:"image"`
}

// ImageLinks ...
type ImageLinks struct {
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail      string `json:"thumbnail"`
}

// SaleInfo ...
type SaleInfo struct {
	Country     string `json:"country"`
	Saleability string `json:"saleability"`
	IsEbook     bool   `json:"isEbook"`
}

// AccessInfo ...
type AccessInfo struct {
	Country                string `json:"country"`
	Viewability            string `json:"viewability"`
	Embeddable             bool   `json:"embeddable"`
	PublicDomain           bool   `json:"publicDomain"`
	TextToSpeechPermission string `json:"textToSpeechPermission"`
	Epub                   *Epub  `json:"epub"`
	Pdf                    *Pdf   `json:"pdf"`
	WebReaderLink          string `json:"webReaderLink"`
	AccessViewStatus       string `json:"accessViewStatus"`
	QuoteSharingAllowed    bool   `json:"quoteSharingAllowed"`
}

// Epub ...
type Epub struct {
	IsAvailable bool `json:"isAvailable"`
}

// Pdf ...
type Pdf struct {
	IsAvailable bool `json:"isAvailable"`
}

// SearchInfo ...
type SearchInfo struct {
	TextSnippet string `json:"textSnippet"`
}

// GetBookItems ...
func (r *BookRepository) GetBookItems() (*Book, error) {
	body, err := r.Source.Body()
	if err != nil {
		return nil, err
	}

	book := &Book{}

	err = json.Unmarshal(body, book)
	if err != nil {
		return nil, err
	}

	return book, nil
}
