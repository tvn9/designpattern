package main

///////////// ISP (Interface Segregation Principle) ////////////////

// Print interface
type Printer interface {
	Print(d Document)
}

// Scanner interface
type Scanner interface {
	Scan(d Document)
}

// FaxMachine
type FaxMachine interface {
	Fax(d Document)
}

// Document
type Document struct {
	Doc []byte
}

// AllInOneMachine
type MultiFuncMachine interface {
	Printer
	Scanner
	FaxMachine
}

// MonoPrinter
type MonoPrinter struct {
	printer Printer
}

// MonoScanner
type MonoScanner struct {
	scanner Scanner
}

// MonoFax
type MonoFax struct {
	faxer FaxMachine
}

// AllInOne
type AllInOnePrinter struct {
	printer Printer
	scanner Scanner
	faxer   FaxMachine
}

// MonoPrinter
func (p *MonoPrinter) Print(d Document) {
}

// MonoScanner
func (s *MonoScanner) Scan(d Document) {
}

// MonoFax
func (f *MonoFax) Fax(d Document) {

}

// AllInOneMachine
func (a *AllInOnePrinter) Print(d Document) {
	a.printer.Print(d)
}

func (a *AllInOnePrinter) Fax(d Document) {
	a.faxer.Fax(d)
}

func (a *AllInOnePrinter) Scan(d Document) {
	a.scanner.Scan(d)
}

// Program entry point
func main() {
	//

	//

	//

	//
}
