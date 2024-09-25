package context

import (
	"context"
	"fmt"
)

type UserInfo struct {
	Name string
}

func GetUser(ctx context.Context) {
	fmt.Println(ctx.Value("name").(UserInfo).Name)
}
