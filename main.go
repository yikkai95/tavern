package main

import (
	"encoding/json"
	"os"
	"text/template"
)

func main() {
	in, err := os.Open("data.json")
	check(err)
	defer in.Close()

	var data []struct {
		Key     int    `json:"key"`
		Caption string `json:"caption"`
		Date    string `json:"date"`
	}

	err = json.NewDecoder(in).Decode(&data)
	check(err)

	t := template.Must(template.ParseFiles("index.tmpl"))

	out, err := os.Create("index.html")
	check(err)
	defer out.Close()

	err = t.Execute(out, data)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
