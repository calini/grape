package output

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// Stdout creates a .csv file from records
func Stdout(headers []string, records <-chan []string) error {
	w := tabwriter.NewWriter(
		os.Stdout,
		5,
		5,
		3,
		' ', // TODO figure out what this does,
		tabwriter.Debug|tabwriter.TabIndent,
	)

	defer w.Flush()

	// print headers
	_, err := fmt.Fprintln(w, strings.Join(headers, "\t"))
	if err != nil {
		return fmt.Errorf("error printing table: %v", err)
	}

	// write all records.
	for r := range records {
		_, err = fmt.Fprintln(w, strings.Join(r, "\t"))
		if err != nil {
			return fmt.Errorf("error printing table: %v", err)
		}
	}
	return nil
}

