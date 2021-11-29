package resize

import (
	"gopkg.in/gographics/imagick.v2/imagick"
)

type MagickWand interface {
	ResizeImage(uint, uint, imagick.FilterType, float64) error
	ReadImageBlob([]byte) error
}
type ConfigResize struct {
	Cols   uint
	Rows   uint
	Filter imagick.FilterType
	Blur   float64
}
type ReadImageBlob struct {
	Byti []byte
}

func (r *ReadImageBlob) ReadImageBlob(mw MagickWand) error {
	return mw.ReadImageBlob(r.Byti)
}
func (r *ConfigResize) Resize(mw MagickWand) error {
	return mw.ResizeImage(r.Cols, r.Rows, r.Filter, r.Blur)
}
