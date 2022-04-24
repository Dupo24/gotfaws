package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	//"strconv"
	// "strings"
	// "time"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

type mainPage struct {
	Provider string
	//resources string
	//link string

}
type resources struct {

}
 func ProviderWebScraper() {
	mainpages := []mainPage{}

	c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))
	
	//callbacks
	// On every a element which has href attribute call callback
	c.OnHTML("link[href]", func(e *colly.HTMLElement) {
		resources := e.DOM
		fmt.Println("found link")
		fmt.Println(e)
		mainpage := mainPage{
			Provider: resources.Text(),//.Find("div.ember-view").Text(),
			// Resources:,
			// Link:,
		} 
		mainpages = append(mainpages, mainpage)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	c.Visit("https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources")
	buildTable(mainpages)
 }

 func buildTable(data []mainPage) { 
		f, err := json.MarshalIndent(data, "", " ")
		if err != nil {
			log.Fatal(err)
			return
		}
		_ = ioutil.WriteFile("resources.json", f, 0644)
 }

func main() {
	ProviderWebScraper()
}


//https://registry.terraform.io/providers/hashicorp/aws/latest/docs