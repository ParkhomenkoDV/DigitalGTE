package utils

import "testing"

func TestAll(t *testing.T) {
	tests := []struct {
		name string
		sl   []bool
		want bool
	}{
		{
			name: "all",
			sl:   []bool{true, true, true},
			want: true,
		},
		{
			name: "!all",
			sl:   []bool{true, false, true},
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := All(test.sl...)
			if got != test.want {
				t.Errorf("All() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestAny(t *testing.T) {
	tests := []struct {
		name string
		sl   []bool
		want bool
	}{
		{
			name: "any",
			sl:   []bool{false, true, false},
			want: true,
		},
		{
			name: "!any",
			sl:   []bool{false, false, false},
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Any(test.sl...)
			if got != test.want {
				t.Errorf("Any() = %v, want %v", got, test.want)
			}
		})
	}
}
