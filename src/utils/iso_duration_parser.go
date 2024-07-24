package utils

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

func ParseISODuration(duration string) (time.Duration, error) {
	pattern := `^P(?:(\d+)Y)?(?:(\d+)M)?(?:(\d+)D)?(?:T(?:(\d+)H)?(?:(\d+)M)?(?:(\d+)S)?)?$`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(duration)

	if matches == nil {
		return 0, errors.New("invalid ISO 8601 duration format")
	}

	var totalDuration time.Duration

	// Parse years, months, days, hours, minutes, and seconds
	if matches[1] != "" {
		years, _ := strconv.Atoi(matches[1])
		totalDuration += time.Duration(years*8760) * time.Hour // Approximation: 1 year = 8760 hours
	}
	if matches[2] != "" {
		months, _ := strconv.Atoi(matches[2])
		totalDuration += time.Duration(months*730) * time.Hour // Approximation: 1 month = 730 hours
	}
	if matches[3] != "" {
		days, _ := strconv.Atoi(matches[3])
		totalDuration += time.Duration(days*24) * time.Hour
	}
	if matches[4] != "" {
		hours, _ := strconv.Atoi(matches[4])
		totalDuration += time.Duration(hours) * time.Hour
	}
	if matches[5] != "" {
		minutes, _ := strconv.Atoi(matches[5])
		totalDuration += time.Duration(minutes) * time.Minute
	}
	if matches[6] != "" {
		seconds, _ := strconv.Atoi(matches[6])
		totalDuration += time.Duration(seconds) * time.Second
	}

	return totalDuration, nil
}
