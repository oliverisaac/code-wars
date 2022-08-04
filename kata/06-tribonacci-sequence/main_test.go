package kata

import (
	"reflect"
	"testing"
)

func TestTribonacci(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name      string
		signature [3]float64
		n         int
		want      []float64
	}{
		{
			name: "No results",
			signature: [3]float64{
				0.0,
				0.0,
				0.0,
			},
			n:    0,
			want: []float64{},
		},
		{
			name: "Basic input handling",
			signature: [3]float64{
				1.0,
				1.0,
				1.0,
			},
			n:    3,
			want: []float64{1, 1, 1},
		},
		{
			name: "Do the math",
			signature: [3]float64{
				1.0,
				1.0,
				1.0,
			},
			n:    4,
			want: []float64{1, 1, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tribonacci(tt.signature, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tribonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
