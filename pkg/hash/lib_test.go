package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash11(t *testing.T) {
	num_bins := 10
	num_samples := 1000
	bins := make([]int, num_bins)
	for i := 0; i < num_samples; i++ {
		bin_index := int(Hash11(float32(i)) * float32(num_bins))
		bins[bin_index]++
	}
	expected := []int{98, 111, 91, 85, 105, 100, 103, 100, 116, 91}
	assert.Equal(t, bins, expected)
}
