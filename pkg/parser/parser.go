package parser

import (
	"errors"
	"fmt"
	"net/http"

	pb "github.com/DimkaTheGreat/testTaskStrafovNet/pkg/api"
	"github.com/PuerkitoBio/goquery"
)

const queryURL = "https://www.rusprofile.ru/search?query=%v"

func GetCompanyInfo(inn string) (*pb.Response, error) {
	resp, err := http.Get(fmt.Sprintf(queryURL, inn))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	companyInfo := &pb.Response{}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	doc.Find("span").EachWithBreak(func(index int, item *goquery.Selection) bool {
		val, ok := item.Attr("id")
		if ok && val == "clip_inn" {
			companyInfo.INN = item.Text()
			return false

		}
		return true
	})

	if companyInfo.INN == "" {
		return nil, errors.New("Company not found")
	}

	doc.Find("span").EachWithBreak(func(index int, item *goquery.Selection) bool {
		val, ok := item.Attr("id")
		if ok && val == "clip_kpp" {
			companyInfo.KPP = item.Text()
			return false

		}
		return true
	})

	doc.Find("div").EachWithBreak(func(index int, item *goquery.Selection) bool {
		val, ok := item.Attr("itemprop")
		if ok && val == "legalName" {
			companyInfo.Name = item.Text()
			return false

		}
		return true
	})

	doc.Find("span").EachWithBreak(func(index int, item *goquery.Selection) bool {
		val, ok := item.Attr("class")
		if ok && val == "company-info__text" {
			companyInfo.Leader = item.Text()
			return false

		}
		return true
	})

	return companyInfo, nil

}
