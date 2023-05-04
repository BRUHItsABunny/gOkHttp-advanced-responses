package gokhttp_advanced_responses

import (
	"bytes"
	"fmt"
	gokhttp_responses "github.com/BRUHItsABunny/gOkHttp/responses"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func ResponseHTML(resp *http.Response) (*goquery.Document, error) {
	respBytes, err := gokhttp_responses.ResponseBytes(resp)
	if err != nil {
		return nil, fmt.Errorf("responses.ResponseBytes: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(respBytes))
	if err != nil {
		return nil, fmt.Errorf("goquery.NewDocumentFromReader: %w", err)
	}
	return doc, err
}
