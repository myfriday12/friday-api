package friday

// PriceBucket - Will hold the individual bucket price range and how
// many ads returned from the search fall into the price range
type PriceBucket struct {
	From     int `json:"from"`
	To       int `json:"to"`
	Quantity int `json:"quantity"`
}

// PriceBuckets - Will hold multiple priceBucket entries
type PriceBuckets struct {
	Prices    []PriceBucket `json:"chart_data"`
	PriceList []int         `json:"price_list"`
}

// ChototAds - This will hold the response returned by Chotots Search
type ChototAds struct {
	Total int `json:"total"`
	Ads   []struct {
		Price int `json:"price,omitempty"`
	} `json:"ads"`
}
