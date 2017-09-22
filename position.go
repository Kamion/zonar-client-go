package zonar

import (
	"encoding/xml"
	"net/url"
)

type CurrentLocations struct {
	XMLName xml.Name `xml:"currentlocations"`
	Assets  []Asset  `xml:"asset"`
}

type Asset struct {
	XMLName xml.Name `xml:"asset"`
	TagID   string   `xml:"tagid,attr"`
	Fleet   string   `xml:"fleet,attr"`
	ID      string   `xml:"id,attr"`
	Type    string   `xml:"type,attr"`
	Lon     float64  `xml:"long"`
	Lat     float64  `xml:"lat"`
	Heading float64  `xml:"heading"`
	Time    string   `xml:"time"`
	Speed   float64  `xml:"speed"`
	Power   string   `xml:"power"`
}

func (c *Client) GetCurrentPosition() (CurrentLocations, error) {
	form := url.Values{}

	form.Add("action", "showposition")
	form.Add("operation", "current")
	form.Add("format", "xml")
	form.Add("logvers", "2")

	result := CurrentLocations{}

	resp, err := c.request(form)
	if err != nil {
		return result, err
	}

	err = xml.Unmarshal(resp, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
