package fridayroutes

import (
	"net/http"
	"testing"
)

func Test_handleGerPricing(t *testing.T) {
	//region=13000&cg=5000&q=Asus ROG Phone 8 GB Đen bóng - Jet black

	status, _ := handleGerPricing("13000", "5000", "Asus ROG Phone 8 GB Đen bóng - Jet black")

	if status != http.StatusOK {
		t.Errorf("handleGerPricing() gotStaus = %v, wantStatus %v", status, http.StatusOK)
	}
}
