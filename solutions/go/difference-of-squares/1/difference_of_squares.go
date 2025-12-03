package diffsquares

func SquareOfSum(n int) (result int) {
	for i := 1; i <= n; i++ {
        result += i
    }
    result *= result
    return 
    // panic("Please implement the SquareOfSum function")
}

func SumOfSquares(n int) (result int) {
	for i := 1; i <= n; i++ {
        result += i * i
    }
    return
    // panic("Please implement the SumOfSquares function")
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
    // panic("Please implement the Difference function")
}
