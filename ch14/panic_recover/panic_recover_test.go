package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	// defer func() {
	// 	fmt.Println("Finnally!")
	// }()

	// use recover() , then program wouldn't stop by panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recoved from ", err)
		}
	}()
	fmt.Println("start")
	panic(errors.New("Something wrong!"))
	//os.Exit(-1) // this will exits without defer
}
