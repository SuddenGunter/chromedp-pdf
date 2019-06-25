package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
)


func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var buf []byte
	err := chromedp.Run(ctx, screenshot(`https://www.google.com/`, `#main`, &buf))
	if err != nil {
		log.Fatal(err)
	}

	// save the screenshot to disk
	if err = ioutil.WriteFile("screenshot.png", buf, 0644); err != nil {
		log.Fatal(err)
	}
}

func screenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(sel, chromedp.ByID),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible, chromedp.ByID),
	}
}