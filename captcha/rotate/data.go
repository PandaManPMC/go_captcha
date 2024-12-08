package rotate

import "github.com/PandaManPMC/pic_captcha/captcha/base/imagedata"

// CaptchaData .
type CaptchaData interface {
	GetData() *Block
	GetMasterImage() imagedata.PNGImageData
	GetThumbImage() imagedata.PNGImageData
}

// CaptData .
type CaptData struct {
	block       *Block
	masterImage imagedata.PNGImageData
	thumbImage  imagedata.PNGImageData
}

var _ CaptchaData = (*CaptData)(nil)

// GetData is to get block
func (c CaptData) GetData() *Block {
	return c.block
}

// GetMasterImage is to get master image
func (c CaptData) GetMasterImage() imagedata.PNGImageData {
	return c.masterImage
}

// GetThumbImage is to get thumb image
func (c CaptData) GetThumbImage() imagedata.PNGImageData {
	return c.thumbImage
}
