// others/context/datatransfer_test.go

package context

import (
	"context"
	"testing"
)

func TestGetUser(t *testing.T) {
	// Create a test context
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", "thomas")

	// Call the GetUser function
	// Assuming it returns a user or an error
	GetUser(ctx)

	// Check if the function returned an error
	// if err != nil {
	// 	t.Errorf("GetUser returned an error: %v", err)
	// }

	// Check if the user is not nil
	// if user == nil {
	// 	t.Errorf("GetUser returned a nil user")
	// }
}
