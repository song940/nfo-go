package nfo

import (
	"encoding/xml"
	"io"
)

type TVShow struct {
	Base

	XMLName        xml.Name `xml:"tvshow"`
	ShowTitle      string   `xml:"showtitle,omitempty"`
	Season         int64    `xml:"season,omitempty"`
	Episode        int64    `xml:"episode,omitempty"`
	DisplaySeason  int64    `xml:"displayseason,omitempty"`
	DisplayEpisode int64    `xml:"displayepisode,omitempty"`
	NamedSeason    struct {
		Number int64  `xml:"number,omitempty,attr"`
		Name   string `xml:",chardata"`
	} `xml:"namedseason,omitempty"`
}

func ReadTVShowNfo(r io.Reader) (*TVShow, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	m := TVShow{}
	err = xml.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
