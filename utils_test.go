package tronwallet

import (
	"math"
	"math/big"
	"testing"
)

func TestValidateAmount(t *testing.T) {
	tests := []struct {
		name    string
		amount  *big.Int
		wantErr bool
	}{
		{"nil", nil, true},
		{"negative", big.NewInt(-1), true},
		{"zero", big.NewInt(0), false},
		{"positive", big.NewInt(1_000_000), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateAmount(tt.amount)
			if (err != nil) != tt.wantErr {
				t.Fatalf("validateAmount(%v) error = %v, wantErr %v", tt.amount, err, tt.wantErr)
			}
		})
	}
}

func TestAmountToInt64(t *testing.T) {
	overflow := new(big.Int).Add(big.NewInt(math.MaxInt64), big.NewInt(1))

	tests := []struct {
		name    string
		amount  *big.Int
		want    int64
		wantErr bool
	}{
		{"nil", nil, 0, true},
		{"negative", big.NewInt(-1), 0, true},
		{"zero", big.NewInt(0), 0, false},
		{"one trx in sun", big.NewInt(1_000_000), 1_000_000, false},
		{"max int64", big.NewInt(math.MaxInt64), math.MaxInt64, false},
		{"overflow", overflow, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := amountToInt64(tt.amount)
			if (err != nil) != tt.wantErr {
				t.Fatalf("amountToInt64(%v) error = %v, wantErr %v", tt.amount, err, tt.wantErr)
			}
			if err == nil && got != tt.want {
				t.Fatalf("amountToInt64(%v) = %d, want %d", tt.amount, got, tt.want)
			}
		})
	}
}
