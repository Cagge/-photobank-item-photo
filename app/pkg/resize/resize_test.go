package resize

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gopkg.in/gographics/imagick.v2/imagick"
)

func TestReadFail(t *testing.T) {
	mw := imagick.NewMagickWand()
	m := ReadImageBlob{Byti: []byte{65, 66, 67, 226, 130, 172}}
	err := m.ReadImageBlob(mw)
	assert.NotNil(t, err)
}

func TestResizeFail(t *testing.T) {
	mw := imagick.NewMagickWand()
	m := ConfigResize{100, 100, 22, 1}
	err := m.Resize(mw)
	assert.NotNil(t, err)
}
