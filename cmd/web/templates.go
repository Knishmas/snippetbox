package main

import "snippetbox.andresmacias.dev/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
