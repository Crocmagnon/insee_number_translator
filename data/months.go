package data

import "time"

var frenchMonths = []string{
	"janvier",
	"février",
	"mars",
	"avril",
	"mai",
	"juin",
	"juillet",
	"août",
	"septembre",
	"octobre",
	"novembre",
	"décembre",
}

func frenchMonth(month time.Month) string {
	return frenchMonths[month-1]
}
