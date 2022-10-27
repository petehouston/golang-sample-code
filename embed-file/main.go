package main

import (
	"embed"
	_ "embed"
)

//go:embed "data/*"
var fs embed.FS

func main() {
	companies, _ := fs.ReadFile("data/companies.txt")
	println("Companies: \n" + string(companies))

	countries, _ := fs.ReadFile("data/countries.txt")
	println("Countries: \n" + string(countries))
}
