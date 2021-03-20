package utils

/*
Every image returned by the Unsplash API is a dynamic image URL, which means that it can be manipulated to create
new transformations of the image by simply adjusting the query parameters of the image URL.

This enables resizing, cropping, compression, and changing the format of the image in realtime client-side, without any
API calls.

Under the hood, Unsplash uses Imgix, a powerful image manipulation service to provide dynamic image URLs.
*/

import (
	"fmt"

	"github.com/eddogola/unsplash-go/unsplash/client"
)

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

// GetResizedPhotoURL takes a picture and a ResizeOptions object, dynamically resizes Photo.URLs.Raw using the options
// and returns the resulting URL
func GetResizedPhotoURL(pic *client.Photo, rOptions ResizeOptions) string {
	return pic.URLs.Raw + rOptions.String()
}
