package friday

import (
	"testing"
)

func TestGetAdPricingSingle(t *testing.T) {
	_, err := GetAdPricing("13000", "5000", "Asus ROG Phone 8 GB Đen bóng - Jet black")

	if err != nil {
		t.Errorf("TestGetAdPricingSingle() gotError = %v, want no error", err.Error())
	}
}

func TestGetAdPricingMulti(t *testing.T) {
	_, err := GetAdPricing("13000", "5010", "iphone 6")

	if err != nil {
		t.Errorf("TestGetAdPricingMulti() gotError = %v, want no error", err.Error())
	}
}
