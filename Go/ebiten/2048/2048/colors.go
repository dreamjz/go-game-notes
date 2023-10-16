package twenty48

import "image/color"

var (
	backgroundColor = color.RGBA{R: 0xfa, G: 0xf8, B: 0xef, A: 0xff}
	frameColor      = color.RGBA{R: 0xbb, G: 0xad, B: 0xa0, A: 0xff}
)

func tileColor(value int) color.Color {
	switch value {
	case 2, 4:
		return color.RGBA{R: 0x77, G: 0x6e, B: 0x65, A: 0xff}
	case 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536:
		return color.RGBA{R: 0xf9, G: 0xf6, B: 0xf2, A: 0xff}
	}
	panic("not reach")
}

func tileBackgroundColor(value int) color.Color {
	switch value {
	case 0:
		return color.NRGBA{R: 0xee, G: 0xe4, B: 0xda, A: 0x59}
	case 2:
		return color.RGBA{R: 0xee, G: 0xe4, B: 0xda, A: 0xff}
	case 4:
		return color.RGBA{R: 0xed, G: 0xe0, B: 0xc8, A: 0xff}
	case 8:
		return color.RGBA{R: 0xf2, G: 0xb1, B: 0x79, A: 0xff}
	case 16:
		return color.RGBA{R: 0xf5, G: 0x95, B: 0x63, A: 0xff}
	case 32:
		return color.RGBA{R: 0xf6, G: 0x7c, B: 0x5f, A: 0xff}
	case 64:
		return color.RGBA{R: 0xf6, G: 0x5e, B: 0x3b, A: 0xff}
	case 128:
		return color.RGBA{R: 0xed, G: 0xcf, B: 0x72, A: 0xff}
	case 256:
		return color.RGBA{R: 0xed, G: 0xcc, B: 0x61, A: 0xff}
	case 512:
		return color.RGBA{R: 0xed, G: 0xc8, B: 0x50, A: 0xff}
	case 1024:
		return color.RGBA{R: 0xed, G: 0xc5, B: 0x3f, A: 0xff}
	case 2048:
		return color.RGBA{R: 0xed, G: 0xc2, B: 0x2e, A: 0xff}
	case 4096:
		return color.NRGBA{R: 0xa3, G: 0x49, B: 0xa4, A: 0x7f}
	case 8192:
		return color.NRGBA{R: 0xa3, G: 0x49, B: 0xa4, A: 0xb2}
	case 16384:
		return color.NRGBA{R: 0xa3, G: 0x49, B: 0xa4, A: 0xcc}
	case 32768:
		return color.NRGBA{R: 0xa3, G: 0x49, B: 0xa4, A: 0xe5}
	case 65536:
		return color.NRGBA{R: 0xa3, G: 0x49, B: 0xa4, A: 0xff}
	}
	panic("not reach")
}
