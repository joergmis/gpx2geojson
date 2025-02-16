package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/joergmis/gpx2geojson/internal/geojson"
	"github.com/joergmis/gpx2geojson/internal/gpx"
)

func main() {
	var (
		input  gpx.GPX
		output geojson.GeoJSON
	)

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("read gpx input: %v\n", err)
	}

	input, err = gpx.Parse(bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	output = convert(input)
	err = json.NewEncoder(os.Stdout).Encode(&output)
	if err != nil {
		log.Fatal(err)
	}
}

func convert(input gpx.GPX) geojson.GeoJSON {
	converted := geojson.GeoJSON{
		Type:     "FeatureCollection",
		Features: []geojson.Feature{},
	}

	for _, track := range input.Tracks {
		for _, segment := range track.Segments {
			feature := geojson.Feature{
				Type: "Feature",
				Geometry: geojson.Geometry{
					Type: "Polygon",
					Coordinates: [][][]float64{
						{},
					},
				},
			}

			for _, point := range segment.Point {
				feature.Geometry.Coordinates[0] = append(feature.Geometry.Coordinates[0], []float64{
					point.Longitude,
					point.Latitude,
				})
			}

			// to be a valid polygon, geojson requires that the first point and the last point are the same
			feature.Geometry.Coordinates[0] = append(feature.Geometry.Coordinates[0], []float64{
				segment.Point[0].Longitude,
				segment.Point[0].Latitude,
			})

			converted.Features = append(converted.Features, feature)
		}
	}

	return converted
}
