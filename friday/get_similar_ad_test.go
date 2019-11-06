package friday

import (
	"testing"
)

func TestGetSimilarAds(t *testing.T) {
	_, err := GetSimilarAds("5000", "250000", "Asus ROG Phone 8 GB Đen bóng - Jet black")
	if err != nil {
		t.Errorf("TestGetSimilarAds() gotError = %v, want no error", err.Error())
	}
}
