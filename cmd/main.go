package main

import (
	"context"
	"fmt"
	"github.com/SuddenGunter/pandaren/pkg/pdfstore"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"log"
)

func main() {
	log.Print("I am 1")
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var buf []byte
	err := chromedp.Run(ctx, navigate(`https://www.google.com/`, `#main`, &buf))
	if err != nil {
		log.Fatal(err)
	}

	store := getDefaultStore()
	err = writeFile(store, buf)
	if err != nil {
		log.Fatal(err)
	}
}

func getDefaultStore() pdfstore.PdfStore {
	config := &pdfstore.FileStoreConfig{
		Path:              "/store",
		Permissions:       0666,
		FileNameGenerator: pdfstore.DefaultFileNameGenerator(),
	}

	fs, err := pdfstore.NewFileStore(config)
	if err != nil {
		log.Fatalln(fmt.Errorf("unable to create default file store: %v", err))
	}

	return fs
}

func writeFile(store pdfstore.PdfStore, bytes []byte) error {
	_, err := store.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func navigate(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(sel, chromedp.ByID),
		pdf(res),
	}
}

func pdf(pdfbuf *[]byte) chromedp.Action {
	if pdfbuf == nil {
		panic("pdfbuf cannot be nil")
	}

	return chromedp.ActionFunc(func(ctx context.Context) error {
		// take page screenshot
		buf, err := page.PrintToPDF().Do(ctx)
		if err != nil {
			return err
		}
		*pdfbuf = buf
		return nil
	})
}
