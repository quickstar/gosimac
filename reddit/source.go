package reddit

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// ErrRequestFailed indicates a general error in service request.
var ErrRequestFailed = errors.New("request failed")

// Source is source implmentation for unsplash image service.
type Source struct {
	response   Response
	Count      int
	Query      string
	Resolution string
}

// Init initiates source and return number of available images.
func (s *Source) Init() (int, error) {
	resp, err := resty.New().
		SetHostURL("https://www.reddit.com").
		R().
		SetResult(&s.response).
		SetQueryParam("q", s.Resolution).
		SetQueryParam("sort", "top").
		SetQueryParam("limit", strconv.Itoa(s.Count)).
		SetQueryParam("restrict_sr", "0").
		SetQueryParam("t", "all").
		Get("/r/wallpaper/search.json")

	if err != nil {
		return 0, err
	}

	if resp.StatusCode() != http.StatusOK {
		return 0, ErrRequestFailed
	}

	return len(s.response.Data.Children), nil
}

// Name returns source name.
func (s *Source) Name() string {
	return "reddit"
}

// Fetch fetches given index from source.
func (s *Source) Fetch(index int) (string, io.ReadCloser, error) {
	image := s.response.Data.Children[index]

	logrus.Infof("Getting %s", image.Data.Title)

	resp, err := resty.New().R().SetDoNotParseResponse(true).Get(image.Data.URL)
	if err != nil {
		return "", nil, err
	}
	logrus.Infof("%s was gotten", image.Data.Title)

	return fmt.Sprintf("%s.jpg", image.Data.Title), resp.RawBody(), nil
}
