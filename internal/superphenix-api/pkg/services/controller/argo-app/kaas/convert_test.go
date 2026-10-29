package kaas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuantityToGiB(t *testing.T) {
	tests := []struct {
		name     string
		quantity string
		want     int
	}{
		{name: "gibibytes", quantity: "24Gi", want: 24},
		{name: "canonical tebibyte", quantity: "1Ti", want: 1024},
		{name: "non canonical gibibytes", quantity: "1536Gi", want: 1536},
		{name: "canonical tebibytes", quantity: "2Ti", want: 2048},
		{name: "largest accepted size", quantity: "999999Gi", want: 999999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, quantityToGiB(tt.quantity))
		})
	}
}

func TestQuantityToBinary(t *testing.T) {
	tests := []struct {
		name     string
		quantity string
		want     string
	}{
		{name: "gibibytes", quantity: "4Gi", want: "4Gi"},
		{name: "fractional gibibytes", quantity: "1.5Gi", want: "1536Mi"},
		{name: "mebibytes", quantity: "2500Mi", want: "2500Mi"},
		{name: "mebibytes as gibibytes", quantity: "1024Mi", want: "1Gi"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, quantityToBinary(tt.quantity))
		})
	}
}
