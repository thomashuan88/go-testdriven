package factory_test

import (
	"go-testdriven/ch11/factory"
	"testing"
)

func TestFactory(t *testing.T) {
	var stu = factory.NewStudent("huan", 11.5)
	t.Log("stu : ", stu)
}
