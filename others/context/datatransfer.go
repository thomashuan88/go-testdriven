package context

import (
	"context"
	"fmt"
)

func GetUser(ctx context.Context) {
	fmt.Println(ctx.Value("name"))
}
