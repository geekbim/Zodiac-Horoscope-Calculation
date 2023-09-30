package calculation_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func FindZodiacByDate(date time.Time) string {
	zodiacs := []string{"Rat", "Ox", "Tiger", "Rabbit", "Dragon", "Snake", "Horse", "Goat", "Monkey", "Rooster", "Dog", "Pig"}

	year := date.Year()

	index := (year - 1912) % 12

	return zodiacs[index]
}

func FindHoroscopeByDate(date time.Time) string {
	day := date.Day()
	month := date.Month()

	horoscopes := map[string]struct {
		startMonth time.Month
		startDay   int
		endMonth   time.Month
		endDay     int
	}{
		"♈ Aries (Ram): March 21–April 19": {
			startMonth: time.March,
			startDay:   21,
			endMonth:   time.April,
			endDay:     19,
		},
		"♉ Taurus (Bull): April 20–May 20": {
			startMonth: time.April,
			startDay:   20,
			endMonth:   time.May,
			endDay:     20,
		},
		"♊ Gemini (Twins): May 21–June 21": {
			startMonth: time.May,
			startDay:   21,
			endMonth:   time.June,
			endDay:     21,
		},
		"♋ Cancer (Crab): June 22–July 22": {
			startMonth: time.June,
			startDay:   22,
			endMonth:   time.July,
			endDay:     22,
		},
		"♌ Leo (Lion): July 23–August 22": {
			startMonth: time.July,
			startDay:   23,
			endMonth:   time.August,
			endDay:     22,
		},
		"♍ Virgo (Virgin): August 23–September 22": {
			startMonth: time.August,
			startDay:   23,
			endMonth:   time.September,
			endDay:     22,
		},
		"♎ Libra (Balance): September 23–October 23": {
			startMonth: time.September,
			startDay:   23,
			endMonth:   time.October,
			endDay:     23,
		},
		"♏ Scorpius (Scorpion): October 24–November 21": {
			startMonth: time.October,
			startDay:   24,
			endMonth:   time.November,
			endDay:     21,
		},
		"♐ Sagittarius (Archer): November 22–December 21": {
			startMonth: time.November,
			startDay:   22,
			endMonth:   time.December,
			endDay:     21,
		},
		"♑ Capricornus (Goat): December 22–January 19": {
			startMonth: time.December,
			startDay:   22,
			endMonth:   time.January,
			endDay:     19,
		},
		"♒ Aquarius (Water Bearer): January 20–February 18": {
			startMonth: time.January,
			startDay:   20,
			endMonth:   time.February,
			endDay:     18,
		},
		"♓ Pisces (Fish): February 19–March 20": {
			startMonth: time.February,
			startDay:   19,
			endMonth:   time.March,
			endDay:     20,
		},
	}

	for i, horoscope := range horoscopes {
		if (month == horoscope.startMonth && day >= horoscope.startDay) ||
			(month == horoscope.endMonth && day <= horoscope.endDay) ||
			(month > horoscope.startMonth && month < horoscope.endMonth) {
			return i
		}
	}

	return "not found"
}

func TestZodiacCalculation(t *testing.T) {
	// 2023-04-15
	assert.Equal(t, "Rabbit", FindZodiacByDate(time.Date(2023, time.April, 15, 0, 0, 0, 0, time.UTC)))

	// 2014-08-07
	assert.Equal(t, "Horse", FindZodiacByDate(time.Date(2014, time.August, 07, 0, 0, 0, 0, time.UTC)))

	// 2007-12-23
	assert.Equal(t, "Pig", FindZodiacByDate(time.Date(2007, time.December, 23, 0, 0, 0, 0, time.UTC)))

	// 2000-02-05
	assert.Equal(t, "Dragon", FindZodiacByDate(time.Date(2000, time.February, 05, 0, 0, 0, 0, time.UTC)))
}

func TestHoroscopeCalculation(t *testing.T) {
	// 2023-04-15
	assert.Equal(t, "♈ Aries (Ram): March 21–April 19", FindHoroscopeByDate(time.Date(2023, time.April, 15, 0, 0, 0, 0, time.UTC)))

	// 2014-08-07
	assert.Equal(t, "♌ Leo (Lion): July 23–August 22", FindHoroscopeByDate(time.Date(2014, time.August, 07, 0, 0, 0, 0, time.UTC)))

	// 2007-12-23
	assert.Equal(t, "♑ Capricornus (Goat): December 22–January 19", FindHoroscopeByDate(time.Date(2007, time.December, 23, 0, 0, 0, 0, time.UTC)))

	// 2000-02-05
	assert.Equal(t, "♒ Aquarius (Water Bearer): January 20–February 18", FindHoroscopeByDate(time.Date(2000, time.February, 05, 0, 0, 0, 0, time.UTC)))
}
