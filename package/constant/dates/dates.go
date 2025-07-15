package dates

const DaysInWeek int = 7

func WeeksWoDays(weeks int) int {
	return weeks * DaysInWeek
}

func DaysToWeeks(days int) float64 {
	return float64(days) / float64(DaysInWeek)
}
