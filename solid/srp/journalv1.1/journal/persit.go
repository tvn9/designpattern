package journal

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

/////////////////////// separation of concerns ////////////////////////

func (j *Journal) String() string {
	return fmt.Sprintln(strings.Join(j.Entries, "\n"))
}

func (j *Journal) Save(fName string) {
	_ = ioutil.WriteFile(fName, []byte(strings.Join(j.Entries, "\n")), 0644)
}

func (j *Journal) Load(fName string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {

}

// or use stand alone function
var newLine = "\n"

// Saive file using the stanalone function
func SaveToFile(j *Journal, fName string) {
	_ = ioutil.WriteFile(fName, []byte(strings.Join(j.Entries, newLine)), 0600)
}

// Persistence struct
type Persistence struct {
	newLine string
}

// or a using a method
func (p *Persistence) SaveToFile(j *Journal, fname string) {
	_ = ioutil.WriteFile(fname, []byte(strings.Join(j.Entries, p.newLine)), 0644)
}
