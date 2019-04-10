package server

import (
	"github.com/drhodes/golorem"
)

// Service define interface
type Service interface {
	Word(min, max int) string
	Sentence(min, max int) string
	Paragraph(min, max int) string
}

// LoremService ..
type LoremService struct{}

// Word ..
func (LoremService) Word(min, max int) string {
	return lorem.Word(min, max)
}

// Sentence ..
func (LoremService) Sentence(min, max int) string {
	return lorem.Sentence(min, max)
}

// Paragraph ..
func (LoremService) Paragraph(min, max int) string {
	return lorem.Paragraph(min, max)
}
