package friday

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func callSearch(search string, limit string) (ChototAds, error) {
	var chototAds ChototAds

	capi, _ := os.LookupEnv("CHOTOT_SEARCH")

	//Replace limit
	capi = strings.Replace(capi, "limit=50", "limit="+limit, -1)

	capi = capi + search

	fmt.Println("capi: " + capi)

	req, err := http.NewRequest("GET", capi, nil)
	if err != nil {
		return chototAds, err
	}

	host, _ := os.LookupEnv("FRIDAY_HOST")

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", host)
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("accept-encoding", "gzip, deflate")
	req.Header.Add("Connection", "keep-alive")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return chototAds, err
	}
	defer res.Body.Close()

	var body []byte
	encoding := res.Header.Get("Content-Encoding")

	// Have data
	if encoding == "gzip" {
		reader, err := gzip.NewReader(res.Body)
		if err != nil {
			return chototAds, err
		}
		defer reader.Close()

		body, err = ioutil.ReadAll(reader)
		if err != nil {
			return chototAds, err
		}

		errJSON := json.Unmarshal(body, &chototAds)

		if errJSON != nil {
			return chototAds, errJSON
		}

	}
	return chototAds, nil
}

//GetSimilarAds - Returns similar ads
func GetSimilarAds(category string, price string, query string) (SimilarAds, error) {
	var similarAds SimilarAds

	intPrice, err := strconv.Atoi(price)
	if err != nil {
		return similarAds, err
	}

	floatPrice := float64(intPrice)
	less1PcPrice := strconv.Itoa(int(floatPrice * 0.99))
	less25PcPrice := strconv.Itoa(int(floatPrice * 0.75))
	more5PcPrice := strconv.Itoa(int(floatPrice * 1.05))
	more25PcPrice := strconv.Itoa(int(floatPrice * 1.25))

	// Generate 1st == cg=5010&price=2699999-2700000 --> Less 1%
	sOne := "&cg=" + category + "&price=" + less1PcPrice + "-" + strconv.Itoa(intPrice)
	fmt.Println(sOne)
	sOneAds, sOneErr := callSearch(sOne, "1")
	if sOneErr != nil {
		return similarAds, sOneErr
	}

	if sOneAds.Ads != nil {
		similarAds.OneSt = sOneAds.Ads
	}

	// Generate 2nd == cg=5010&q=iphone6
	sTwo := "&cg=" + category + "&q=" + url.QueryEscape(query)
	fmt.Println(sTwo)
	sTwoAds, sTwoErr := callSearch(sTwo, "1")
	if sTwoErr != nil {
		return similarAds, sTwoErr
	}

	if sTwoAds.Ads != nil {
		similarAds.TwoNd = sTwoAds.Ads
	}

	// Generate 3rd == cg=5010&price=2800000-4000000 --> 1% more - 25% more
	sThree := "&cg=" + category + "&price=" + more5PcPrice + "-" + more25PcPrice
	fmt.Println(sThree)
	sThreeAds, sThreeErr := callSearch(sThree, "1")
	if sThreeErr != nil {
		return similarAds, sThreeErr
	}

	if sThreeAds.Ads != nil {
		similarAds.ThreeRd = sThreeAds.Ads
	}

	// Generate 4th == cg=5010&price=0-2000000 --> 0  to Less 25%
	sFour := "&cg=" + category + "&price=0" + "-" + less25PcPrice
	fmt.Println(sFour)
	sFourAds, sFourErr := callSearch(sFour, "2")
	if sFourErr != nil {
		return similarAds, sFourErr
	}

	if sFourAds.Ads != nil {
		similarAds.FourTh = sFourAds.Ads
	}

	return similarAds, nil
}
