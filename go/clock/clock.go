package clock

import "fmt"

const testVersion = 4

// Clock API as stub definitions.  No, it doesn't compile yet.
// More details and hints are in clock_test.go.

type Clock struct {
	hour int
	minute int
}

func New(hour, minute int)  {
	h := hour
	m := minute
	fmt.Println(h)
	fmt.Println(m)

	Clock.String(time string) string {

	}

}

func (Clock) String() string {
}


//
//func (Clock) Add(minutes int) Clock {
//}

// Remember to delete all of the stub comments.
// They are just noise, and reviewers will complain.
