package main

import (
	"errors"
	"fmt"
	"time"
)

type Countdown struct {
	Days    int
	Hours   int
	Minutes int
}

func calculateCountdown(targetDate string) (Countdown, error) {
	layout := "2006-01-02 15:04:05"
	targetTime, err := time.Parse(layout, targetDate)
	if err != nil {
		// If parsing fails, return the error.
		return Countdown{}, err
	}
	timeNow := time.Now()
	if targetTime.Before(timeNow) {
		return Countdown{}, errors.New("target date is in the past")
	}
	duration := targetTime.Sub(timeNow)
	totalHours := int(duration.Hours())
	days := totalHours / 24
	remainingHours := totalHours % 24

	totalMinutes := int(duration.Minutes())
	remainingMinutes := totalMinutes % 60
	countdown := Countdown{
		Days:    days,
		Hours:   remainingHours,
		Minutes: remainingMinutes,
	}
	return countdown, nil
}

func main() {
	target := "2025-10-07 12:30:00"
	countdown, err := calculateCountdown(target)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Time until %s: %d days, %d hours, %d minutes\n", target, countdown.Days, countdown.Hours, countdown.Minutes)
	}

	pastTarget := "2024-01-01 00:00:00"
	_, err = calculateCountdown(pastTarget)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
