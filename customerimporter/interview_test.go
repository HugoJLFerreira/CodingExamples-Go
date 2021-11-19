package customerimporter

import (
	"fmt"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestFileReadingNoFile(t *testing.T) {
	_, err := ReadFileContents("../data/customers_none.csv")
	assertEqual(t, err.Error(), "open ../data/customers_none.csv: The system cannot find the file specified.", ("Unexpected error reading from file:" + err.Error()))
}

func TestFileReading(t *testing.T) {
	lines, err := ReadFileContents("../data/customers.csv")
	assertEqual(t, err, nil, "There was a problem reading from FileContents")
	assertEqual(t, len(lines), 3003, "Amount of entries do not match.")
}

func TestDomainParsingHappyPath(t *testing.T) {
	wellformedData := [][]string{
		{"Mildred", "Hernandez", "mhernandez0@github.io", "Female", "38.194.51.128"},
		{"Bonnie", "Ortiz", "bortiz1@cyberchimps.com", "Female", "197.54.209.129"},
		{"Dennis", "Henry", "dhenry2@hubpages.com", "Male", "155.75.186.217"},
		{"Justin", "Hansen", "jhansen3@360.cn", "Male", "251.166.224.119"},
		{"Carlos", "Garcia", "cgarcia4@statcounter.com", "Male", "57.171.52.110"},
	}

	domains := []string{"github.io", "cyberchimps.com", "hubpages.com", "360.cn", "statcounter.com"}

	for i := 0; i < len(wellformedData); i++ {
		domain, err := ParseEntry(wellformedData[i])
		assertEqual(t, err, nil, "There was a problem reading from FileContents")
		assertEqual(t, domain, domains[i], ("Domains do not match Actual:" + domain + " - Expected:" + domains[i]))
	}
}

func TestDomainParsingFailurePath(t *testing.T) {
	malformedData := [][]string{
		{"first_name", "last_name", "email", "gender", "ip_address"},
		{"Bonnie", "Ortiz", "@cybe@rchimps.com", "Female", "197.54.209.129"},
		{"Dennis", "Henry", "dhenry2@hubpages.", "Male", "155.75.186.217"},
		{"Justin", "Hansen", "jhansen3@", "Male", "251.166.224.119"},
		{"Carlos", "Garcia", "c@...", "Male", "57.171.52.110"},
		{"Carlos"}}

	failures := []string{EMAIL_INVALID, EMAIL_INVALID, EMAIL_INVALID, EMAIL_INVALID, EMAIL_INVALID, NO_EMAIL_PRESENT}

	for i := 0; i < len(malformedData); i++ {
		_, err := ParseEntry(malformedData[i])
		// assertEqual(t, domain, nil, ("No domains should be returned. Returned:" + domain))
		assertEqual(t, err.Error(), failures[i], "There was a problem reading from FileContents")
	}
}
