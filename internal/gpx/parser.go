package gpx

import (
	"encoding/xml"
	"io"

	"github.com/pkg/errors"
)

func Parse(data io.Reader) (GPX, error) {
	gpx := GPX{}

	err := xml.NewDecoder(data).Decode(&gpx)
	if err != nil {
		return gpx, errors.Wrap(err, "parse gpx data")
	}

	return gpx, nil
}
