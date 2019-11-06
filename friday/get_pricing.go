package friday

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
)

// GetAdPricing - This function will return an array of Pricebuckets with quantity
func GetAdPricing(region string, category string, query string) (PriceBuckets, error) {
	var buckets PriceBuckets
	var chototAdResp ChototAds

	url, hasValue := os.LookupEnv("CHOTOT_SEARCH")

	if hasValue == false {
		log.Println("Error reading ENV for CHOTOT_SEARCH")
		return buckets, errors.New("Cannot read ENV")
	}

	url = url + "&region_v2=" + region + "&cg=" + category
	if query != "" {
		url = url + "&q=" + query
	}

	log.Println("URL: " + url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return buckets, err
	}

	host, _ := os.LookupEnv("FRIDAY_HOST")

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", host)
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("accept-encoding", "gzip, deflate")
	req.Header.Add("Connection", "keep-alive")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return buckets, err
	}
	defer res.Body.Close()

	var body []byte
	encoding := res.Header.Get("Content-Encoding")

	// Have data
	if encoding == "gzip" {
		reader, err := gzip.NewReader(res.Body)
		if err != nil {
			return buckets, err
		}
		defer reader.Close()

		body, err = ioutil.ReadAll(reader)
		if err != nil {
			return buckets, err
		}

		errJSON := json.Unmarshal(body, &chototAdResp)

		if errJSON != nil {
			return buckets, errJSON
		}

		sort.Slice(chototAdResp.Ads, func(i, j int) bool {
			return chototAdResp.Ads[i].Price < chototAdResp.Ads[j].Price
		})

		for _, v := range chototAdResp.Ads {
			buckets.PriceList = append(buckets.PriceList, v.Price)
		}

		l := len(chototAdResp.Ads)

		min := chototAdResp.Ads[0].Price
		max := chototAdResp.Ads[l-1].Price
		delta := (max - min) / 10

		from := min
		to := min + delta

		for to <= max {
			var pb PriceBucket
			pb.From = from
			pb.To = to
			pb.Quantity = 0
			for _, v := range chototAdResp.Ads {
				if v.Price >= from && v.Price <= to {
					pb.Quantity++
				}
			}
			buckets.Prices = append(buckets.Prices, pb)
			from = to + 1
			to += delta
		}
	}

	return buckets, nil
}
