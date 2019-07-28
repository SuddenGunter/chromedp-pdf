package main

import (
	"context"
	"fmt"
	"os"

	"github.com/SuddenGunter/pandaren/pkg/pdf"
	"github.com/caarlos0/env"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
)

type config struct {
	PdfStorePath string `env:"PANDAREN_PDF_PATH,required"`
}

func main() {

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var buf []byte
	err := chromedp.Run(ctx, navigate(`https://www.google.com/`, `#main`, &buf))
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	chromedp.NewContext(context.Background(), &chromedp.ContextOption{})

	store, err := getDefaultStore(cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	err = writeFile(store, buf)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}

func getDefaultStore(cfg config) (pdf.Store, error) {
	config := &pdf.FileSystemStoreConfig{
		Path:        cfg.PdfStorePath,
		Permissions: 0666,
	}

	fs, err := pdf.NewFileSystemStore(config)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create default file store")
	}

	return fs, nil
}

func writeFile(store pdf.Store, bytes []byte) error {
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
		makePdf(res),
	}
}

func makePdf(pdfbuf *[]byte) chromedp.Action {
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
