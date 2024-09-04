package spider

import (
	"encoding/json"
	"fmt"
	"web_spider/data"
	"web_spider/data/structure"

	"github.com/gocolly/colly/v2"
	gJSON "github.com/tidwall/gjson"
)

func Scrapper(c *colly.Collector, targetUrl string, function data.HtmlFuncProto) {
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error: ", err)
	})
	function(c)
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Successfully scrapped: ", r.Request.URL)
	})
	c.Visit(targetUrl)
}

func GetProductIDs(c *colly.Collector) {
	c.OnHTML("script#__NEXT_DATA__", func(html *colly.HTMLElement) {
		if _, ok := gJSON.Parse(html.Text).Value().(map[string]interface{}); ok {
			//TODO:: need to give a value check for preventing data anomaly; also need to optimize the process
			gJSON.Parse(html.Text).Get("props").Get("pageProps").Get("apis").Get("plpInitialProps").Get("productListApi").Get("articles_sort_list").ForEach(func(key, value gJSON.Result) bool {
				data.PropertyList = append(data.PropertyList, value.Str)
				return true
			})
		}
	})
}
func GetProductDetails(c *colly.Collector) {
	c.OnHTML("script#__NEXT_DATA__", func(html *colly.HTMLElement) {
		var (
			productBreadCrumb = []structure.BreadCrumb{}
			ProductDetails    = structure.ProductDetails{}
		)
		if _, ok := gJSON.Parse(html.Text).Value().(map[string]interface{}); ok {
			//TODO:: need to give a value check for preventing data anomaly; also need to optimize the process
			mainFrame := gJSON.Parse(html.Text).Get("props").Get("pageProps").Get("apis").Get("pdpInitialProps").Get("detailApi")
			page := mainFrame.Get("page")
			err := json.Unmarshal([]byte(page.Get("breadcrumbs").String()), &productBreadCrumb)
			if err != nil {
				fmt.Println(err)
			}
			product := mainFrame.Get("product")
			err = json.Unmarshal([]byte(product.Get("breadcrumbs").String()), &ProductDetails)
			if err != nil {
				fmt.Println(err)
			}
		}
	})
}
