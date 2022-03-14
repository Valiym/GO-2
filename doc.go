// Package doc implements functions to demonstrate godoc features
//
// The ErrorWithPanic при выполнении этой функции происходит паника
//
// func ErrorWithPanic() (err error) {
//	defer func() {
//		if v := recover(); v != nil {
//			err = New(fmt.Sprintf("detected panic error: %s", v))
//		}
//	}()
//
// So, if call this function, you get panic
package doc

import (
	"fmt"
	"time"
)

// ErrorWithTime добавляет текущее время к ошибке
type ErrorWithTime struct {
	text string
	time time.Time
}

func New(text string) error {
	return &ErrorWithTime{
		text: text,
		time: time.Now(),
	}
}

func (e *ErrorWithTime) Error() string {
	return fmt.Sprintf("error: %s\n, %v", e.text, e.time)
}

func main() {
	err := ErrorWithPanic()
	if err != nil {
		fmt.Println(err)
	}
}

// ErrorWithPanic при выполнении этой функции происходит паника
func ErrorWithPanic() (err error) {
	defer func() {
		if v := recover(); v != nil {
			err = New(fmt.Sprintf("detected panic error: %s", v))
		}
	}()

	arr := []int{1, 2}

	for i := 0; i <= 5; i++ {
		fmt.Printf("%d %d\n", i, arr[i])
	}
	return
}
