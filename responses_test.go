package gokhttp_advanced_responses

import (
	"context"
	"fmt"
	gokhttp "github.com/BRUHItsABunny/gOkHttp"
	gokhttp_requests "github.com/BRUHItsABunny/gOkHttp/requests"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHTMLParser(t *testing.T) {
	hClient, err := gokhttp.NewHTTPClient()
	require.NoError(t, err, "NewHTTPClient: errored unexpectedly.")

	req, err := gokhttp_requests.MakeGETRequest(context.Background(), "https://github.com")
	require.NoError(t, err, "requests.MakeGETRequest: errored unexpectedly.")

	resp, err := hClient.Do(req)
	require.NoError(t, err, "hClient.Do: errored unexpectedly.")

	respHTML, err := ResponseHTML(resp)
	require.NoError(t, err, "ResponseHTML: errored unexpectedly.")

	fmt.Println(respHTML.Text())
}
