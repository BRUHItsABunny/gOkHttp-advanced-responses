package gokhttp_advanced_responses

import (
	"fmt"
	gokhttp_responses "github.com/BRUHItsABunny/gOkHttp/responses"
	"github.com/Jeffail/gabs"
	"github.com/PuerkitoBio/goquery"
	"github.com/t14raptor/go-fast/ast"
	"github.com/t14raptor/go-fast/parser"
	"net/http"
	"strings"
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

func ResponseAST(resp *http.Response) (*ast.Program, error) {
	respText, err := gokhttp_responses.ResponseText(resp)
	if err != nil {
		return nil, fmt.Errorf("gokhttp_responses.ResponseText: %w", err)
	}

	astParser, err := parser.ParseFile(respText)
	if err != nil {
		return nil, fmt.Errorf("parser.ParseFile: %w", err)
	}

	return astParser, nil
}

func ResponseASTFromHTML(resp *http.Response) (*ast.Program, error) {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("goquery.NewDocumentFromReader: %w", err)
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("resp.Body.Close: %w", err)
	}

	jsSB := strings.Builder{}

	doc.Find("script").Each(func(index int, element *goquery.Selection) {
		jsSB.WriteString(element.Text())
		jsSB.WriteString("\n")
	})

	astParser, err := parser.ParseFile(jsSB.String())
	if err != nil {
		return nil, fmt.Errorf("parser.ParseFile: %w", err)
	}

	return astParser, nil
}
