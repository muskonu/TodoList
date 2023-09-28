package globalkey

import "fmt"

var recurrenceMap = map[int64]string{0: "once", 1: "daily", 2: "weekly", 3: "monthly", 4: "annually"}

func Email(studentEmail string) string {
	return fmt.Sprintf("email:%s", studentEmail)
}

func Recurrence(recurrence int64) string {
	return recurrenceMap[recurrence]
}
