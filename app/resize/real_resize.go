package resize

import (
	"github.com/pkg/errors"
	"gopkg.in/gographics/imagick.v2/imagick"
	"photobank-item-photo/app/pkg/downloader"
	. "photobank-item-photo/app/pkg/downloader/models"
	"photobank-item-photo/app/pkg/resize"
	. "photobank-item-photo/app/pkg/resize/models"
	. "photobank-item-photo/app/resize/hendler"
)

type httpClientI interface {
	downloader.HttpClientI
}

func ControlResize(a ConfigDataPhoto, c ConfigResize, dhc httpClientI) (int, [][]byte, error) { //return count photo
	count := 0
	byter := make([][]byte, len(c.Resize.OutSize), cap(c.Resize.OutSize))
	b := make([]resize.ConfigResize, len(c.Resize.OutSize), cap(c.Resize.OutSize))
	dc := downloader.Client{
		HttpClient: dhc,
	}
	f, err := dc.GetJsonPhoto(a)
	if err != nil {
		return 0, byter, err
	}
	for num, size := range c.Resize.OutSize {
		err := ControlSize(size)
		if err != nil {
			return 0, byter, err
		}
		b[num] = resize.ConfigResize{size, size, imagick.FILTER_LANCZOS, 1}
	}
	for i := range f.Photos {
		m := PhotoDownloadCfg{a.Photodate.Urldownload, f.Photos[i].Sid, f.Photos[i].Number}
		mw := imagick.NewMagickWand()
		count, byter, err = ResizePhoto(m, c, b, count, mw, dhc)
		if (errors.Is(ErrDownload, err))||(errors.Is(ErrValueSid, err)) {
			continue
		} else if err != nil {
			return count, byter, err
		}
	}
	return count, byter, nil
}
