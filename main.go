package main

import (
	"strconv"
	"strings"

	"github.com/calini/grape/flags"
	"github.com/calini/grape/output"
	"github.com/calini/grape/scraper"

	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
)

const (
	DefaultConcurrency = 1
)

var (
	urlTemplate flags.String
	concurrency flags.Int
	dictfile    flags.String
	idLow       flags.Int
	idHigh      flags.Int
	query       flags.String
	outfile     flags.String
)

// init runs before main and parses the CLI flags
func init() {
	flag.Var(&urlTemplate, "url", "The URL you wish to scrape, containing \"%s\" or \"%d\" where the token will be substituted")
	flag.Var(&concurrency, "concurrency", "How many scrapers to run in parallel. (More scrapers are faster, but more prone to rate limiting or bandwidth issues)")
	flag.Var(&dictfile, "dict", "Filename to import a dictionary from (optional: if not provided, to and from will be used to generate integer IDs to be scraped; if present along to and from, will select the range of words in the dictionary)")
	flag.Var(&idLow, "from", "Start of the search range - inclusive")
	flag.Var(&idHigh, "to", "End of the search range - exclusive")
	flag.Var(&query, "query", "JQuery-style query for the element")
	flag.Var(&outfile, "outfile", "Filename to export the CSV results")
	flag.Parse()

	// check if concurrency set
	if !concurrency.IsSet() {
		concurrency.Value = DefaultConcurrency
	}

	// if no input is set
	if !dictfile.IsSet() && (!idLow.IsSet() || !idHigh.IsSet()) {
		log.Fatal("you must either provide a dictionary file or an index range")
	}
}

func main() {
	headers := strings.Split(query.Value, " ")
	queries := strings.Split(query.Value, " ")
	// url and id are added as the first two columns.
	headers = append([]string{"id"}, headers...)

	// create tasks and send them to the channel.
	tasks := make(chan task)
	go createTasks(tasks)

	// create workers and schedule closing results when all work is done.
	results := make(chan []string)
	var wg sync.WaitGroup
	wg.Add(concurrency.Value)
	go func() {
		wg.Wait()
		close(results)
	}()

	for i := 0; i < concurrency.Value; i++ {
		go func() {
			defer wg.Done()
			for t := range tasks {
				r, err := scraper.Fetch(t.url, queries)
				if err != nil {
					log.Printf("could not fetch %v: %v", t.url, err)
					continue
				} else {
					log.Printf("fetched %v", t.url)
				}
				results <- append([]string{t.token}, r...)
			}
		}()
	}

	if outfile.IsSet() {
		// if we have an outfile, print to csv.
		if err := output.CSV(outfile.Value, headers, results); err != nil {
			log.Printf("could not write to %s: %v", outfile.Value, err)
		}
	} else {
		// else print a table to stdout.
		if err := output.Stdout(headers, results); err != nil {
			log.Printf("could not print table: %v", err)
		}
	}
}

type task struct {
	url   string
	token string
}

func createTasks(tasks chan task) {
	defer close(tasks)
	if dictfile.IsSet() {
		if idLow.IsSet() && idHigh.IsSet() {
			passTasksFromDictRange(urlTemplate.Value, tasks, dictfile.Value, idLow.Value, idHigh.Value)
		} else {
			passTasksFromDict(urlTemplate.Value, tasks, dictfile.Value)
		}
	} else if idLow.IsSet() && idHigh.IsSet() {
		passTasksFromRange(urlTemplate.Value, tasks, idLow.Value, idHigh.Value)
	} else {
		log.Fatal("you must either provide a dictionary file or an index range")
	}
}

func passTasksFromDict(url string, tasks chan task, dictFile string) {
	file, err := os.Open(dictFile)
	if err != nil {
		log.Fatalf("cannot open dictionary file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		tasks <- task{url: fmt.Sprintf(url, t), token: t}
	}
}

func passTasksFromDictRange(url string, tasks chan task, dictFile string, idLow, idHigh int) {
	file, err := os.Open(dictFile)
	if err != nil {
		log.Fatalf("cannot open dictionary file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// skip the first idLow tokens
	for i := 0; i < idLow; i++ {
		scanner.Scan()
		fmt.Println(scanner.Text())
	}
	for i := idLow; scanner.Scan() && i < idHigh; i++ {
		t := scanner.Text()
	}
}

func passTasksFromRange(url string, tasks chan task, idLow, idHigh int) {
	for i := idLow; i < idHigh; i++ {
		tasks <- task{url: fmt.Sprintf(url, i), token: strconv.Itoa(i)}
	}
}
