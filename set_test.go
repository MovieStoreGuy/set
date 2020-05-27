package set_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/MovieStoreGuy/set"
)

func TestHasItems(t *testing.T) {
	t.Parallel()

	assert.False(t, set.New().HasItem(nil), "can not have an entry for nil with a new set")
	s := set.New()
	s.Insert(nil)
	assert.True(t, s.HasItem(nil))
	assert.Equal(t, []interface{}{nil}, s.Items())
}

func TestFromValues(t *testing.T) {
	t.Parallel()
	values := []interface{}{1, 3, 3, 7}
	s := set.New(values...)
	for _, i := range values {
		assert.True(t, s.HasItem(i))
	}
	assert.Equal(t, 3, len(s.Items()))
}

func TestInsert(t *testing.T) {
	inserts := []interface{}{
		1,
		1.8,
		1 + 10i,
		"oh damn",
		struct{}{},
		new(struct{}),
		set.New('a', 'b'),
	}
	require.NotPanics(t, func() {
		set.New(inserts...)
	})
	require.Panics(t, func() {
		set.New([]interface{}{nil})
	})
}
