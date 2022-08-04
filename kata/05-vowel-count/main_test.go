package kata

import "testing"

func TestGetCount(t *testing.T) {
	tests := []struct {
		name      string
		str       string
		wantCount int
	}{
		{
			name:      "Given test",
			str:       "abracadabra",
			wantCount: 5,
		},
		{
			name:      "Test different vowels",
			str:       "bubba",
			wantCount: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := GetCount(tt.str); gotCount != tt.wantCount {
				t.Errorf("input = %v, GetCount() = %v, want %v", tt.str, gotCount, tt.wantCount)
			}
		})
	}
}
