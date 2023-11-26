package main

import (
	"github.com/RB-PRO/sibintek/internal/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		return
	}

	err = a.Run()
	if err != nil {
		return
	}
}
