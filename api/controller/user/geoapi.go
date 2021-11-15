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

type geoAddress struct {
	HouseNumber string `json:"house_number"`
	Road        string `json:"road"`
	Suburb      string `json:"suburb"`
	Borough     string `json:"borough"`
	City        string `json:"city"`
	State       string `json:"state"`
	Postcode    string `json:"postcode"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
}

type geoLocations struct {
	Lat     string     `json:"lat"`
	Lon     string     `json:"lon"`
	OsmID   int64      `json:"osm_id"`
	OsmType string     `json:"osm_type"`
	Address geoAddress `json:"address"`
}

func getValidAddress(street string, housenumber string, city string) ([]geoLocations, error) {
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
	}
	return locations, nil
}

func getValidLookup(osmID int64, osmType string) ([]geoLocations, error) {
	// TODO can verify list of osm_id (up to 50)
	url := "https://nominatim.openstreetmap.org/lookup?format=json&addressdetails=1&osm_ids=" + url.QueryEscape(strings.ToUpper(osmType[0:1])+strconv.Itoa(int(osmID)))
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
	}
	return locations, nil
}
