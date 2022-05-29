package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type parallelStruct struct {
	i int
	s string
	e error
}

//var localeList []string = []string{"en-US, nl-NL"}
var currencyList []string = []string{"EUR", "USD"}

func NewParallelStruct(i int, s string, e error) parallelStruct {
	return parallelStruct{i: i, s: s, e: e}
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	ok := false
	for _, c := range currencyList {
		if currency == c {
			ok = true
			break
		}
	}

	if !ok {
		return "", errors.New("non existant currency")
	}

	var entriesCopy []Entry
	entriesCopy = append(entriesCopy, entries...)

	sort.Slice(entriesCopy,
		func(i, j int) (ret bool) {
			iDay, _ := time.Parse("2006-01-02", entriesCopy[i].Date)
			jDay, _ := time.Parse("2006-01-02", entriesCopy[j].Date)

			if iDay.Before(jDay) { //Check Date
				ret = true
			} else if jDay.Before(iDay) {
				ret = false
			} else { //if equal, check description
				if entriesCopy[i].Description < entriesCopy[j].Description {
					ret = true
				} else if entriesCopy[i].Description > entriesCopy[j].Description {
					ret = false
				} else {
					if entriesCopy[i].Change < entriesCopy[j].Change {
						ret = true
					} else {
						ret = false
					}
				}
			}

			return
		},
	)

	var s string
	switch locale {
	case "nl-NL":
		s = getBeginningLine("Datum", "Omschrijving", "Verandering")
	case "en-US":
		s = getBeginningLine("Date", "Description", "Change")
	default:
		return "", errors.New("")
	}

	// Parallelism, always a great idea
	co := make(chan parallelStruct)
	for i, et := range entriesCopy {

		go func(i int, entry Entry) {

			day, err := time.Parse("2006-01-02", entry.Date)
			if err != nil {
				co <- NewParallelStruct(0, "", errors.New(""))
			}

			var d string
			if locale == "nl-NL" {
				d = day.Format("02-01-2006")
			} else if locale == "en-US" {
				d = day.Format("01/02/2006")
			}

			de := entry.Description
			if len(de) > 25 {
				de = de[:22] + "..."
			} else {
				de = de + strings.Repeat(" ", 25-len(de))
			}

			var a string = formatPrice(entry.Change, locale, currency)

			al := len([]rune(a))
			co <- NewParallelStruct(i,
				d+strings.Repeat(" ", 10-len(d))+" | "+de+" | "+strings.Repeat(" ", 13-al)+a+"\n",
				nil)
		}(i, et)
	}

	ss := make([]string, len(entriesCopy))
	for range entriesCopy {
		v := <-co
		if v.e != nil {
			return "", v.e
		}
		ss[v.i] = v.s
	}

	s += strings.Join(ss, "")
	return s, nil
}

func getBeginningLine(date, description, change string) string {
	return date +
		strings.Repeat(" ", 10-len(date)) +
		" | " +
		description +
		strings.Repeat(" ", 25-len(description)) +
		" | " + change + "\n"
}

func formatPrice(cents int, locale string, currency string) string {
	negative := false
	if cents < 0 {
		cents *= -1
		negative = true
	}
	var money float64 = float64(cents) / 100

	var floatingPoint, milSep string
	switch locale {
	case "nl-NL":
		floatingPoint = ","
		milSep = "."
	default:
		floatingPoint = "."
		milSep = ","
	}
	moneyStr := strings.Replace(fmt.Sprintf("%.2f", money), ".", floatingPoint, -1)

	for i := len(moneyStr) - 6; i > 0; i -= 3 {
		moneyStr = moneyStr[:i] + milSep + moneyStr[i:]
	}

	var currRune string
	switch currency {
	case "EUR":
		currRune = "€"
	case "USD":
		currRune = "$"
	default:
		currRune = " "
	}
	if locale == "en-US" {
		moneyStr = currRune + moneyStr
	}

	if negative {
		switch locale {
		case "en-US":
			moneyStr = "(" + moneyStr + ")"
		default:
			moneyStr = " " + moneyStr + "-"
		}
	} else {
		moneyStr = " " + moneyStr + " "
	}

	if locale == "nl-NL" {
		moneyStr = currRune + moneyStr
	}

	return moneyStr
}

/*
func spaceRepeat(str string, neededLength int) string {
	return strings.Repeat(" ", max(0, neededLength-len(str)))
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
*/
