package analyzer

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

type AnalysisResult struct {
	HTMLVersion       string         `json:"html_version"`
	Title             string         `json:"title"`
	Headings          map[string]int `json:"headings"`
	InternalLinks     int            `json:"internal_links"`
	ExternalLinks     int            `json:"external_links"`
	InaccessibleLinks int            `json:"inaccessible_links"`
	LoginForm         bool           `json:"login_form"`
}

func AnalyzeURL(url string) (*AnalysisResult, int, error) {
	if url == "" {
		return nil, http.StatusBadRequest, fmt.Errorf("URL cannot be empty")
	}

	resp, err := http.Get(url)

	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("failed to reach url")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, http.StatusNotFound, fmt.Errorf("page is not accessible")
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, fmt.Errorf("failed to parse page")
	}

	result := &AnalysisResult{
		Headings: make(map[string]int),
	}

	var allLinks []string

	var parseNode func(*html.Node)
	parseNode = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "title":
				if n.FirstChild != nil {
					result.Title = n.FirstChild.Data
				}
			case "h1", "h2", "h3", "h4", "h5", "h6":
				result.Headings[n.Data]++
			case "a":
				for _, attr := range n.Attr {
					if attr.Key == "href" {
						allLinks = append(allLinks, attr.Val)
						if strings.HasPrefix(attr.Val, "http") || strings.HasPrefix(attr.Val, "https") {
							result.ExternalLinks++
						} else {
							result.InternalLinks++
						}
					}
				}
			case "button":
				buttonText := strings.ToLower(n.FirstChild.Data)
				if strings.Contains(buttonText, "log in") ||
					strings.Contains(buttonText, "sign in") ||
					strings.Contains(buttonText, "sign up") {
					result.LoginForm = true
				}
			case "html":
				if strings.Contains(n.FirstChild.Data, "html") {
					if strings.Contains(n.FirstChild.Data, "PUBLIC \"-//W3C//DTD HTML 4.01//EN\"") {
						result.HTMLVersion = "HTML4"
					} else if strings.Contains(n.FirstChild.Data, "PUBLIC \"-//W3C//DTD XHTML 1.0 Strict//EN\"") {
						result.HTMLVersion = "XHTML1.0"
					} else if strings.Contains(n.FirstChild.Data, "PUBLIC \"-//W3C//DTD HTML 3.2//EN\"") {
						result.HTMLVersion = "HTML3.2"
					} else {
						result.HTMLVersion = "HTML5"
					}
				}
			}
		}
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			parseNode(child)
		}
	}

	parseNode(doc)

	var wg sync.WaitGroup
	inaccessibleLinksChan := make(chan string, len(allLinks))

	for _, link := range allLinks {
		wg.Add(1)
		go func(link string) {
			defer wg.Done()
			_, err := http.Get(link)
			if err != nil {
				inaccessibleLinksChan <- link
			}
		}(link)
	}

	go func() {
		wg.Wait()
		close(inaccessibleLinksChan)
	}()

	for range inaccessibleLinksChan {
		result.InaccessibleLinks++
	}

	return result, http.StatusOK, nil
}
