package downloader

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"

	//"photobank-item-photo/app/resize/hendler"
	//_ "photobank-item-photo/app/resize/hendler"
)

type GetJsonPhotoConf struct {
	data []byte
	v    interface{}
}

type DownloadPhotoConf struct {
	Url string
}

type HttpClientG interface {
	Do(*http.Request) (*http.Response, error)
}
type Client2 struct {
	HttpClient HttpClientG
}

var ErrStatusCode = errors.New("error statuscode")

func (a *GetJsonPhotoConf) GetJsonPhoto() error {
	return json.Unmarshal(a.data, &a.v)
}
func (c Client) DownloadPhoto(a DownloadPhotoConf) ([]byte, error) {
	resp, err := HttpGet(a.Url)
	if resp.StatusCode != 200 {
		return nil, ErrStatusCode
	}
	req, err := http.NewRequest("GET", a.Url, nil)
	if err != nil {
		return nil, err
	}
	response, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	buf, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = response.Body.Close()
	if err != nil {
		return nil, err
	}
	return buf, err
}
