package user

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type GeoAddress struct {
	House_number string `json:"house_number"`
	Road         string `json:"road"`
	Suburb       string `json:"suburb"`
	Borough      string `json:"borough"`
	City         string `json:"city"`
	State        string `json:"state"`
	Postcode     string `json:"postcode"`
	Country      string `json:"country"`
	Country_code string `json:"country_code"`
}

type geoLocations struct {
	Lat      string     `json:"lat"`
	Lon      string     `json:"lon"`
	Osm_id   int64      `json:"osm_id"`
	Osm_type string     `json:"osm_type"`
	Address  GeoAddress `json:"address"`
}

func GetValidAddress(street string, housenumber string, city string) ([]geoLocations, error) {
	url := "https://nominatim.openstreetmap.org/search?addressdetails=1&format=json&street=" + url.QueryEscape(housenumber+" ") + url.QueryEscape(street) + "&city=" + url.QueryEscape(city)
	// url := "https://nominatim.openstreetmap.org/search?addressdetails=1&format=json&street=" + url.QueryEscape(street)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return []geoLocations{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		return []geoLocations{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []geoLocations{}, err
	}
	var locations []geoLocations
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return []geoLocations{}, err
	}
	if len(locations) == 0 {
		return []geoLocations{}, errors.New("no locations found")
	} else {
		return locations, nil
	}
}

func GetValidLookup(osm_id int64, osm_type string) ([]geoLocations, error) {
	// TODO can verify list of osm_id (up to 50)
	url := "https://nominatim.openstreetmap.org/lookup?format=json&addressdetails=1&osm_ids=" + url.QueryEscape(strings.ToUpper(osm_type[0:1])+strconv.Itoa(int(osm_id)))
	// url := "https://nominatim.openstreetmap.org/search?addressdetails=1&format=json&street=" + url.QueryEscape(street)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return []geoLocations{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		return []geoLocations{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []geoLocations{}, err
	}
	var locations []geoLocations
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return []geoLocations{}, err
	}
	if len(locations) == 0 {
		return []geoLocations{}, errors.New("no locations found")
	} else {
		return locations, nil
	}
}
