package models

// Source ...
type Source interface {
	Body() ([]byte, error)
}
