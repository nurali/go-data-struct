package singly_test

import (
	"testing"

	"github.com/nurali/go-data-struct/linklist/singly"
)

func TestInsert(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		s := singly.NewSingly()

		// test
		s.Insert(10)

		res := s.ToArray()
		checkEqual(t, []int{10}, res)
	})

	t.Run("multiple", func(t *testing.T) {
		s := singly.NewSingly()

		// test
		s.Insert(10)
		s.Insert(20)
		s.Insert(30)

		checkEqual(t, []int{10, 20, 30}, s.ToArray())
	})

	t.Run("duplicate", func(t *testing.T) {
		s := singly.NewSingly()

		// test
		s.Insert(10)
		s.Insert(10)

		checkEqual(t, []int{10, 10}, s.ToArray())
	})
}

func TestDelete(t *testing.T) {
	t.Run("unique", func(t *testing.T) {
		s := singly.NewSingly(10, 20, 30)

		// test
		s.Delete(20)

		checkEqual(t, []int{10, 30}, s.ToArray())
	})

	t.Run("first", func(t *testing.T) {
		s := singly.NewSingly(10)

		// test
		s.Delete(10)

		checkEqual(t, []int{}, s.ToArray())
	})

	t.Run("last", func(t *testing.T) {
		s := singly.NewSingly(10, 20, 30)

		// test
		s.Delete(30)

		checkEqual(t, []int{10, 20}, s.ToArray())
	})

	t.Run("duplicate", func(t *testing.T) {
		s := singly.NewSingly(10, 10, 10)

		// test
		s.Delete(10)

		checkEqual(t, []int{10, 10}, s.ToArray())
	})

	t.Run("multiplle_duplicate", func(t *testing.T) {
		s := singly.NewSingly(10, 10, 10)

		// test
		s.Delete(10)
		checkEqual(t, []int{10, 10}, s.ToArray())

		// test
		s.Delete(10)
		checkEqual(t, []int{10}, s.ToArray())

		// test
		s.Delete(10)
		checkEqual(t, []int{}, s.ToArray())
	})
}

func checkEqual(t *testing.T, wantArr []int, gotArr []int) {
	t.Helper()

	if len(wantArr) != len(gotArr) {
		t.Errorf("array not equal as their length is different, want:%d, got:%d", len(wantArr), len(gotArr))
	}

	for ind := range wantArr {
		if wantArr[ind] != gotArr[ind] {
			t.Errorf("array not equal as their element is different, want:%d, got%d", wantArr[ind], gotArr[ind])
		}
	}
}
