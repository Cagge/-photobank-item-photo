package hendler

import (
	"fmt"
	"github.com/pkg/errors"
	. "photobank-item-photo/app/pkg/downloader"

	"strconv"
)

type PhotoDownloadCfg struct {
	Urldownload string
	Sid         int
	Number      int
}

var ErrValueSid = errors.New("error value size")

func ProcessImageByUrl(p PhotoDownloadCfg,dhc httpClientH) ([]byte, error) {
	if p.Sid < 1000000 {
		return nil, ErrValueSid
	}
	url := fmt.Sprint(p.Urldownload, strconv.Itoa(p.Sid)[0:3], "/", p.Sid, "/", p.Number, "/src.jpg")
	a := DownloadPhotoConf{url}
	dc := Client{
		HttpClient: dhc,
	}
	return dc.DownloadPhoto(a)
}
