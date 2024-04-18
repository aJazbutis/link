package link

import (
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func extractText(n *html.Node, b *strings.Builder) {
	if n == nil {
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			b.WriteString(c.Data)
		}
		if c.Type == html.ElementNode {
			extractText(c, b)
		}
	}
}

func extractLink(n *html.Node) Link {
	var ret Link
	for _, a := range n.Attr {
		if a.Key == "href" {
			ret.Href = a.Val
			break
		}
	}
	var b strings.Builder
	extractText(n, &b)
	ret.Text = strings.Join(strings.Fields(b.String()), " ")
	return ret
}

func ExtractLinks(path *string) []Link {
	file, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(file)
	if err != nil {
		log.Fatal(err)
	}
	var links []Link
	var f func(*html.Node, *[]Link)
	f = func(root *html.Node, links *[]Link) {
		if root == nil {
			return
		}
		if root.Type == html.ElementNode && root.Data == "a" {
			*links = append(*links, extractLink(root))
		}
		for c := root.FirstChild; c != nil; c = c.NextSibling {
			f(c, links)
		}
	}
	f(doc, &links)
	return links
}
