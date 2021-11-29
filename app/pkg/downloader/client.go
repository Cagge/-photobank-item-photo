package downloader

import (
	"io/ioutil"
	"net/http"
	"photobank-item-photo/app/pkg/downloader/models"
)

type HttpClientI interface {
	Do(*http.Request) (*http.Response, error)
}
type Client struct {
	HttpClient HttpClientI
}

func (c Client) GetJsonPhoto(a models.ConfigDataPhoto) (f models.Data, err error) {
	URLlisting := a.Photodate.BaseUrlListing
	User := a.Photodate.User
	Pass := a.Photodate.Pass
	bodyText, err := c.ClientAuth(URLlisting, User, Pass)
	if err != nil {
		return
	}
	m := GetJsonPhotoConf{bodyText, &f}
	err = m.GetJsonPhoto()
	return
}
func (c Client) ClientAuth(URLlisting string, User string, Pass string) ([]byte, error) {
	req, err := http.NewRequest("GET", URLlisting, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(User, Pass)
	return c.GetByteWithURL(req)
}
func (c Client) GetByteWithURL(req *http.Request) ([]byte, error) {
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	return bodyText, err
}
