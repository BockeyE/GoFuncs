package captcha

import (
	"testing"
)

func TestRandom(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(Random(0, 1))
		//fmt.Println(Random(0, 1))
	}
}
