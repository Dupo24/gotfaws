package main

import (
	"fmt"
	// "strings"
	// "time"
	"github.com/gocolly/colly"
)

type mainPage struct {
	provider string
	resources string
	resource_pagelink string

}
type resources struct {

}
 func ProviderWebScraper() {
	c := colly.NewCollector(
		//colly.AllowedDomains("registry.terraform.io"),
		colly.MaxDepth(1),
		colly.Async(true),
	)
 }


func main() {
	
	c := colly.NewCollector(
		//colly.AllowedDomains("registry.terraform.io"),
		colly.MaxDepth(1),
		colly.Async(true),
	)

	//callbacks
	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})
	
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
	
	c.OnXML("//h1", func(e *colly.XMLElement) {
		fmt.Println(e.Text)
	})
	
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	//grab main AWS provider page
	c.Visit("https://registry.terraform.io/providers/hashicorp/aws/latest/docs/")

}


//https://registry.terraform.io/providers/hashicorp/aws/latest/docs