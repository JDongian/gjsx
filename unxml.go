package main

import (
    "os"
    "fmt"
    "xml"
)

type Result struct {
    XMLName xml.Name `xml:"div"`
}

func main() {
    var v Result

    jsxFile, err := os.Open("samples/BuildWithReact.jsx")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer xmlFile.Close()

    xml.Unmarshal(jsxFile, &v)

    fmt.Printf("XMLName: %#v\n", v.XMLName)
}
