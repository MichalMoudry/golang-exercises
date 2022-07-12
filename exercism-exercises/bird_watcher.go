package main

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var result int
	for i := 0; i < len(birdsPerDay); i++ {
		result += birdsPerDay[i]
	}
	return result
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	return TotalBirdCount(birdsPerDay[(week-1)*7 : week*7])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if (i+1)%2 != 0 {
			birdsPerDay[i] += 1
		}
	}
	return birdsPerDay
}

func main() {
	birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	println("Total bird count (expected - 34):", TotalBirdCount(birdsPerDay))

	println("Birds in week (expected - 12):", BirdsInWeek(birdsPerDay, 2))
	println("Birds in week (expected - 12):", BirdsInWeek([]int{3, 0, 3, 3, 2, 1, 0}, 1))
	println("Birds in week (expected - 14):", BirdsInWeek([]int{3, 0, 5, 1, 0, 4, 1, 0, 3, 4, 3, 0, 8, 0}, 1))

	birdsPerDayMistake := []int{2, 5, 0, 7, 4, 1}
	correctedBirdsPerDay := FixBirdCountLog(birdsPerDayMistake)
	print("Fix bird count log (expected - [3 5 1 7 5 1]): ")
	print("|")
	for i := 0; i < len(correctedBirdsPerDay); i++ {
		print(correctedBirdsPerDay[i], "|")
	}
}
