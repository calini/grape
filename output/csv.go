package output

import (
	"encoding/csv"
	"fmt"
	"os"
)

// CSV creates a file at path then writes the contents in CSV format
func CSV(path string, headers []string, records <-chan []string) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to create file %s: %v", path, err)
	}
	defer f.Close()

	w := csv.NewWriter(f)

	// write headers to file.
	if err := w.Write(headers); err != nil {
		return fmt.Errorf("error writing record to csv: %v", err)
	}

	// write all records.
	for r := range records {
		if err := w.Write(r); err != nil {
			return fmt.Errorf("could not write record to csv: %v", err)
		}
	}

	w.Flush()

	// check for extra errors.
	if err := w.Error(); err != nil {
		return fmt.Errorf("writer failed: %v", err)
	}
	return nil
}
