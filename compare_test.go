package set_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MovieStoreGuy/set"
)

func TestEqual(t *testing.T) {
	t.Parallel()

	values := []struct{ a, b set.Set }{
		{nil, nil},
		{set.New('a', 'b', 'c'), set.New('a', 'b', 'c')},
		{set.New('a', 'a', 'a', 'a', 'a', 'a', 'a'), set.New('a')},
	}

	for _, v := range values {
		assert.True(t, set.Equal(v.a, v.b))
	}
}

func TestContains(t *testing.T) {
	t.Parallel()

	cases := []struct {
		a, b set.Set
		msg  string
	}{
		{nil, nil, "comparing empty sets"},
		{nil, set.New(), "An empty set is contains within a set"},
		{set.New(), nil, "An empty set is equal to an empty set"},
		{set.New('a'), set.New('a', 'b'), "Sub set contains within b"},
	}

	for _, c := range cases {
		assert.True(t, set.Contains(c.a, c.b), c.msg)
	}
}

func TestUnion(t *testing.T) {
	t.Parallel()

	values := []struct{ a, b set.Set }{
		{nil, nil},
		{nil, set.New()},
		{set.New(), nil},
		{nil, set.New('a')},
		{set.New('a'), nil},
		{set.New('a'), set.New('a')},
		{set.New('a'), set.New('b')},
	}

	for _, v := range values {
		u := set.Union(v.a, v.b)
		assert.True(t, set.Contains(v.a, u), "Set a must be contained within U")
		assert.True(t, set.Contains(v.b, u), "Set b must be contained within U")
	}
}

func TestIntersection(t *testing.T) {
	t.Parallel()
	values := []struct{ a, b set.Set }{
		{nil, nil},
		{nil, set.New('a')},
		{set.New('a'), nil},
		{set.New('a'), set.New('a')},
		{set.New('a'), set.New('b')},
	}

	for _, v := range values {
		u := set.Union(v.a, v.b)
		assert.True(t, set.Equal(set.Intersection(v.a, u), v.a))
		assert.True(t, set.Equal(set.Intersection(v.b, u), v.b))
	}
}

func TestSubtraction(t *testing.T) {
	t.Parallel()
	values := []struct{ a, b set.Set }{
		{nil, nil},
		{nil, set.New('a')},
		{set.New('a'), nil},
		{set.New('a'), set.New('b')},
	}

	for _, v := range values {
		u := set.Union(v.a, v.b)
		assert.True(t, set.Equal(set.Subtract(v.a, u), v.b))
		assert.True(t, set.Equal(set.Subtract(v.b, u), v.a))
	}
}
