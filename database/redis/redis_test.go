package redis

import (
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {

	//test()
	//a := []int{1,2,3}
	aa(1, 2)
}

func aa(v ...int) {
	fmt.Println(v[0])
}

func TestGet(t *testing.T) {
	_ = InitRedis()
	cmd := RDB.LRange(ctxRedis, "list", 0, 10)
	list := cmd.Val()
	fmt.Println(list)
}
