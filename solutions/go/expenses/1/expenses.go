package expenses

import "errors"
// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var out []Record
    for _, r := range in {
        if (predicate(r)) {
            out = append(out, r)
        }
    }
    return out
    // panic("Please implement the Filter function")
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
    return func(r Record) bool {
        if r.Day >= p.From && r.Day <= p.To {
            return true
        }
        return false
    }
    // panic("Please implement the ByDaysPeriod function")
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
        if r.Category == c {
            return true
        }
        return false
    }
    // panic("Please implement the ByCategory function")
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var total float64 = 0
    for _, r := range in {
        if p.From <= r.Day && r.Day <= p.To {
            total += r.Amount
        }
    }
    return total
    // panic("Please implement the TotalByPeriod function")
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	var totalExpenses float64 = 0
    var found bool = false
    for _, r := range in {
        bc := ByCategory(c)
        bdp := ByDaysPeriod(p)
        if bc(r) {
            found = true
        }
        if bc(r) && bdp(r) {
            totalExpenses += r.Amount
        }
    }
    if !found {
        return 0, errors.New("unknown category " + c)
    }
    return totalExpenses, nil
    // panic("Please implement the CategoryExpenses function")
}
