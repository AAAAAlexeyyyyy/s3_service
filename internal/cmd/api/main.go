package main

import (
	"github.com/AAAAAlexeyyyyy/s3_service/internal/run/app"
)

func main() {
	l, err := app.Start()
	if err != nil {
		l.Errorw("Error start application", "due to error", err)
	}
}
