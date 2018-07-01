// Package ledger implements a simple ledger
package ledger

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Entry is an entry in the ledger
type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
	dateTime    time.Time
}

type ledger []Entry

func (l ledger) Less(i, j int) bool {
	e1, e2 := l[i], l[j]

	if e1.dateTime.Equal(e2.dateTime) {
		if e1.Description == e2.Description {
			return e1.Change < e2.Change
		}
		return e1.Description < e2.Description
	}

	return e1.dateTime.Before(e2.dateTime)
}

func (l ledger) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l ledger) Len() int {
	return len(l)
}

type localisation struct {
	translation
	dateFormat                                 string
	decimalSeparator, thousandsSeparator       byte
	positiveFormatString, negativeFormatString string
}

type translation struct {
	date, description, change string
}

const headerFormatString = "%-10s | %-25s | %s\n"
const entryFormatString = "%-10s | %-25s | %13s\n"

var localisations = map[string]localisation{
	"nl-NL": {
		translation: translation{
			"Datum", "Omschrijving", "Verandering",
		},
		dateFormat:           "02-01-2006",
		decimalSeparator:     ',',
		thousandsSeparator:   '.',
		positiveFormatString: "%c %s ",
		negativeFormatString: "%c %s-",
	},
	"en-US": {
		translation: translation{
			"Date", "Description", "Change",
		},
		dateFormat:           "01/02/2006",
		decimalSeparator:     '.',
		thousandsSeparator:   ',',
		positiveFormatString: "%c%s ",
		negativeFormatString: "(%c%s)",
	},
}

var currencySymbols = map[string]rune{
	"EUR": '€',
	"USD": '$',
}

// header returns the translated header of the ledger
func header(locale string) string {
	t := localisations[locale]
	return fmt.Sprintf(headerFormatString, t.date, t.description, t.change)
}

// formatAmount returns the string representation for an amount for the given locale and currency
func formatAmount(amount int, locale, currency string) string {
	isNegative := amount < 0

	if isNegative {
		amount *= -1
	}

	major, minor := amount/100, amount%100
	majorString := strconv.Itoa(major)

	loc, _ := localisations[locale]
	tSep, dSep := loc.thousandsSeparator, loc.decimalSeparator

	var out strings.Builder

	majorLength := len(majorString)
	if majorLength%3 != 0 {
		out.WriteString(majorString[:majorLength%3])
	}

	for i := len(majorString) % 3; i < majorLength; i += 3 {
		if i != 0 {
			out.WriteByte(tSep)
		}
		out.WriteString(majorString[i : i+3])
	}

	out.WriteByte(dSep)
	out.WriteString(fmt.Sprintf("%02d", minor))

	converted := out.String()

	if isNegative {
		converted = fmt.Sprintf(loc.negativeFormatString, currencySymbols[currency], converted)
	} else {
		converted = fmt.Sprintf(loc.positiveFormatString, currencySymbols[currency], converted)
	}

	return converted
}

// newLedger returns a new sorted ledger from the given entries
func newLedger(entries []Entry) (ledger, error) {
	var entriesCopy ledger
	for _, e := range entries {
		t, err := time.Parse("2006-01-02", e.Date)

		if err != nil {
			return nil, fmt.Errorf("Invalid date format: %s", e.Date)
		}

		e.dateTime = t

		if len(e.Description) > 25 {
			e.Description = fmt.Sprintf("%s...", e.Description[:22])
		}

		entriesCopy = append(entriesCopy, e)
	}

	sort.Sort(entriesCopy)

	return entriesCopy, nil
}

// FormatLedger returns a string representation of the ledger
func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if _, ok := localisations[locale]; !ok {
		return "", fmt.Errorf("Invalid locale: %s", locale)
	}

	if _, ok := currencySymbols[currency]; !ok {
		return "", fmt.Errorf("Invalid currency: %s", currency)
	}

	if len(entries) == 0 {
		return header(locale), nil
	}

	l, err := newLedger(entries)

	if err != nil {
		return "", err
	}

	var ledgerBuilder strings.Builder
	ledgerBuilder.WriteString(header(locale))

	for _, entry := range l {
		date := entry.dateTime.Format(localisations[locale].dateFormat)
		amount := formatAmount(entry.Change, locale, currency)
		ledgerBuilder.WriteString(fmt.Sprintf(entryFormatString, date, entry.Description, amount))
	}

	return ledgerBuilder.String(), nil
}
