package resize

import (
	"net/http"
	. "photobank-item-photo/app/pkg/downloader/models"
	. "photobank-item-photo/app/pkg/resize/models"
	"photobank-item-photo/app/resize/hendler"
	"testing"

	"github.com/stretchr/testify/assert"
)

var UUint = []uint{25}
var f = REsize{
	OutSize:       UUint,
	PathSavePhoto: "/tmp/",
}
var ff = ConfigResize{f}
var v = ConfigDataPhoto{
	PHotodate{"***", "***", "http://photobank.sima-land.ru/photos?type=regular,3d&status_site=0,3&limit=100", "http://photobank.sima-land.ru:8080/static/items/"}}

func TestControlResizeValidFail(t *testing.T) {
	var UUintUnvalid = []uint{0}
	var f = REsize{
		OutSize:       UUintUnvalid,
		PathSavePhoto: "/tmp/",
	}
	var fff = ConfigResize{f}
	_, _, err := ControlResize(v, fff, http.DefaultClient)
	assert.NotNil(t, err)
}
func TestControlResizeJsonFail(t *testing.T) {
	var v = ConfigDataPhoto{
		PHotodate{"***", "***", "ht0", "http://photobank.sima-land.ru:8080/static/items/"}}
	_, _, err := ControlResize(v, ff, http.DefaultClient)
	assert.NotNil(t, err)
}

func TestProcessImageByUrlSuccess(t *testing.T) {
	p := hendler.PhotoDownloadCfg{"http://photobank.sima-land.ru:8080/static/items/", 5076211, 1}
	_, err := hendler.ProcessImageByUrl(p, http.DefaultClient)
	assert.Equal(t, err, nil)
}
