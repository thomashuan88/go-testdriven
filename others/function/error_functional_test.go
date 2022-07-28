package function

import (
	"errors"
	"testing"
)

// type appHandler func(writer http.ResponseWriter, request *http.Request) error

// func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		err := handler(w, r)
// 		if err != nil {
// 			switch {
// 			case os.IsNotExist(err):
// 				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 			}
// 		}
// 	}
// }

type appHandlerT func(aa int) error

func aWrap(ah appHandlerT) func(int) int {
	return func(a int) int {
		err := ah(a)
		if err != nil {
			return 0
		}
		return a + 5
	}
}

func tt(a int) error {
	if a != 5 {
		return errors.New("not 5!")
	}
	return nil
}
func TestFilefunc(t *testing.T) {
	aa := aWrap(tt)
	t.Log("check: ", aa(4))
}
