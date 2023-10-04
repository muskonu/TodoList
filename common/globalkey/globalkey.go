package globalkey

import (
	"fmt"
	"time"
)

var recurrenceMap = map[int64]string{0: "once", 1: "daily", 2: "weekly", 3: "monthly", 4: "annually"}

func Email(studentEmail string) string {
	return fmt.Sprintf("email:%s", studentEmail)
}

func Recurrence(recurrence int64) string {
	return recurrenceMap[recurrence]
}

func Period(recurrence int64) time.Duration {
	switch recurrence {
	case 1:
		return 24 * time.Hour
	case 2:
		return 7 * 24 * time.Hour
	case 3:
		return getYearMonthToDay(time.Now()) * 24 * time.Hour
	case 4:
		return getYearToDay(time.Now()) * 24 * time.Hour
	default:
		return 0
	}
}

func getYearToDay(now time.Time) time.Duration {
	year := now.Year()
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return 366
	}
	return 365
}

func getYearMonthToDay(now time.Time) time.Duration {
	year, month := now.Year(), int(now.Month())
	// 有31天的月份
	day31 := map[int]struct{}{
		1:  {},
		3:  {},
		5:  {},
		7:  {},
		8:  {},
		10: {},
		12: {},
	}
	if _, ok := day31[month]; ok {
		return 31
	}
	// 有30天的月份
	day30 := map[int]struct{}{
		4:  {},
		6:  {},
		9:  {},
		11: {},
	}
	if _, ok := day30[month]; ok {
		return 30
	}
	// 计算是平年还是闰年
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		// 得出2月的天数
		return 29
	}
	// 得出2月的天数
	return 28
}
