package main

import (
	"context"
	"github.com/dennyback/interview/application"
)

func main() {
	app := application.Application{context.Background()}
	_ = app.RunAndServeApplication("service-name")
}
