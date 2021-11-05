package spider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func PrintAndGetNextUrl (url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch reading %s: %v\n", url, err)
		os.Exit(1)
	}
	body := string(b[:])
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse html failed: %v\n", err)
		os.Exit(1)
	}
	doc.Find("a[title]").Each(func(i int, selection *goquery.Selection) {
		title, exists := selection.Attr("title")
		if exists {
			href, e := selection.Attr("href")
			if e {
				_, ex := selection.Children().Attr("src")
				if ex {
					fmt.Printf("title: %s, href: %s\n", title, href)
				}
			}
		}
	})
	next := ""
	doc.Find("a.next").Each(func(i int, selection *goquery.Selection) {
		href, exists := selection.Attr("href")
		if exists {
			next = href
		}
	})
	return next
}