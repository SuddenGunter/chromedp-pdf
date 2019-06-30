package main

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"io/ioutil"
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

	// save the screenshot to disk
	if err = ioutil.WriteFile("/store/test.pdf", buf, 0666); err != nil {
		log.Fatal(err)
	}
	files, err := ioutil.ReadDir("/store")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(cap(files))
	log.Print("I am 2")
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
