package kata

import "testing"

func TestDuplicateEncode(t *testing.T) {
	tests := []struct {
		name string
		word string
		want string
	}{
		{
			name: "test simple word",
			word: "din",
			want: "(((",
		},
		{
			name: "test repeat of single letter",
			word: "recede",
			want: "()()()",
		},
		{
			name: "Test multiple repeat letters",
			word: "Success",
			want: ")())())",
		},
		{
			name: "Test with parens",
			word: "(( @",
			want: "))((",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DuplicateEncode(tt.word); got != tt.want {
				t.Errorf("Input = %v, DuplicateEncode() = %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}
