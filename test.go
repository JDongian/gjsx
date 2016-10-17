package main

import (
    "github.com/mamaar/risotto/generator"
    "io/ioutil"
    "os"
    "fmt"
)

func main() {
    jsxFile, err := os.Open("samples/BuildWithReact.jsx")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer jsxFile.Close()

    raw, err := ioutil.ReadFile("samples/BuildWithReact.jsx")
    fmt.Printf("jsx: " + string(raw))

    gen, err := generator.ParseAndGenerate(jsxFile)
    fmt.Println(err)
    result, err := ioutil.ReadAll(gen)
    fmt.Println(err)
    fmt.Printf("js: " + string(result))
}
