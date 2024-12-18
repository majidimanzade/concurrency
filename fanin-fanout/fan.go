package faninfanout

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

func Crawl(websites ...[]string) {
	websitesList := map[int][]string{}
	for k, v := range websites {
		websitesList[k] = v
	}

	var wg sync.WaitGroup
	numJobs := 3
	inputs := make(map[int]chan string)
	results := make(chan string)

	for k, w := range websites {
		inputs[k] = make(chan string)
		go generateData(w, inputs[k])
	}

	for k := range websites {
		for i := 0; i < numJobs; i++ {
			wg.Add(1)
			go worker(inputs[k], results, fmt.Sprintf("%d %d", k, i), &wg)
		}
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println(r)
	}
}

func worker(inputs chan string, results chan string, identiti string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range inputs {
		title, err := getWebsiteTitle(i)
		if err != nil {
			results <- fmt.Sprintf("Error! Worker %s Proceed %s", identiti, i)
			continue
		}
		results <- fmt.Sprintf("Worker %s Proceed %s Title: %s", identiti, i, title)
	}
}

func generateData(websites []string, ch chan string) {
	defer close(ch)
	for _, v := range websites {
		ch <- v
	}
}

func getWebsiteTitle(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status %d", resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML: %w", err)
	}

	var title string
	var findTitle func(*html.Node)
	findTitle = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			title = strings.TrimSpace(n.FirstChild.Data)
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findTitle(c)
		}
	}
	findTitle(doc)

	if title == "" {
		return "", fmt.Errorf("title not found")
	}
	return title, nil
}
