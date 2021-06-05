package utils

type Output interface {
	// Write takes in group of points to be written to the Output
	Write(points *[]*PKGContent) error
}

type Input interface {
	// Gather takes in an accumulator and adds the inputInfo to the Input
	Gather() (interface{}, error)
}

type PKGContent struct {
	// gorm.Model
	Name     string `gorm:"primary_key"`
	Parent   string `gorm:"primary_key"`
	Synopsis string
	Href     string
}
