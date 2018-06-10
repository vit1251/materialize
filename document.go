package materialize

import "fmt"
import "html/template"

type Style struct {
	URL	string
}

type JavaScript struct {
	URL	string
}

type Node struct {
	cc	*ComponentContent
}

func NewStyle(URL string) (*Style) {
	s := &Style{
		URL: URL,
	}
	return s
}

func NewJavaScript(URL string) (*JavaScript) {
	js := &JavaScript{
		URL: URL,
	}
	return js
}

type DocumentHead struct {
	Title string
	Styles []Style
	JavaScripts []JavaScript
}

type DocumentBody struct {
	Nodes	[]Node
}

type Document struct {
    Head DocumentHead
    Body DocumentBody
}

func NewDocument() (*Document, error) {
	doc := &Document{}
	return doc, nil
}

func (h *DocumentHead) AddStyle(s Style) {
	h.Styles = append(h.Styles, s)
}

func (h *DocumentHead) AddJavaScript(js JavaScript) {
	h.JavaScripts = append(h.JavaScripts, js)
}

func (b *DocumentBody) AddNode(cc *ComponentContent) {
	node := Node{
		cc: cc,
	}
	b.Nodes = append(b.Nodes, node)
}

func (doc *Document) Content() (template.HTML, error) {

	var body string

	body += "<!DOCTYPE html>"
	body += "<html>"
	body += "<head>"
	body += "<meta charset=\"utf-8\">"
	body += fmt.Sprintf("<title>%s</title>", doc.Head.Title)

	// Step 1. Pack CSS styles
	for _, s := range doc.Head.Styles {
		link := fmt.Sprintf("<link rel=\"stylesheet\" href=\"%s\">", s.URL)
		body += link
	}
	// Step 2. Pack JavaScript
	for _, js := range doc.Head.JavaScripts {
		script := fmt.Sprintf("<script src=\"%s\"></script>", js.URL)
		body += script
	}

	body += "</head>"
	body += "<body>"
	// Step 3. Pack component
	for _, node := range doc.Body.Nodes {
		body += fmt.Sprintf( "%s", node.cc.Content() )
	}
	body += "</body>"
	body += "</html>"

	return template.HTML(body), nil
}
