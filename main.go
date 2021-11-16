package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
)

type Link struct {
	url   string
	title string
}

func CrawlLink(url string) []Link {
	var result []Link

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
		colly.Async(),
		colly.Debugger(&debug.LogDebugger{}),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// time.Sleep(5 * time.Second)
		link := e.Attr("href")
		str := strings.TrimSpace(e.Text)
		result = append(result, Link{url: link, title: str})
	})

	c.Visit(url)
	c.Wait()

	return result
}

func main() {
	data := CrawlLink("https://ssr-assessment.miqdadyyy.vercel.app/")
	// data := CrawlLink("https://csr-assessment.miqdadyyy.vercel.app/")
	log.Println(data)
}
