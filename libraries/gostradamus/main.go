package main

import "github.com/bykof/gostradamus"

func main() {
	// Easy parsing
	dateTime, err := gostradamus.Parse("14.07.2017 02:40:00", "DD.MM.YYYY HH:mm:ss")
	if err != nil {
		panic(err)
	}

	// Easy manipulation
	dateTime = dateTime.ShiftMonths(-5).ShiftDays(2)

	// Easy formatting
	println(dateTime.Format("DD.MM.YYYY HH:mm:ss"))
	// 16.02.2017 02:40:00

	// Easy helper functions
	start, end := dateTime.SpanWeek()

	println(start.String(), end.String())
	// 2017-02-13T00:00:00.000000Z 2017-02-19T23:59:59.999999Z
}
