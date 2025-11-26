package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var result float64
    result = float64(productionRate) * (successRate / 100 )
    return result
    
    //panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	result := (float64(productionRate) * (successRate / 100)) / 60
    return int(result)
    
    //panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupOfTen := carsCount / 10
    otherCars := carsCount % 10
    return uint(groupOfTen * 95000 + otherCars * 10000)
    //panic("CalculateCost not implemented")
}
