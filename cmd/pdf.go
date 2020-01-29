package main

import (
	"net/http"

	"github.com/SuddenGunter/pandaren/browser"
)

func PdfHandlerFunc(b *browser.Browser) func(http.ResponseWriter, *http.Request) {
	timeout := 10
	return func(http.ResponseWriter, *http.Request) {
		// todo bind body
		err := b.PrintAsPdf()

	}
}
