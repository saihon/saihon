## saihon

JavaScript-like HTML parser for Go language.

<br/>

[![GoDoc](https://pkg.go.dev/badge/github.com/saihon/saihon)](https://pkg.go.dev/github.com/saihon/saihon) 
[![Test](https://github.com/saihon/saihon/actions/workflows/go.yml/badge.svg)](https://github.com/saihon/saihon/actions/workflows/go.yml)

<br>
<br>

## Usage


```go

import (
    "github.com/saihon/saihon"
)

func main() {
    text := "<html><head></head><body></body></html>"
    
    // parse from text HTML
    document, err := saihon.Parse(strings.NewReader(text))
    if err != nil {
       return
    }

    documentElement := document.DocumentElement()
    all     := document.All()
    body    := document.Body()
    title   := document.Title() // title string
    head    := document.Head()
    form    := document.Form()
    images  := document.Images()
    links   := document.Links()
    anchors := document.Anchors()


    element := document.GetElementById("id")
    element = document.QuerySelector("div > p")
    // should be verified
    if element != nil {
        textcontent := element.TextContent()
        // ...
    }


    // returns collection
    elements := document.GetElementsByClassName("class")
    elements = document.QuerySelectorAll("div > p")
    elements = document.GetElementsByName("name")
    elements = document.GetElementsByTagName("p")

    // each element
    for i := 0; i < elements.Length(); i++ {
        outerhtml := elements.Get(i).OuterHTML()
        // ...
    }
    // or 
    for element := range elements.Enumerator() {
        outerhtml := element.OuterHTML()
        // ...
    }


    // set
    element.TextContent("hello")
    // get
    textcontent := element.TextContent()
    // set
    element.InnerHTML("<p>hello</p>")
    // get
    innerhtml := element.InnerHTML()

    // get id
    id := element.HasAttribute("id")
    // get class name
    classname := element.GetAttribute("class")
    // set attribute
    element.SetAttribute("key", "value")
    // remove
    element.RemoveAttribute("key")
}


```

<br>


## License

[MIT License](https://github.com/saihon/saihon/blob/master/LICENSE)

<br>
<br>
