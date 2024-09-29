package main

import (
	"ddd-template/app"
	"fmt"
	"net/http"
)

func main() {
	appSetup, err := app.Setup()
	if err != nil {
		panic(err)
	}
	if err := http.ListenAndServe(fmt.Sprintf(":%d", appSetup.Config.Port), appSetup.Handler); err != nil {
		panic(err)
	}
}
