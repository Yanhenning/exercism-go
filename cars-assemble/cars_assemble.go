package cars

const OneHour float64 = 60
const DiscountPrice int = 95000
const NormalPrice int = 10000
const BulkQuantity int = 10

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	percentRate := successRate / 100
	return float64(productionRate) * percentRate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	percentRate := successRate / 100
	productionRatePerHour := float64(productionRate) * percentRate
	return int(productionRatePerHour / OneHour)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	decimals := carsCount / BulkQuantity
	remaining := carsCount % BulkQuantity

	cost := decimals*DiscountPrice + remaining*NormalPrice
	return uint(cost)
}
