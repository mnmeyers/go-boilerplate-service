package http

import "github.com/go-chi/render"

// RenderJSON is an alias of method to render JSON for easy mocking in tests
var RenderJSON = render.JSON