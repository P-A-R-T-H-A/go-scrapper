package data

import "github.com/gocolly/colly/v2"

var (
	PropertyList []string
)

type HtmlFuncProto func(*colly.Collector)
