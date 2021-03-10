package unsplash

import "fmt"

// ResizeOptions defines parameters that resize a photo
// Resizing is done using Imgix()
type ResizeOptions struct {
	Width       string
	Height      string
	Crop        string // https://docs.imgix.com/apis/rendering/size/crop
	ImageFormat string // https://docs.imgix.com/apis/rendering/format/fm
	Auto        string // https://docs.imgix.com/apis/rendering/auto/auto
	Quality     string //for changing the compression quality when using lossy file formats
	// https://docs.imgix.com/apis/rendering/format/q
	Fit string // for changing the fit of the image within the specified dimensions
	// https://docs.imgix.com/apis/rendering/size/fit
	DevicePixelRatio string // min-value is 1, max-value is 5
	//https://docs.imgix.com/apis/rendering/pixel-density/dpr
}

// NewDefaultResizeOptions takes in width and height to return a new default ResizeOptions
// with custom dimensions
func NewDefaultResizeOptions(w, h int) *ResizeOptions {
	width := fmt.Sprint(w)
	height := fmt.Sprint(h)
	return &ResizeOptions{Width: width,
		Height: height,
		Auto:   "format"}
}

func (rOptions ResizeOptions) String() string {
	var result string
	options := map[string]string{
		"w":    rOptions.Width,
		"h":    rOptions.Height,
		"crop": rOptions.Crop,
		"fm":   rOptions.ImageFormat,
		"auto": rOptions.Auto,
		"q":    rOptions.Quality,
		"fit":  rOptions.Fit,
		"dpr":  rOptions.DevicePixelRatio,
	}
	for key, val := range options {
		if val != "" {
			result += fmt.Sprintf("&%v=%v", key, val)
		}
	}
	return result
}

// Resize takes a ResizeOptions object, dynamically resizes Photo.URLs.Raw using the options
// and returns the resulting URL
func (pic *Photo) Resize(rOptions ResizeOptions) string {
	return pic.URLs.Raw + rOptions.String()
}
