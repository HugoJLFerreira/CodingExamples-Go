// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package customerimporter

import (
	"encoding/csv"
	"errors"
	"fmt"
	"net/mail"
	"os"
	"sort"
	"strings"
)

//Error constants
const NO_EMAIL_PRESENT = "No email present"
const EMAIL_INVALID = "Email not in a valid format"
const PARSING_ERROR = "Parsing error:"

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func SortDomainEntries(domainEntries map[string]int) PairList {
	keys := make(PairList, len(domainEntries))
	i := 0
	for k, v := range domainEntries {
		keys[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(keys))
	return keys
}

//
func ParseEntry(line []string) (string, error) {
	if len(line) <= 2 {
		return "", errors.New(NO_EMAIL_PRESENT)
	}

	//Check if email and domain are valid
	_, err := mail.ParseAddress(line[2])
	if err != nil {
		return "", errors.New("Email not in a valid format")
	}
	components := strings.Split(line[2], "@")
	return components[1], nil
}

func CollateDomainEntries(csvLines [][]string) map[string]int {
	emailDom := make(map[string]int)

	for _, line := range csvLines {
		domain, err := ParseEntry(line)
		if err != nil {
			fmt.Println("Parsing error:", err)
			continue
		}

		if emailDom[domain] == 0 {
			emailDom[domain] = 1 //Creates domain key on the map and assigns to 1
		} else {
			emailDom[domain]++
		}
	}
	return emailDom
}

func ReadFileContents(file string) ([][]string, error) {
	csvFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		return nil, err
	}

	return csvLines, nil
}
