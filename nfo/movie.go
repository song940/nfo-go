package nfo

import (
	"encoding/xml"
	"io"
)

type Movie struct {
	XMLName xml.Name `xml:"movie"`

	Base
}

type Rating struct {
	Value   float64 `xml:"value,omitempty"`
	Votes   int64   `xml:"votes,omitempty"`
	Name    string  `xml:"name,attr,omitempty"`
	Max     int64   `xml:"max,attr,omitempty"`
	Default bool    `xml:"default,attr,omitempty"`
}

type Fanart struct {
	Thumb []Thumb `xml:"thumb,omitempty"`
}

type Thumb struct {
	Spoof   string `xml:"spoof,omitempty,attr"`
	Aspect  string `xml:"aspect,omitempty,attr"`
	Cache   string `xml:"cache,omitempty,attr"`
	Preview string `xml:"preview,omitempty,attr"`
	Colors  string `xml:"colors,omitempty,attr"`
	Link    string `xml:",chardata"`
}

type UniqueId struct {
	Type    string `xml:"type,omitempty,attr"`
	Default bool   `xml:"default,omitempty,attr"`
	Id      string `xml:",chardata"`
}

type StreamVideo struct {
	Codec             string `xml:"codec,omitempty"`
	Aspect            string `xml:"aspect,omitempty"`
	Width             int64  `xml:"width,omitempty"`
	Height            int64  `xml:"height,omitempty"`
	DurationInSeconds int64  `xml:"durationinseconds,omitempty"`
	StereoMode        string `xml:"stereomode,omitempty"`
}

type StreamAudio struct {
	Codec    string `xml:"codec,omitempty"`
	Language string `xml:"language,omitempty"`
	Channels int64  `xml:"channels,omitempty"`
}

type StreamSubtitle struct {
	Language string `xml:"language,omitempty"`
}

type Actor struct {
	Name  string `xml:"name,omitempty"`
	Type  string `xml:"type,omitempty"`
	Role  string `xml:"role,omitempty"`
	Order int64  `xml:"order,omitempty"`
	Thumb string `xml:"thumb,omitempty"`
}

func ReadMovieNfo(r io.Reader) (movie *Movie, err error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return
	}
	movie = &Movie{}
	err = xml.Unmarshal(b, &movie)
	if err != nil {
		return
	}
	return
}
