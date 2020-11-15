package services

import (
	"github.com/PuerkitoBio/goquery"
)

type RequestService struct{}

type ResultStruct struct {
	Html string `json:"html"`
}

func (_ RequestService) Call(link string) ResultStruct {
	doc, err := goquery.NewDocument(link)
	if err != nil {
		panic(err)
	}
	node := doc.Find(".answercell .s-prose")
	html, err := node.Html()
	return ResultStruct{Html: html}
}
