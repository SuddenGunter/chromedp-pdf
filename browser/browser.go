package browser

import (
	"context"
	"time"

	"github.com/chromedp/cdproto/page"

	"github.com/chromedp/chromedp"
)

// Browser is responsible for dedicated browser instance handling
type Browser struct {
	chromedp context.Context
	// Close browser
	Close context.CancelFunc
}

func New() *Browser {
	ctx, cancel := chromedp.NewContext(context.Background())
	return &Browser{chromedp: ctx, Close: cancel}
}

// NavigateAndWaitLoad navigates to website and waits until element with selector is loaded
// where:
// 	timeout is the operation timeout in seconds
//	url is url to navigate, example http://localhost
//  selector is valid query selector, example "#searchbar"
func (b *Browser) PrintAsPdf(timeout time.Duration, url string, selector string) *[]byte {
	c, _ := context.WithTimeout(b.chromedp, timeout)
	tab, cancel := chromedp.NewContext(c)

	defer cancel()

	var buf []byte
	if err := chromedp.Run(tab,
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			buf, _, err = page.PrintToPDF().
				WithDisplayHeaderFooter(false).
				WithLandscape(true).
				Do(ctx)
			return err
		}),
	); err != nil {
		panic(err)
	}

	return &buf
}
