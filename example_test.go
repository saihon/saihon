package gohtml_test

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/saihon/gohtml"
	"github.com/saihon/gohtml/utils"
)

func Example() {

	text := `
<html>
<head></head>
<body>
	<div id="id">hello</div>
</body>
</html>`

	doc, err := gohtml.Parse(strings.NewReader(text))
	if err != nil {
		log.Fatal(err)
	}

	v := doc.GetElementById("id")
	if v == nil {
		return
	}
	fmt.Println(v.TextContent()) // hello

	// Attribute
	// set
	v.SetAttribute("class", "class-1")
	// get
	classname := v.GetAttribute("class")
	fmt.Println(classname) // class-1

	// get body
	body := doc.Body()

	// remove
	body.RemoveChild(v)

	// create element
	div := gohtml.CreateElement("div")
	// create text node
	textnode := gohtml.CreateTextNode("hello world")
	div.AppendChild(textnode)
	body.AppendChild(div)

	// remove itself
	div.Remove()

	for _, v := range []string{"foo", "bar", "baz"} {
		p := gohtml.CreateElement("p")
		t := gohtml.CreateTextNode(v)
		p.AppendChild(t)
		body.AppendChild(p)
	}

	collection := body.GetElementsByTagName("p")
	for i := 0; i < collection.Length(); i++ {
		element := collection.Get(i)
		fmt.Println(element.InnerText())
	}

	utils.Empty(body.Node)
}

// This example shows how to use parse from string
func ExampleParse_string() {
	text := `<html><head></head><body></body></html>`

	document, err := gohtml.Parse(strings.NewReader(text))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(document.Title())
}

// This example shows how to use parse from bytes
func ExampleParse_bytes() {
	text := []byte(`<html><head></head><body></body></html>`)

	document, err := gohtml.Parse(bytes.NewReader(text))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(document.Title())
}

// This example shows how to use parse from file
func ExampleParse_file() {
	fp, err := os.Open("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	document, err := gohtml.Parse(fp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(document.Title())
}

// This example shows how to use parse from http response
func ExampleParse_httpresponse() {
	resp, err := http.Get("https://example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	document, err := gohtml.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(document.Title())
}
