package kata

import (
	"reflect"
	"testing"
)

func TestSumDigPow(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name  string
		start uint64
		end   uint64
		want  []uint64
	}{
		{
			name:  "Test from 1 to 9",
			start: 1,
			end:   9,
			want:  []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:  "Test from 1 to 100",
			start: 1,
			end:   100,
			want:  []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 89},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumDigPow(tt.start, tt.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumDigPow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_doSum(t *testing.T) {
	tests := []struct {
		name    string
		val     uint64
		wantSum uint64
	}{
		{
			name:    "test 89",
			val:     89,
			wantSum: 89,
		},
		{
			name:    "test 3",
			val:     3,
			wantSum: 3,
		},
		{
			name:    "test 10",
			val:     10,
			wantSum: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := doSum(tt.val); gotSum != tt.wantSum {
				t.Errorf("doSum() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
