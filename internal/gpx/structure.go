package gpx

type GPX struct {
	Metadata Metadata `xml:"metadata"`
	Tracks   []Track  `xml:"trk"`
}

type Metadata struct{}

type Track struct {
	Segments []Segment `xml:"trkseg"`
}

type Segment struct {
	Point []Point `xml:"trkpt"`
}

type Point struct {
	Latitude  float64 `xml:"lat,attr"`
	Longitude float64 `xml:"lon,attr"`
	Elevation float64 `xml:"ele"`
}
