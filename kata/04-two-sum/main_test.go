package kata

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    [2]int
	}{
		{
			name:    "basic test",
			numbers: []int{1, 2, 3},
			target:  4,
			want:    [2]int{0, 2},
		},
		{
			name:    "backwards test",
			numbers: []int{3, 1},
			target:  4,
			want:    [2]int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.numbers, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
