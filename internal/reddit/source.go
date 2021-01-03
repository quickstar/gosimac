package reddit

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/pterm/pterm"
)

// ErrRequestFailed indicates a general error in service request.
var ErrRequestFailed = errors.New("request failed")

// Reddit is source implmentation for unsplash image service.
type Reddit struct {
	N      int
	Query  string
	Size   string
	Path   string
	Prefix string
	Client *resty.Client
}

func New(count int, query string, path string, size string) *Reddit {
	return &Reddit{
		N:      count,
		Query:  query,
		Path:   path,
		Prefix: "reddit",
		Client: resty.New().
			SetBaseURL("https://www.reddit.com").
			SetHeader("User-Agent", "wally:v1.2.69"),
	}
}

// gather images urls from unplash based on given critarias.
func (r *Reddit) gather() (*Response, error) {
	var response Response

	resp, err := r.Client.R().
		SetResult(&response).
		SetQueryParam("q", r.Query).
		SetQueryParam("sort", "top").
		SetQueryParam("limit", strconv.Itoa(r.N)).
		SetQueryParam("restrict_sr", "on").
		SetQueryParam("t", "all").
		Get("/r/wallpapers/search.json")

	if err != nil {
		return nil, fmt.Errorf("network failure: %w", err)
	}

	if resp.IsError() {
		pterm.Error.Printf("reddit response code is %d: %s", resp.StatusCode(), resp.String())

		return nil, ErrRequestFailed
	}

	return &response, nil
}

// Fetch fetches given index from source.
func (r *Reddit) Fetch() error {
	response, err := r.gather()
	if err != nil {
		return fmt.Errorf("gatering information from reddit failed %w", err)
	}

	for _, image := range response.Data.Children {
		pterm.Info.Printf("Getting %s (%s)\n", image.Data.ID, image.Data.Title)

		resp, err := resty.New().R().SetDoNotParseResponse(true).Get(image.Data.URL)
		if err != nil {
			return fmt.Errorf("network failure: %w", err)
		}

		if resp.IsError() {
			pterm.Error.Printf("reddit response code is %d: %s", resp.StatusCode(), resp.String())

			return ErrRequestFailed
		}

		pterm.Success.Printf("%s was gotten\n", image.Data.Title)

		go r.Store(image.Data.Title, resp.RawBody())
	}

	return nil
}

func (r *Reddit) Store(name string, content io.ReadCloser) {
	path := path.Join(r.Path, fmt.Sprintf("%s-%s.jpg", r.Prefix, name))

	if _, err := os.Stat(path); err == nil {
		pterm.Warning.Printf("%s is already exists\n", path)

		return
	}

	file, err := os.Create(path)
	if err != nil {
		pterm.Error.Printf("os.Create: %v\n", err)

		return
	}

	bytes, err := io.Copy(file, content)
	if err != nil {
		pterm.Error.Printf("io.Copy (%d bytes): %v\n", bytes, err)
	}

	if err := file.Close(); err != nil {
		pterm.Error.Printf("(*os.File).Close: %v", err)
	}

	if err := content.Close(); err != nil {
		pterm.Error.Printf("(*io.ReadCloser).Close: %v", err)
	}
}
