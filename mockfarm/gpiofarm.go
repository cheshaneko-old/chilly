package mockfarm

import (
	"fmt"
)

type MockFarm struct {
	Name string
}


func (f *MockFarm) OnLight() error {
	fmt.Println("OnLight")
	return nil
}

func (f *MockFarm) OffLight() error {
	fmt.Println("OffLight")
	return nil
}

func (f *MockFarm) Open() error {
	fmt.Println("Open")
	return nil
}

func (f *MockFarm) Close() error {
	fmt.Println("Close")
	return nil
}

func (f *MockFarm) OnMotor() error {
	fmt.Println("OnMotor")
	return nil
}
func (f *MockFarm) OffMotor() error {
	fmt.Println("OffMotor")
	return nil
}
func (f *MockFarm) WaterGreen() error {
	fmt.Println("Water green")
	return nil
}
func (f *MockFarm) WaterViolet() error {
	fmt.Println("Water violet")
	return nil
}
func (f *MockFarm) WaterOrange() error {
	fmt.Println("Water orange")
	return nil
}
func (f *MockFarm) WaterBlue() error {
	fmt.Println("Water blue")
	return nil
}
func (f *MockFarm) TemperatureAndHumidityRoom() (uint, uint, error) {
	fmt.Println("Temperatur room")
	return 42, 24, nil
}

