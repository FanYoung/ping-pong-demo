package server

import (
	"testing"
	"time"
)

func Test_Time(t *testing.T) {
	a := time.Now()
	b := time.Now()
	t.Log(a)
	t.Log(b)

}
