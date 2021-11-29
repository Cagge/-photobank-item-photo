package hendler

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/gographics/imagick.v2/imagick"
	"photobank-item-photo/app/pkg/downloader"
	"photobank-item-photo/app/pkg/resize"
	. "photobank-item-photo/app/pkg/resize/models"
)

type httpClientH interface {
	downloader.HttpClientI
}
var ErrDownload = errors.New("error download photo")

func ResizePhoto(p PhotoDownloadCfg, cc ConfigResize, b []resize.ConfigResize, count int,mw *imagick.MagickWand,dhc httpClientH ) (int, [][]byte, error) {
	byti, err := ProcessImageByUrl(p,dhc)
	switch {
	case errors.Is(downloader.ErrStatusCode, err):
		return count, nil, ErrDownload
	case err != nil:
		return count, nil, err
	default:
	}
	byter := make([][]byte, len(cc.Resize.OutSize))
	m := resize.ReadImageBlob{byti}
	err = m.ReadImageBlob(mw)
	if err != nil {
		return count, byter, err
	}
	for num, size := range b {
		mwc := mw.Clone()
		err = size.Resize(mwc)
		if err != nil {
			return count, byter, err
		}
		byter[num] = mwc.GetImageBlob()
		count++
		fmt.Println(fmt.Sprint(cc.Resize.PathSavePhoto, p.Sid, "_", p.Number, "_", size.Cols, ".jpg"))
		fmt.Println(count, "img")
		mwc.Destroy()
	}
	mw.Destroy()
	return count, byter, err
}