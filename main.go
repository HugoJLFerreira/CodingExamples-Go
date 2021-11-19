package main

import (
	"TeamworkGoTests/customerimporter"
	"fmt"
)

func main() {
	lines, err := customerimporter.ReadFileContents("data/customers.csv")
	if err != nil {
		fmt.Println(err)
	} else {
		//In Go 1.12+, you can just print a map value and it will be sorted by key automatically. Thus records returns ordered by domain name.
		records := customerimporter.CollateDomainEntries(lines)

		//Ordered by highest email count.
		sortedEntries := customerimporter.SortDomainEntries(records)
		fmt.Println(sortedEntries)
	}
}
