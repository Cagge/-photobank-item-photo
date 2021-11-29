package downloader

import (
	"io"
	"net/http"
	"photobank-item-photo/app/pkg/downloader/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpGetFail(t *testing.T) {
	_, err := HttpGet("nic.ru/")
	assert.NotNil(t, err)
}
func TestHttpGetSuccess(t *testing.T) {
	_, err := HttpGet("https://www.nic.ru/")
	assert.Equal(t, err, nil)
}
func TestReadURLSuccess(t *testing.T) {
	dc := Client{
		HttpClient: http.DefaultClient,
	}
	m := DownloadPhotoConf{"https://www.nic.ru/"}
	_, err := dc.DownloadPhoto(m)
	assert.Equal(t, err, nil)
}
func TestReadURLFail2(t *testing.T) {
	dc := Client{
		HttpClient: http.DefaultClient,
	}
	m := DownloadPhotoConf{"https://www.nic.ru/2222222222"}
	_, err := dc.DownloadPhoto(m)
	assert.NotNil(t, err)
}
func TestClientDoSuccess(t *testing.T) {
	client := HttpClient(nil)
	req, err := http.NewRequest("GET", "https://www.nic.ru/", nil)
	assert.Equal(t, err, nil)
	_, err = ClientDo(req, client)
	assert.Equal(t, err, nil)
}
func TestClientDoFail(t *testing.T) {
	client := HttpClient(nil)
	req, err := http.NewRequest("GET", "htt/www.nic.ru/", nil)
	_, err = ClientDo(req, client)
	assert.NotNil(t, err)
}

var v = models.ConfigDataPhoto{
	models.PHotodate{"asd", "asd", "123", "1243"}}
var f = models.Data{
	Count: 2,
	Photos: []models.Phot{
		{
			555, 2555555, "regular",
		},
		{
			45555, 5222222, "3d",
		},
	},
}

func TestGetJsonPhotoFail(t *testing.T) {
	dc := Client{
		HttpClient: http.DefaultClient,
	}
	_, err := dc.GetJsonPhoto(v)
	assert.NotNil(t, err)
}
func TestGetJsonPhotoFail2(t *testing.T) {
	v = models.ConfigDataPhoto{
		models.PHotodate{"asd",
			"asd",
			"https://www.sima-land.ru/",
			"https://www.sima-land.ru/"}}
	dc := Client{
		HttpClient: http.DefaultClient,
	}
	_, err := dc.GetJsonPhoto(v)
	assert.NotNil(t, err)
}

func TestReedAllSuccess(t *testing.T) {
	response, err := HttpGet("https://github.com/")
	assert.Equal(t, err, nil)
	_, err = io.ReadAll(response.Body)
	assert.Equal(t, err, nil)
	err = response.Body.Close()
	assert.Equal(t, err, nil)
}
func TestReedAllFail(t *testing.T) {
	_, err := HttpGet("htps://github.com/")
	assert.NotNil(t, err)
}
