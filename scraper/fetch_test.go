package scraper

import (
	"fmt"
	"testing"
)

type testData struct {
	url     string
	queries []string
	result  []string
}

func TestFetch(t *testing.T) {

	Test := []testData{
		{
			"https://github.com/calini",
			[]string{".p-name", ".p-org"},
			[]string{"Calin Ilie", "The Hut Group"},
		},
		{
			"https://github.com/philipithomas",
			[]string{".p-name", ".p-org"},
			[]string{"Philip I. Thomas", "@moonlightwork"},
		},
	}

	for _, td := range Test {
		o, err := Fetch(td.url, td.queries)
		if err != nil {
			fmt.Printf("could not get %s: %v", td.url, err)
			continue
		}
		for i, v := range o {
			if v != td.result[i] {
				t.Error("Expected", td.result[i], "Got", v)
			}
		}
	}

}

func ExampleFetch() {
	o, _ := Fetch("https://github.com/calini", []string{".p-name", ".p-org"})
	fmt.Println(o)
	// Output:
	// [Calin Ilie The Hut Group]
}
