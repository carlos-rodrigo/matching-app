package storage

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"

	matching "github.com/carlos-rodrigo/matching-app/pkg/matching/model"
	"googlemaps.github.io/maps"
)

type CsvParticipantRepository struct {
	Participants []matching.Participant
}

func NewCsvParticipantsRepository(csvPath string) CsvParticipantRepository {
	participants := readCsvAndLoadParticipants(csvPath)
	return CsvParticipantRepository{
		Participants: participants,
	}
}

func readCsvAndLoadParticipants(csvPath string) []matching.Participant {
	csvFile, err := os.Open(csvPath)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvFile)
	participants := []matching.Participant{}

	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("error reading all lines: %v", err)
	}

	for i, line := range lines {
		if i == 0 {
			// skip header line
			continue
		}

		latitude, errLat := strconv.ParseFloat(line[5], 64)
		if errLat != nil {
			log.Fatalf("error converting latitude %q", errLat)
		}
		longitude, errLon := strconv.ParseFloat(line[6], 64)
		if errLon != nil {
			log.Fatalf("error converting latitude %q", errLon)
		}

		formattedAddress, errFormattedAddress := getFormattedAddressForLocation(latitude, longitude)
		if errFormattedAddress != nil {
			log.Fatalf("FormattedAddress can't be obtained %q", errFormattedAddress)
		}

		participants = append(participants, matching.Participant{
			Name:             line[0],
			Gender:           line[1],
			JobTitle:         line[2],
			Industry:         strings.Split(line[3], ","),
			FormattedAddress: formattedAddress,
			Location: matching.Location{
				Latitude:  latitude,
				Longitude: longitude,
			},
		})
	}
	return participants
}

func getFormattedAddressForLocation(latitude, longitude float64) (string, error) {
	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyA5QdtBhRRhPbJAx-oXUffbpQy3ARuSRt8")) //TODO: Change this hardcoded key. Make it an environment variable
	if err != nil {
		log.Fatalf("fatal error: %s", err)
		return "", err
	}

	r := &maps.GeocodingRequest{
		LatLng: &maps.LatLng{Lat: latitude, Lng: longitude},
	}

	resp, errReverseGeocode := c.ReverseGeocode(context.Background(), r)

	if errReverseGeocode != nil {
		log.Fatalf("error retriving formattedAddress %q", errReverseGeocode)
		return "", errReverseGeocode
	}

	return resp[0].FormattedAddress, nil

}
