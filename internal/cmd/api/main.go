package main

import (
	"github.com/AAAAAlexeyyyyy/s3_service/internal/run/app"
)

func main() {
	l, err := app.Run()
	if err != nil {
		l.Error("Error running run: ", err)
	}
}
