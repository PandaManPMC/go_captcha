package yrdzst

import (
	"github.com/PandaManPMC/pic_captcha/captcha/bindata/fonts/yrdzst"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

func GetFont() (font *truetype.Font, err error) {
	asset, err := yrdzst.Asset("sourcedata/fonts/yrdzst-bold.ttf")
	if err != nil {
		return font, err
	}

	font, err = freetype.ParseFont(asset)
	if err != nil {
		return nil, err
	}

	return font, nil
}
