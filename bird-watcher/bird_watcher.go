package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	totalBirdCount := 0

	for index := 0; index < len(birdsPerDay); index++ {
		totalBirdCount += birdsPerDay[index]
	}
	return totalBirdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	lastIndex := week * 7
	firstIndex := lastIndex - 7
	birdsInWeek := 0
	for index := firstIndex; index < lastIndex; index++ {
		birdsInWeek += birdsPerDay[index]
	}
	return birdsInWeek
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for index := 0; index < len(birdsPerDay); index++ {
		isEven := index % 2

		if isEven == 0 {
			birdsPerDay[index] += 1
		}
	}

	return birdsPerDay
}
