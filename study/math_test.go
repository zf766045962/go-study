package study

import (
	"math/rand"
	"testing"
)

//单元测试
func TestCalNumber(t *testing.T) {

	result := CalNumber(1, 2)
	if result != 3 {
		t.Errorf("wrong: result=%d actual=%d", result, 3)
	}

}

// 性能测试
func BenchmarkCalNumber(b *testing.B) {
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		CalNumber(i+1, rand.Intn(i+1))
	}
}
