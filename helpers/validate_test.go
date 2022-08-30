package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------> A test for IsItEmpty function
func TestIsItEmpty(t *testing.T) {
	data1, data2, data3, data4, data5 := "apple", "974#!0726AD", " ", "w", ""

	assert.Equal(t, false, IsItEmpty(data1))
	assert.Equal(t, false, IsItEmpty(data2))
	assert.Equal(t, false, IsItEmpty(data3))
	assert.Equal(t, false, IsItEmpty(data4))

	assert.Equal(t, true, IsItEmpty(data5))

}
