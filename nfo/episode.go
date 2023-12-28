package nfo

import (
	"encoding/xml"
	"io"
)

type Episode struct {
	XMLName xml.Name `xml:"episodedetails"`

	TVShow

	FileInfo struct {
		StreamDetails struct {
			Video    []StreamVideo    `xml:"video,omitempty"`
			Audio    []StreamAudio    `xml:"audio,omitempty"`
			Subtitle []StreamSubtitle `xml:"subtitle,omitempty"`
		} `xml:"streamdetails,omitempty"`
	} `xml:"fileinfo,omitempty"`
}

func ReadEpisodeNfo(r io.Reader) (*Episode, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	m := Episode{}
	err = xml.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
