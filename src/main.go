package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

var (
	app *App
)

type GooglePlaceReview struct {
	AuthorName              string `json:"author_name"`
	AuthorUrl               string `json:"author_url"`
	Language                string `json:"language"`
	ProfilePhotoUrl         string `json:"profile_photo_url"`
	Rating                  int8   `json:"rating"`
	RelativeTimeDescription string `json:"relative_time_description"`
	Text                    string `json:"text"`
	Time                    int64  `json:"time"`
}

type GooglePlace struct {
	// AddressComponents
	Address              string `json:"adr_address"`
	FormattedAddress     string `json:"formatted_address"`
	FormattedPhoneNumber string `json:"formatted_phone_number"`
	// Geometry
	Icon                     string `json:"icon"`
	Id                       string `json:"id"`
	InternationalPhoneNumber string `json:"international_phone_number"`
	Name                     string `json:"name"`
	// OpeningHours
	// Photos
	PlaceId   string              `json:"place_id"`
	Rating    float64             `json:"rating"`
	Reference string              `json:"reference"`
	Reviews   []GooglePlaceReview `json:"reviews"`
	// Scope
	// Types
	Url       string `json:"url"`
	UtcOffset int16  `json:"utc_offset"`
	Vicinity  string `json:"vicinity"`
	Website   string `json:"Website"`
}

type GooglePlacesStandartResponce struct {
	Result GooglePlace `json:"result"`
	Status string      `json:"status"`
}

func handleError(err error, w *http.ResponseWriter) {
	fmt.Print(err)
	http.Error(*w, err.Error(), http.StatusInternalServerError)
	(*w).WriteHeader(http.StatusInternalServerError)
}

func GetRating(w http.ResponseWriter, r *http.Request) {
	placeID := r.FormValue("placeid")

	rating, err := app.Redis.Get(placeID).Result()

	if err == redis.Nil {
		key := r.Header.Get("Authorization")
		fmt.Printf("Authorization: %s\n", key)
		link := getGoolePlacesLink(key, placeID)
		response, err := http.Get(link)

		if err != nil {
			// Issue with connection to Goole Places API.
			handleError(err, &w)
			return
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			var googlePlace GooglePlacesStandartResponce
			json.Unmarshal(data, &googlePlace)

			rating = fmt.Sprintf("%f", googlePlace.Result.Rating)
			app.Redis.Set(placeID, rating, app.Experation).Err()
			fmt.Printf("%s was stored for 1h, value: %s\n", placeID, rating)
		}
	} else if err != nil {
		// Issue with connection to redis server.
		handleError(err, &w)
		return
	} else {
		fmt.Printf("%s exist in redis\n", rating)
	}

	ratingAsFloat, err := strconv.ParseFloat(rating, 32)

	if err != nil {
		// Issue with converting rating to float64.
		handleError(err, &w)
		return
	}

	fmt.Println("Float ", ratingAsFloat)

	response, err := json.Marshal(
		ServerResponce{
			Rating: ratingAsFloat,
		},
	)

	if err != nil {
		// Issue with marshal ServerResponce to json.
		handleError(err, &w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func initRouter() *mux.Router {
	router := mux.NewRouter()

	ratingPath := router.Path("/")
	ratingPath.HandlerFunc(GetRating).Methods("GET")
	ratingPath.Queries("placeid", "{placeid}")

	return router
}

func initRedis() *redis.Client {
	redisHost := getEnv("REDIS_HOST", "localhost")
	redisPort := getEnv("REDIS_PORT", "6379")
	redisPassword := getEnv("REDIS_PASSWORD", "")
	redisURL := redisHost + ":" + redisPort

	redis := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: redisPassword,
		DB:       0,
	})

	pong, err := redis.Ping().Result()
	fmt.Println(pong, err)

	return redis
}

func main() {
	app = GetApp()

	log.Printf("Listening on " + app.getAddress())
	log.Fatal(http.ListenAndServe(app.getAddress(), *app.GetHandler()))
}
