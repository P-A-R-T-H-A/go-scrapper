package main

import (
	"os"
	"web_spider/spider"

	"github.com/gocolly/colly/v2"
	env "github.com/joho/godotenv"
)

func main() {
	env.Load(".env")
	c := colly.NewCollector(
		colly.AllowedDomains(os.Getenv("BASE_DOMAIN")),
	)
	fetched := spider.FetchCategoryIds(c)
	if fetched {
		spider.FetchProductDetails(c)
	}

}
