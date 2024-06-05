package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var interestRate float32 = 2.475
	if balance < 0 {
		interestRate = 3.213
	} else if balance >= 0 && balance < 1000 {
		interestRate = 0.5
	} else if balance >= 1000 && balance < 5000 {
		interestRate = 1.621
	}
	return interestRate
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	var interestRate float32 = InterestRate(balance) / 100
	return float64(interestRate) * balance
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var years int = 0
	var balanceWithInterest float64 = balance
	for {
		if balanceWithInterest >= targetBalance {
			break
		} else {
			balanceWithInterest += Interest(balanceWithInterest)
			years += 1
		}
	}
	return years
}

func Run() {
	println("Interest rate (expected - 0.5):", InterestRate(200.75))
	println("Interest (expected - 1.003750):", Interest(200.75))
	println("Annual balance update (expected - 201.75375)", AnnualBalanceUpdate(200.75))
	println("Years before reaching the desired balance (expected - 14):", YearsBeforeDesiredBalance(200.75, 214.88))
}
