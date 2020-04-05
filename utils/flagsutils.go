package utils

import (
	"image"
	"math"

	"github.com/nfnt/resize"
)

// ResizeFlags function resizes two images to a specified size
func ResizeFlags(flag1, flag2 *image.Image, minWidth, minHeight int) {
	*flag1 = resize.Resize(uint(minWidth), uint(minHeight), *flag1, resize.NearestNeighbor)
	*flag2 = resize.Resize(uint(minWidth), uint(minHeight), *flag2, resize.NearestNeighbor)
}

// GetFlagsMinSize function claculate the minimum size of the rectangle which contains the two flags
func GetFlagsMinSize(flag1, flag2 *image.Image) (minWidth, minHeight int) {
	firstFlagBounds, secondFlagBounds := GetFlagsBounds(flag1, flag2)

	minWidth = int(math.Min(float64(firstFlagBounds.Dx()), float64(secondFlagBounds.Dx())))
	minHeight = int(math.Min(float64(firstFlagBounds.Dy()), float64(secondFlagBounds.Dy())))
	return
}

// GetFlagsBounds function returns the bounds of two flags
func GetFlagsBounds(flag1, flag2 *image.Image) (firstFlagBounds, secondFlagBounds image.Rectangle) {
	firstFlagBounds, secondFlagBounds = (*flag1).Bounds(), (*flag2).Bounds()
	return
}

// AdaptFlags function resizes two flags to the same size
func AdaptFlags(flag1, flag2 *image.Image) (minWidth, minHeight int) {
	minWidth, minHeight = GetFlagsMinSize(flag1, flag2)
	ResizeFlags(flag1, flag2, minWidth, minHeight)
	return
}
