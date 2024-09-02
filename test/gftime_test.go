package test

import (
	"testing"

	"github.com/gogf/gf/v2/os/gtime"
)

func TestTime(t *testing.T) {
	now := gtime.Now()
	t.Log(now)

	t.Log(now.StartOfDay())
	t.Log(now.EndOfDay())
}
