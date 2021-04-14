package library 

import (
	"time"
)

func StringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func Age(birthdate, today time.Time) int {
	today = today.In(birthdate.Location())
	ty, tm, td := today.Date()
	today = time.Date(ty, tm, td, 0, 0, 0, 0, time.UTC)
	by, bm, bd := birthdate.Date()
	birthdate = time.Date(by, bm, bd, 0, 0, 0, 0, time.UTC)
	if today.Before(birthdate) {
		return 0
	}
	age := ty - by
	anniversary := birthdate.AddDate(age, 0, 0)
	if anniversary.After(today) {
		age--
	}
	return age
}

func Diffday(fromDate, toDate time.Time) int {
	yfrom, mfrom, dfrom := fromDate.Date()
	yto, mto, dto := toDate.Date()
	
	from := theDate(yfrom, int(mfrom), dfrom)
    to := theDate(yto, int(mto), dto)
    days := to.Sub(from).Hours() / 24
	
	return int(days)
}

func theDate(year, month, day int) time.Time {
    return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}