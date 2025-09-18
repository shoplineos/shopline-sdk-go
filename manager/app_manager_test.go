package manager

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

var m sync.Map

func TestSyncMap(t *testing.T) {
	m.Store("name", 1)

	fmt.Println(m.Load("name"))

	value, _ := m.Load("name")
	assert.Equal(t, 1, value)
}
