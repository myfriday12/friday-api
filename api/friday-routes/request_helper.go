package fridayroutes

import (
	fridayf "friday-api/friday"
	"net/http"
)

//handleGerPricing - This function will return an array of Pricebuckets with quantity
func handleGerPricing(region string, category string, query string) (int, interface{}) {
	if region == "" {
		return http.StatusUnprocessableEntity, "(HGPR) Missing region"
	}
	if category == "" {
		return http.StatusUnprocessableEntity, "(HGPC) Missing category"
	}

	buckets, err := fridayf.GetAdPricing(region, category, query)

	if err != nil {
		return JsonapiErrorResp(http.StatusNotAcceptable, err.Error())
	}

	return http.StatusOK, buckets
}

//handleRateAdUp - This function will rate a specific ad up
func handleRateAdUp(listID string) (int, interface{}) {
	//TODO:
	if listID == "" {
		return http.StatusUnprocessableEntity, "(HRAU) Missing List ID"
	}
	return http.StatusNotImplemented, nil
}

//handleRateAdDown - This function will rate a specific ad down
func handleRateAdDown(listID string) (int, interface{}) {
	//TODO:
	if listID == "" {
		return http.StatusUnprocessableEntity, "(HRAU) Missing List ID"
	}
	return http.StatusNotImplemented, nil
}

//handlegetAdRating - This function will return the current rating of an ad
func handlegetAdRating(listID string) (int, interface{}) {
	//TODO:
	if listID == "" {
		return http.StatusUnprocessableEntity, "(HRAU) Missing List ID"
	}
	return http.StatusNotImplemented, nil
}
