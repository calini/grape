package scraper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Fetch returns a list of rsults parsed from the link corresponding to the token
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

	// parse body with goquery.
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("could not parse page: %v", err)
	}

	// extract info we want.
	r := []string{}
	for _, q := range queries {
		r = append(r, strings.TrimSpace(doc.Find(q).Text()))
	}
	return r, nil
}
