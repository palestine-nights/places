package main

import "os"

const GoogleApi = "https://maps.googleapis.com/maps/api/place/details/json"

// Get environment variable with ability to specify default value.
func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		value = fallback
	}

	return value
}

func getGoolePlacesLink(key, placeID string) string {
	return GoogleApi + "?key=" + key + "&placeid=" + placeID
}
