package data

import (
	"math"
	"testing"
)

func TestGetInsertID(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		want    int
		wantErr bool
	}{
		{
			name:    "int64 positive",
			input:   int64(42),
			want:    42,
			wantErr: false,
		},
		{
			name:    "int64 zero",
			input:   int64(0),
			want:    0,
			wantErr: false,
		},
		{
			name:    "int",
			input:   123,
			want:    123,
			wantErr: false,
		},
		{
			name:    "uint64 valid",
			input:   uint64(999),
			want:    999,
			wantErr: false,
		},
		{
			name:    "uint valid",
			input:   uint(555),
			want:    555,
			wantErr: false,
		},
		{
			name:    "int64 negative",
			input:   int64(-1),
			want:    0,
			wantErr: true,
		},
		{
			name:    "uint64 overflow",
			input:   uint64(math.MaxInt64) + 1,
			want:    0,
			wantErr: true,
		},
		{
			name:    "unsupported type string",
			input:   "not a number",
			want:    0,
			wantErr: true,
		},
		{
			name:    "unsupported type float",
			input:   3.14,
			want:    0,
			wantErr: true,
		},
		{
			name:    "nil",
			input:   nil,
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getInsertID(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("getInsertID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getInsertID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInsertID_MaxInt(t *testing.T) {
	// Test that max int value works
	maxInt := int(^uint(0) >> 1)

	got, err := getInsertID(int64(maxInt))
	if err != nil {
		t.Errorf("getInsertID(maxInt) should not error, got: %v", err)
	}
	if got != maxInt {
		t.Errorf("getInsertID(maxInt) = %v, want %v", got, maxInt)
	}
}
