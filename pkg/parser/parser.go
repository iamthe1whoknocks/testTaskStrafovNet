package parser

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	pb "github.com/DimkaTheGreat/testTaskStrafovNet/proto/testTaskStrafovNet"
	"github.com/PuerkitoBio/goquery"
)

const queryURL = "https://www.rusprofile.ru/search?query=%v"

func GetCompanyInfo(inn string) (*pb.Response, error) {
	resp, err := http.Get(fmt.Sprintf(queryURL, inn))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	isSuccessSearch(doc)

	description, err := getDescription(doc)

	if err != nil {
		return nil, err
	}

	num := getResultNumber(description)

	if num == "0" {
		return nil, errors.New("Company not found")
	}

}

//Получение description для анализа (при вводе ИНН в поисковую строку возможна множественная выдача)
func getDescription(doc *goquery.Document) (resultNum string, err error) {
	var metaDescription string
	doc.Find("meta").EachWithBreak(func(index int, item *goquery.Selection) bool {
		if item.AttrOr("name", "") == "description" {
			metaDescription = item.AttrOr("content", "")
			return false
		}
		return true
	})
	return metaDescription, nil

}

//количество результатов поисковой выдачи
func getResultNumber(desc string) (resultNum string) {
	slice := strings.Split(desc, " ")
	return slice[5]

}

func parseCompanyInfo(doc *goquery.Document) (companyInfo *pb.Response, err error) {

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

func isSuccessSearch(doc *goquery.Document) (successSearch bool) {
	doc.Find("title").EachWithBreak(func(index int, item *goquery.Selection) bool {
		if strings.Contains(item.Text(), "результаты поиска") == true {
			successSearch = false
			return false
		}
		successSearch = true
		return true
	})
	return false
}

//Если множественный вывод (пример : 1656002652)
