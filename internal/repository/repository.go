package repository

import (
	"time"
)

var costs = make(map[time.Time]map[string]int)

func AddCost(date time.Time, category string, sum int) error {
	if _, ok := costs[date]; ok {
		if _, ko := costs[date][category]; ko {
			costs[date][category] += sum
		} else {
			costs[date][category] = sum
		}
	} else {
		costs[date] = map[string]int{category: sum}
	}
	return nil
}

func GetCostsForPeriod(begin, end time.Time) (map[string]int, error) {
	result := map[string]int{}
	for date, spending := range costs {
		if (date.After(begin) || date.Equal(begin)) && (date.Before(end) || date.Equal(end)) {
			for category, sum := range spending {
				if _, ok := result[category]; ok {
					result[category] += sum
				} else {
					result[category] = sum
				}
			}
		}
	}
	return result, nil
}
