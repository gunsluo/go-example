package timeutil

import (
	"testing"
	"time"
)

func TestTimePeriodCalculator(t *testing.T) {
	var cases = []struct {
		periods []struct {
			s, e time.Time
		}
		expect Result
	}{
		{
			periods: []struct {
				s, e time.Time
			}{},
			expect: Result{},
		},
		{
			periods: []struct {
				s, e time.Time
			}{
				{
					s: time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC),
					e: time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			expect: Result{Years: 2, Days: 730},
		},
		{
			periods: []struct {
				s, e time.Time
			}{
				{
					s: time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
					e: time.Date(2011, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					s: time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
					e: time.Date(2011, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					s: time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
					e: time.Date(2011, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			expect: Result{Years: 1, Days: 365},
		},
		{
			periods: []struct {
				s, e time.Time
			}{
				{
					s: time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
					e: time.Date(2013, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					s: time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC),
					e: time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			expect: Result{Years: 5, Days: 1826},
		},
		{
			periods: []struct {
				s, e time.Time
			}{
				{
					s: time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
					e: time.Date(2013, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					s: time.Date(2011, 1, 1, 0, 0, 0, 0, time.UTC),
					e: time.Date(2011, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			expect: Result{Years: 3, Days: 1096},
		},

		{
			periods: []struct {
				s, e time.Time
			}{
				{
					s: time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC),
					e: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					s: time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
					e: time.Date(2013, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			expect: Result{Years: 6, Days: 2192},
		},
		{
			periods: []struct {
				s, e time.Time
			}{
				{
					s: time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
					e: time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					s: time.Date(2012, 2, 1, 0, 0, 0, 0, time.UTC),
					e: time.Date(2013, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					s: time.Date(2012, 2, 1, 0, 0, 0, 0, time.UTC),
					e: time.Date(2013, 2, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					s: time.Date(2014, 1, 1, 0, 0, 0, 0, time.UTC),
					e: time.Date(2015, 2, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			expect: Result{Years: 4, Days: 1492},
		},
	}

	for _, c := range cases {
		tc := NewTimePeriodCalculator()
		for _, p := range c.periods {
			tc.TimePeriod(p.s, p.e)
		}

		r := tc.Calculate()
		if r.Days != c.expect.Days {
			t.Fatalf("days - expect: %d, got %d ", c.expect.Days, r.Days)
		}

		if r.Years != c.expect.Years {
			t.Fatalf("years - expect: %d, got %d ", c.expect.Years, r.Years)
		}
	}
}
