package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	GoogleApi = "https://maps.googleapis.com/maps/api/place/details/json"
)

type GooglePlaceReview struct {
	Author_name               string `json:"author_name"`
	Author_url                string `json:"author_url"`
	Language                  string `json:"language"`
	Profile_photo_url         string `json:"profile_photo_url"`
	Rating                    int8   `json:"rating"`
	Relative_time_description string `json:"relative_time_description"`
	Text                      string `json:"text"`
	Time                      int64  `json:"time"`
}

type GooglePlace struct {
	// address_components
	Address              string `json:"adr_address"`
	FormattedAddress     string `json:"formatted_address"`
	FormattedPhoneNumber string `json:"formatted_phone_number"`
	// geometry
	Icon                     string `json:"icon"`
	Id                       string `json:"id"`
	InternationalPhoneNumber string `json:"international_phone_number"`
	Name                     string `json:"name"`
	// opening_hours
	// photos
	PlaceId   string              `json:"place_id"`
	Rating    float64             `json:"rating"`
	Reference string              `json:"reference"`
	Reviews   []GooglePlaceReview `json:"reviews"`
	// scope
	// types
	Url       string `json:"url"`
	UtcOffset int16  `json:"utc_offset"`
	Vicinity  string `json:"vicinity"`
	Website   string `json:"Website"`
}

type GooglePlacesStandartResponce struct {
	Result GooglePlace `json:"result"`
	Status string      `json:"status"`
}

func GetRating(w http.ResponseWriter, r *http.Request) {
	placeId := r.FormValue("placeid")
	key := r.FormValue("key")

	fmt.Print(key)

	link := GoogleApi + "?key=" + key + "&placeid=" + placeId

	response, err := http.Get(link)

	if err != nil {
		fmt.Printf("%s", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		var googlePlace GooglePlacesStandartResponce
		json.Unmarshal(data, &googlePlace)

		ratingResult := map[string]interface{}{
			"rating": googlePlace.Result.Rating,
		}

		response, _ := json.Marshal(ratingResult)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

func main() {
	router := mux.NewRouter()

	ratingPath := router.Path("/")
	ratingPath.HandlerFunc(GetRating)
	ratingPath.Queries(
		"key", "{key}",
		"placeid", "{placeid}",
	)

	log.Fatal(http.ListenAndServe(":8000", router))
}