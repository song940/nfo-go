package nfo

import (
	"encoding/xml"
	"io"
)

type IMovie interface {
	isNFO() bool
}

type Movie struct {
	XMLName xml.Name `xml:"movie"`

	Title         string  `xml:"title,omitempty"`
	OriginalTitle string  `xml:"originaltitle,omitempty"`
	SortTitle     string  `xml:"sorttitle,omitempty"`
	Rating        float64 `xml:"rating,omitempty"`
	Ratings       struct {
		Rating []Rating `xml:"rating,omitempty"`
	} `xml:"ratings"`
	UserRating float64    `xml:"userrating,omitempty"`
	Top250     int64      `xml:"top250,omitempty"`
	Outline    string     `xml:"outline,omitempty"`
	Plot       string     `xml:"plot,omitempty"`
	Tagline    string     `xml:"tagline,omitempty"`
	Runtime    int64      `xml:"runtime,omitempty"`
	Thumb      []Thumb    `xml:"thumb,omitempty"`
	Fanart     *Fanart    `xml:"fanart,omitempty"`
	MPAA       string     `xml:"mpaa,omitempty"`
	Playcount  int64      `xml:"playcount,omitempty"`
	Lastplayed string     `xml:"lastplayed,omitempty"`
	Id         string     `xml:"id,omitempty"`
	Uniqueid   []UniqueId `xml:"uniqueid,omitempty"`
	Genre      []string   `xml:"genre,omitempty"`
	Tag        []string   `xml:"tag,omitempty"`
	Country    string     `xml:"country,omitempty"`
	Credits    string     `xml:"credits,omitempty"`
	Director   string     `xml:"director,omitempty"`
	Premiered  string     `xml:"premiered,omitempty"`
	Year       string     `xml:"year,omitempty"`
	Status     string     `xml:"status,omitempty"`
	Code       string     `xml:"code,omitempty"`
	Aired      string     `xml:"aired,omitempty"`
	Studio     string     `xml:"studio,omitempty"`
	Trailer    []string   `xml:"trailer,omitempty"`
	Actor      []Actor    `xml:"actor,omitempty"`
	DateAdded  string     `xml:"dateadded,omitempty"`
	IMDbId     string     `xml:"imdbid,omitempty"`
	TMDBId     string     `xml:"tmdbid,omitempty"`
	TVDBId     string     `xml:"tvdbid,omitempty"`

	Set struct {
		Name     string `xml:"name,omitempty"`
		Overview string `xml:"overview,omitempty"`
	} `xml:"set,omitempty"`
	Resume struct {
		Position float64 `xml:"position,omitempty"`
		Total    float64 `xml:"total,omitempty"`
	} `xml:"resume,omitempty"`
	FileInfo struct {
		StreamDetails struct {
			Video    []StreamVideo    `xml:"video,omitempty"`
			Audio    []StreamAudio    `xml:"audio,omitempty"`
			Subtitle []StreamSubtitle `xml:"subtitle,omitempty"`
		} `xml:"streamdetails,omitempty"`
	} `xml:"fileinfo,omitempty"`
}

type Rating struct {
	Value   float32 `xml:"value,omitempty"`
	Votes   int     `xml:"votes,omitempty"`
	Name    string  `xml:"name,attr,omitempty"`
	Max     float32 `xml:"max,attr,omitempty"`
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
	ID    int    `xml:"id,omitempty"`
	Name  string `xml:"name,omitempty"`
	Type  string `xml:"type,omitempty"`
	Role  string `xml:"role,omitempty"`
	Order int    `xml:"order,omitempty"`
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

func (m *Movie) isNFO() bool {
	return true
}
