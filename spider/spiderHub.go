package spider

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"web_spider/data"
	"web_spider/utils"

	"github.com/gocolly/colly/v2"
)

func FetchCategoryIds(c *colly.Collector) bool {
	targetUrl := os.Getenv("CATEGORY_PAGE_URL")
	pageNumber := os.Getenv("PAGE_NUMBER")
	success := false
	pageNum, _ := strconv.Atoi(pageNumber)
	counter := 1
	for {
		url := targetUrl + strconv.Itoa(counter)
		Scrapper(c, url, GetProductIDs)
		if counter == pageNum {
			break
		}
		counter++
	}
	//write PropertyIds in Json
	if len(data.PropertyList) > 0 {
		utils.WriteToJson(data.PropertyList, "/data/json/product-list.json")
		success = true
	}
	return success
}

func FetchProductDetails(c *colly.Collector) {
	targetUrl := os.Getenv("DETAILS_PAGE_URL")
	propertyIds := []string{}
	path, err := os.Getwd()
	if err != nil {
		log.Printf("Error in reading current path")
	}
	directory := path + "/data/json/product-list.json"
	byteArray, err := os.ReadFile(directory)
	if err != nil {
		log.Println("Error in reading file: ", err)
	} else {
		err := json.Unmarshal(byteArray, &propertyIds)
		if err != nil {
			log.Println("Error in unmarshal propertyIds: ", err)
		} else {
			for _, Id := range propertyIds {
				url := targetUrl + Id
				Scrapper(c, url, GetProductDetails)
			}
		}
	}

}
