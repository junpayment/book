package models

type Source interface {
  Body() ([]byte, error)
}
