// Copyright © 2018 tzx All rights reserved.

package unsafe

import (
	"testing"

	. "gopkg.in/check.v1"
)

type MySuite struct{}

var _ = Suite(&MySuite{})

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func (s *MySuite) Test_Suite1(c *C) {
	var x = []byte("Hello World!")
	var y = String(x)
	var z = string(x)

	// c.Check(y, Not(Equals), z)
	c.Check(y, Equals, z)
}

func (s *MySuite) Test_Suite2(c *C) {
	var x = "Hello World!"
	var y = Bytes(x)
	var z = []byte(x)

	c.Check(y, DeepEquals, z)
}

func (s *MySuite) BenchmarkLogic(c *C) {
	// sa := []interface{}{"q", "w", "e", "r", "t"}
	// sb := []interface{}{"q", "w", "a", "s", "z", "x"}
	var x = "Hello World!"
	c.ResetTimer()
	for i := 0; i < c.N; i++ {
		// Logic to benchmark
		// _ = SliceEqual(sa, sb)
		// c.Check(sa, Not(DeepEquals), sb)
		_ = StringByte(x)
	}
}

func Test_String(t *testing.T) {
	var x = []byte("Hello World!")
	var y = String(x)
	var z = string(x)

	if y != z {
		t.Fail()
	}
}

func Test_Bytes(t *testing.T) {
	var x = "Hello World!"
	var y = Bytes(x)
	var z = []byte(x)

	if !SliceReflectEqual(y, z) {
		t.Fail()
	}
}

func Test_ByteString(t *testing.T) {
	var x = []byte("Hello World!")
	var y = ByteString(x)
	var z = string(x)

	if y != z {
		t.Fail()
	}
}

func Test_StringByte(t *testing.T) {
	var x = "Hello World!"
	var y = StringByte(x)
	var z = []byte(x)

	u := make([]interface{}, len(y))
	v := make([]interface{}, len(z))

	for i, w := range y {
		u[i] = w
	}
	for i, w := range z {
		v[i] = w
	}
	if !SliceEqual(u, v) {
		t.Fail()
	}
}

func Benchmark_Normal1(b *testing.B) {
	var x = []byte("Hello World!")
	for i := 0; i < b.N; i++ {
		_ = string(x)
	}
}

func Benchmark_Normal2(b *testing.B) {
	var x = "Hello World!"
	for i := 0; i < b.N; i++ {
		_ = []byte(x)
	}
}

func Benchmark_String(b *testing.B) {
	var x = []byte("Hello World!")
	for i := 0; i < b.N; i++ {
		_ = String(x)
	}
}

func Benchmark_Bytes(b *testing.B) {
	var x = "Hello World!"
	for i := 0; i < b.N; i++ {
		_ = Bytes(x)
	}
}

func Benchmark_StringByte(b *testing.B) {
	var x = "Hello World!"
	for i := 0; i < b.N; i++ {
		_ = StringByte(x)
	}
}

func Benchmark_ByteString(b *testing.B) {
	var x = []byte("Hello World!")
	for i := 0; i < b.N; i++ {
		_ = ByteString(x)
	}
}

func BenchmarkSliceEqual(b *testing.B) {
	sa := []interface{}{"q", "w", "e", "r", "t"}
	sb := []interface{}{"q", "w", "a", "s", "z", "x"}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = SliceEqual(sa, sb)
	}
}

func BenchmarkSliceEqualBCE(b *testing.B) {
	sa := []interface{}{"q", "w", "e", "r", "t"}
	sb := []interface{}{"q", "w", "a", "s", "z", "x"}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = SliceEqualBCE(sa, sb)
	}
}
func BenchmarkDeepEqual(b *testing.B) {
	sa := []string{"q", "w", "e", "r", "t"}
	sb := []string{"q", "w", "a", "s", "z", "x"}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = SliceReflectEqual(sa, sb)
	}
}
