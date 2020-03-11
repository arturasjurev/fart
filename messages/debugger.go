package messages

import (
	"errors"
	"fmt"
	"os"
)

type Debugger struct {
	ON bool
}

func (d *Debugger) Println(a ...interface{}) {
	if d.ON {
		fmt.Println(a...)
	}
}

func (d *Debugger) Printf(format string, a ...interface{}) {
	if d.ON {
		fmt.Printf(format, a...)
	}
}

func (d *Debugger) UnwrapAndPrint(err error) {
	fmt.Println("Error:", errors.Unwrap(err))

	os.Exit(1)
}
