package kata

import "testing"

func TestTwoToOne(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want string
	}{
		{
			name: "Test different strings",
			s1:   "abcd",
			s2:   "efgh",
			want: "abcdefgh",
		},
		{
			name: "Test backwards strings",
			s1:   "dcba",
			s2:   "hgfe",
			want: "abcdefgh",
		},
		{
			name: "Duplicate letters",
			s1:   "aaaaa",
			s2:   "a",
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoToOne(tt.s1, tt.s2); got != tt.want {
				t.Errorf("TwoToOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
