package inverted

import (
	"math/rand"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestContainer(t *testing.T) {
	TestCases := map[string]struct {
		push   []int64
		pop    []int64
		expect []int64
	}{
		"normal": {
			push:   []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			pop:    []int64{0, 2, 4, 6, 9},
			expect: []int64{1, 3, 5, 7, 8},
		},
		"repetitive": {
			push:   []int64{0, 0, 1, 2, 1, 4, 4, 5, 4, 4, 3, 5, 6, 7, 3, 2, 3, 5, 6, 0},
			pop:    []int64{0, 4, 2, 2, 10, 7, 12},
			expect: []int64{1, 5, 3, 6},
		},
	}

	for name, tcase := range TestCases {
		t.Run(name, func(t *testing.T) {
			c := NewIDContainer()
			for _, v := range tcase.push {
				c.Push(v)
			}
			for _, v := range tcase.pop {
				c.Pop(v)
			}
			got := c.Array()
			if !cmp.Equal(tcase.expect, got) {
				t.Errorf("unintended ： %s", cmp.Diff(tcase.expect, got))
			}
		})
	}
}

func TestIndexIndex(t *testing.T) {
	type cases struct {
		key    string
		push   []int64
		pop    []int64
		expect []int64
	}

	TestCases := map[string][]cases{
		"normal": {
			cases{
				key:    "first",
				push:   []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				pop:    []int64{1, 2, 3, 4, 0},
				expect: []int64{5, 6, 7, 8, 9},
			},
			cases{
				key:    "second",
				push:   []int64{1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3},
				pop:    []int64{0, 2, 4},
				expect: []int64{1, 3},
			},
		},
	}

	for name, tcases := range TestCases {
		t.Run(name, func(t *testing.T) {
			index := NewIndex()
			for _, tce := range tcases {
				for _, v := range tce.push {
					index.Push(tce.key, v)
				}
			}
			for _, tce := range tcases {
				for _, v := range tce.pop {
					index.Pop(tce.key, v)
				}
			}
			for _, tce := range tcases {
				got := index.GetAllID(tce.key)
				if !cmp.Equal(tce.expect, got) {
					t.Errorf("unintended : %s", cmp.Diff(tce.expect, got))
				}
			}
			got := index.GetAllID(name)
			if got != nil {
				t.Errorf("unintended : %s", cmp.Diff(nil, got))
			}
		})
	}
}

func BenchmarkContainer(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	c := NewIDContainer()
	// 重置计时器
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Push(r.Int63())
		}
	})
}

func BenchmarkIndex(b *testing.B) {
	keywords := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	l := len(keywords)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := NewIndex()
	// 重置计时器
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			v := r.Int63()
			i.Push(keywords[int(v)%l], v)
		}
	})
}
