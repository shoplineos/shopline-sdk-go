package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildTimestamp(t *testing.T) {
	timestamp := BuildTimestamp()

	//fmt.Printf("timestamp: %s\n", timestamp)
	a := assert.New(t)
	a.True(len(timestamp) == 13)
}
