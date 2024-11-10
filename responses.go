package gokhttp_advanced_responses

import (
	"fmt"
	gokhttp_responses "github.com/BRUHItsABunny/gOkHttp/responses"
	"github.com/Jeffail/gabs"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func ResponseHTML(resp *http.Response) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("goquery.NewDocumentFromReader: %w", err)
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("resp.Body.Close: %w", err)
	}
	return doc, err
}

func ResponseDynamicJSON(resp *http.Response) (*gabs.Container, error) {
	respBytes, err := gokhttp_responses.ResponseBytes(resp)
	if err != nil {
		return nil, fmt.Errorf("gokhttp_responses.ResponseBytes: %w", err)
	}

	container, err := gabs.ParseJSON(respBytes)
	if err != nil {
		return nil, fmt.Errorf("gabs.ParseJSON: %w", err)
	}
	return container, err
}
