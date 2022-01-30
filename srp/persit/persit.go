package persit

import (
	"io/ioutil"
	"net/url"
	"strings"
)

type Journal struct {
	Entries []string
}

// separation of concerns

func (j Journal) String() string {
	return strings.Join(j.Entries, "\n")
}

func (j *Journal) Save(fName string) {
	_ = ioutil.WriteFile(fName, []byte(j.String()), 0644)
}

func (j *Journal) Load(fName string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {

}

// or use stand alone function
var newLine = "\n"

func SaveToFile(j *Journal, fName string) {
	_ = ioutil.WriteFile(fName, []byte(strings.Join(j.Entries, newLine)), 0600)
}
