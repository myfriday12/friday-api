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
		AdID           int    `json:"ad_id"`
		ListID         int    `json:"list_id"`
		ListTime       int64  `json:"list_time"`
		Date           string `json:"date"`
		AccountID      int    `json:"account_id"`
		AccountOid     string `json:"account_oid"`
		AccountName    string `json:"account_name"`
		Subject        string `json:"subject"`
		Body           string `json:"body"`
		Category       int    `json:"category"`
		Area           int    `json:"area"`
		AreaName       string `json:"area_name"`
		Region         int    `json:"region"`
		RegionName     string `json:"region_name"`
		CompanyAd      bool   `json:"company_ad"`
		Type           string `json:"type"`
		Price          int    `json:"price"`
		PriceString    string `json:"price_string"`
		Image          string `json:"image"`
		NumberOfImages int    `json:"number_of_images"`
		Avatar         string `json:"avatar"`
		RegionV2       int    `json:"region_v2"`
		AreaV2         int    `json:"area_v2"`
		Ward           int    `json:"ward"`
		WardName       string `json:"ward_name"`
	} `json:"ads"`
}

//SimilarAds - Holds similar ads data
type SimilarAds struct {
	OneSt []struct {
		AdID           int    `json:"ad_id"`
		ListID         int    `json:"list_id"`
		ListTime       int64  `json:"list_time"`
		Date           string `json:"date"`
		AccountID      int    `json:"account_id"`
		AccountOid     string `json:"account_oid"`
		AccountName    string `json:"account_name"`
		Subject        string `json:"subject"`
		Body           string `json:"body"`
		Category       int    `json:"category"`
		Area           int    `json:"area"`
		AreaName       string `json:"area_name"`
		Region         int    `json:"region"`
		RegionName     string `json:"region_name"`
		CompanyAd      bool   `json:"company_ad"`
		Type           string `json:"type"`
		Price          int    `json:"price"`
		PriceString    string `json:"price_string"`
		Image          string `json:"image"`
		NumberOfImages int    `json:"number_of_images"`
		Avatar         string `json:"avatar"`
		RegionV2       int    `json:"region_v2"`
		AreaV2         int    `json:"area_v2"`
		Ward           int    `json:"ward"`
		WardName       string `json:"ward_name"`
	} `json:"1st"`
	TwoNd []struct {
		AdID           int    `json:"ad_id"`
		ListID         int    `json:"list_id"`
		ListTime       int64  `json:"list_time"`
		Date           string `json:"date"`
		AccountID      int    `json:"account_id"`
		AccountOid     string `json:"account_oid"`
		AccountName    string `json:"account_name"`
		Subject        string `json:"subject"`
		Body           string `json:"body"`
		Category       int    `json:"category"`
		Area           int    `json:"area"`
		AreaName       string `json:"area_name"`
		Region         int    `json:"region"`
		RegionName     string `json:"region_name"`
		CompanyAd      bool   `json:"company_ad"`
		Type           string `json:"type"`
		Price          int    `json:"price"`
		PriceString    string `json:"price_string"`
		Image          string `json:"image"`
		NumberOfImages int    `json:"number_of_images"`
		Avatar         string `json:"avatar"`
		RegionV2       int    `json:"region_v2"`
		AreaV2         int    `json:"area_v2"`
		Ward           int    `json:"ward"`
		WardName       string `json:"ward_name"`
	} `json:"2nd"`
	ThreeRd []struct {
		AdID           int    `json:"ad_id"`
		ListID         int    `json:"list_id"`
		ListTime       int64  `json:"list_time"`
		Date           string `json:"date"`
		AccountID      int    `json:"account_id"`
		AccountOid     string `json:"account_oid"`
		AccountName    string `json:"account_name"`
		Subject        string `json:"subject"`
		Body           string `json:"body"`
		Category       int    `json:"category"`
		Area           int    `json:"area"`
		AreaName       string `json:"area_name"`
		Region         int    `json:"region"`
		RegionName     string `json:"region_name"`
		CompanyAd      bool   `json:"company_ad"`
		Type           string `json:"type"`
		Price          int    `json:"price"`
		PriceString    string `json:"price_string"`
		Image          string `json:"image"`
		NumberOfImages int    `json:"number_of_images"`
		Avatar         string `json:"avatar"`
		RegionV2       int    `json:"region_v2"`
		AreaV2         int    `json:"area_v2"`
		Ward           int    `json:"ward"`
		WardName       string `json:"ward_name"`
	} `json:"3rd"`
	FourTh []struct {
		AdID           int    `json:"ad_id"`
		ListID         int    `json:"list_id"`
		ListTime       int64  `json:"list_time"`
		Date           string `json:"date"`
		AccountID      int    `json:"account_id"`
		AccountOid     string `json:"account_oid"`
		AccountName    string `json:"account_name"`
		Subject        string `json:"subject"`
		Body           string `json:"body"`
		Category       int    `json:"category"`
		Area           int    `json:"area"`
		AreaName       string `json:"area_name"`
		Region         int    `json:"region"`
		RegionName     string `json:"region_name"`
		CompanyAd      bool   `json:"company_ad"`
		Type           string `json:"type"`
		Price          int    `json:"price"`
		PriceString    string `json:"price_string"`
		Image          string `json:"image"`
		NumberOfImages int    `json:"number_of_images"`
		Avatar         string `json:"avatar"`
		RegionV2       int    `json:"region_v2"`
		AreaV2         int    `json:"area_v2"`
		Ward           int    `json:"ward"`
		WardName       string `json:"ward_name"`
	} `json:"4th"`
}
