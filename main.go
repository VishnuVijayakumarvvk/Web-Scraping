package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type cloth struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	ImageURL string `json:"image"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("j2store.net"),
	)
	var cloths []cloth
	c.OnHTML("div.col-sm-9 div[itemprop=itemListElement]", func(h *colly.HTMLElement) {
		cloth := cloth{
			Name:     h.ChildText("h2.product-title"),
			Price:    h.ChildText("div.sale-price"),
			ImageURL: h.ChildAttr("img", "src"),
		}
		cloths = append(cloths, cloth)

		//fmt.Println(h.ChildText("h2.product-title"))
		//fmt.Println(h.ChildText("div.sale-price"))
		//fmt.Println(h.ChildAttr("img","src"))
	})
	c.OnHTML("", func(h *colly.HTMLElement) {
		Nlink := h.Request.AbsoluteURL(h.Attr("href"))
		c.Visit(Nlink)
	})
	c.Visit("https://j2store.net/demo/index.php/shop")
	sourse, err := json.MarshalIndent(cloths, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("./ClothData.json", sourse, 0644)
}
