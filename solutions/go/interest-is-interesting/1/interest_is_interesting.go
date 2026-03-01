package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
        case balance < 0: return 3.213
        case balance >= 0 && balance < 1000: return 0.5
        case balance >= 1000 && balance < 5000: return 1.621
        case balance >= 5000: return 2.475
    }
    return 0
    
    // panic("Please implement the InterestRate function")
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	rate := float64(InterestRate(balance))
    result := (balance / 100.0) * rate
    return result
    // panic("Please implement the Interest function")
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	rate := float64(InterestRate(balance))
    result := (balance / 100.0) * rate
    return balance + result
    // panic("Please implement the AnnualBalanceUpdate function")
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	currentYears := 0
    if balance < targetBalance {
        for currentYears = 0; balance < targetBalance; currentYears++ {
            balance = AnnualBalanceUpdate(balance)
        } 
    }
    return currentYears
    // panic("Please implement the YearsBeforeDesiredBalance function")
}
