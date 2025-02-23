package analyzer

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"

)

type AnalysisResult struct {
	HTMLVersion    string `json:"html_version"`
	Title          string `json:"title"`
	Headings       map[string]int `json:"headings"`
	InternalLinks  int `json:"internal_links"`
	ExternalLinks  int `json:"external_links"`
	InaccessibleLinks int `json:"inaccessible_links"`
	LoginForm      bool `json:"login_form"`
}

func AnalyzeURL(url string) (*AnalysisResult, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Failed to reach URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error fetching page: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse HTML")
	}

	result := &AnalysisResult{
		Headings: make(map[string]int),
	}

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
						if strings.HasPrefix(attr.Val, "http") {
							result.ExternalLinks++
						} else {
							result.InternalLinks++
						}
					}
				}
			case "form":
				for _, attr := range n.Attr {
					if attr.Key == "action" && strings.Contains(attr.Val, "login") {
						result.LoginForm = true
					}
				}
			case "!DOCTYPE":
				result.HTMLVersion = "HTML5"
			}
		}
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			parseNode(child)
		}
	}

	parseNode(doc)

	return result, nil
}
