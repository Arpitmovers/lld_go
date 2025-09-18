package main

type VechileType string

const (
	TwoWheel  VechileType = "2W"
	FourWheel VechileType = "4W"
)

type Vechile struct {
	size    VechileType
	license string
}
