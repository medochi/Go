package main

import "fmt"

const (
	Admin = 1
	Manger
	Contractor = 2
	Member     = 3
	Guest      = 4
)
const (
	Friday    = 0
	Saturaday = 1
	Sunday    = 2
	Monday    = 3
	Tuseday   = 4
	Wednsday  = 5
	Thursday  = 6
)

func accessGranted(daylol int) {
	fmt.Println("You can access the app, Enjoy")
}
func accessDenied(user int) {
	if user == Contractor {
		fmt.Println("You can't access today Your days are weekends ")
	} else if user == Member {
		fmt.Println("You can't access today Your days are weekdays ")
	} else if user == Guest {
		fmt.Println("You can't access today Your days are Mondays, Wednesdays and Fridays ")
	}
}
func checker(usr int, day int) {
	if usr == Manger || usr == Admin {
		accessGranted(day)
	} else if usr == Contractor && (day == Friday || day == Saturaday) {
		accessGranted(day)
	} else if usr == Member && (day != Friday && day != Saturaday) {
		accessGranted(day)
	} else if usr == Guest && ((day == Friday || day == Wednsday) || (day == Monday)) {
		accessGranted(day)
	} else {
		accessDenied(usr)
	}
}

func main() {
	rule, today := Contractor, Friday
	checker(rule, today)
}
