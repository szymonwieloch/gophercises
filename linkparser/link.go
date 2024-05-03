package linkparser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Link struct {
	Href string
	Text string
}

// Parses the provided Reader
func ParseReader(r io.Reader) ([]Link, error) {
	node, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	result := recursiveSearch(node)
	return result, nil

}

func getNodeText(n *html.Node) string {
	var parts []string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			parts = append(parts, getNodeText(c))
		} else
		// fmt.Println("node data", c.Data, "type", c.Type)
		if c.Type == html.TextNode {
			parts = append(parts, c.Data)
		}
	}
	return strings.Join(parts, "")
}

func getHrefAttr(n *html.Node) string {
	href := ""
	for _, a := range n.Attr {
		if a.Key == "href" {
			href = a.Val
			break
		}
	}
	return href
}

func recursiveSearch(n *html.Node) []Link {
	result := []Link{}
	if n.Type == html.ElementNode && n.DataAtom == atom.A {

		result = append(result, Link{
			Href: getHrefAttr(n),
			Text: getNodeText(n),
		})
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		subLinks := recursiveSearch(c)
		result = append(result, subLinks...)
	}
	return result
}
