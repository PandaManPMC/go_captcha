package rotate

import (
	"github.com/PandaManPMC/pic_captcha/captcha/base/option"
)

// defaultOptions is to the default configuration
func defaultOptions() Option {
	return func(opts *Options) {
		opts.imageSquareSize = 220
		opts.rangeAnglePos = []*option.RangeVal{
			{Min: 30, Max: 330},
		}

		opts.thumbImageAlpha = 1
		opts.rangeThumbImageSquareSize = []int{140, 150, 160, 170}
	}
}

// defaultResource is to the default resource
func defaultResource() Resource {
	return func(resources *Resources) {
		// ...
	}
}
