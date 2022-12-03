package timeutil

import (
	"math"
	"sort"
	"time"
)

type TimePeriodCalculator struct {
	periods timePeriods
}

type Result struct {
	Years int
	Days  int
}

func NewTimePeriodCalculator() *TimePeriodCalculator {
	return &TimePeriodCalculator{}
}

func (c *TimePeriodCalculator) TimePeriod(s, e time.Time) {
	c.periods = append(c.periods, timePeriod{s: s, e: e})
}

func (c *TimePeriodCalculator) Calculate() Result {
	if len(c.periods) == 0 {
		return Result{}
	}

	// 1. sort
	c.sort()

	// 2. calculate
	var redundant float64
	for len(c.periods) >= 2 {
		var e time.Time
		if c.periods[0].e.After(c.periods[1].e) {
			// contain c.periods[1]
			e = c.periods[0].e
		} else if c.periods[0].e.After(c.periods[1].s) {
			// a partial repetition of the time period
			e = c.periods[1].e
		} else {
			// a discontinuous time period
			redundant += c.periods[1].s.Sub(c.periods[0].e).Hours()
			e = c.periods[1].e
		}

		c.periods[1].s = c.periods[0].s
		c.periods[1].e = e
		c.periods = c.periods[1:]
	}

	hours := c.periods[0].e.Sub(c.periods[0].s).Hours() - redundant
	days := math.Round(hours / 24)
	years := math.Round(days / 365)

	// 3. clear
	c.periods = timePeriods{}

	return Result{Days: int(days), Years: int(years)}
}

func (c *TimePeriodCalculator) sort() {
	sort.Sort(c.periods)
}

type timePeriods []timePeriod

type timePeriod struct {
	s, e time.Time
}

func (n timePeriods) Len() int           { return len(n) }
func (n timePeriods) Less(i, j int) bool { return n[i].s.Before(n[j].s) }
func (n timePeriods) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
