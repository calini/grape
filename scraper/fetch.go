package scraper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const AttributeSeparator = "§"

// Fetch returns a list of results parsed from the link corresponding to the token
func Fetch(url string, queries []string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("could not get %s: %v", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusTooManyRequests {
			return nil, fmt.Errorf("you are being rate limited")
		}

		return nil, fmt.Errorf("bad response from server: %s", res.Status)
	}

	// parse body with GOQuery.
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("could not parse page: %v", err)
	}

	// extract info we want.
	var r []string
	for _, q := range queries {
		var result string
		if !strings.Contains(q, AttributeSeparator) {
			result = doc.Find(q).Text()
		} else {
			query := strings.Split(q, AttributeSeparator)
			result, _ = doc.Find(query[0]).Attr(query[1])
		}
		r = append(r, strings.TrimSpace(result))
	}
	return r, nil
}
