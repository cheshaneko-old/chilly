package gpiofarm

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
    "github.com/d2r2/go-i2c"
	"time"
)

type GpioFarm struct {
	Name string
}

var (
	// Use mcu pin 10, corresponds to physical pin 19 on the pi
	pin = rpio.Pin(10)
	pin2 = rpio.Pin(9)

    pwmPin = rpio.Pin(19)

    pinGreen = rpio.Pin(17)
    pinViolet = rpio.Pin(27)
    pinOrange = rpio.Pin(22)
    pinBlue = rpio.Pin(11)

)

const WATER_TIME = 3

func (f *GpioFarm) OnLight() error {
	pin.High()
	pin2.High()
	return nil
}

func (f *GpioFarm) OffLight() error {
	pin.Low()
	pin2.Low()
	return nil
}

func (f *GpioFarm) Open() error {
    if err := rpio.Open(); err != nil {
		return err
	}
	pin.Output()
	pin2.Output()
	pinGreen.Output()
	pinViolet.Output()
	pinOrange.Output()
	pinBlue.Output()
	pin.Low()
	pin2.Low()
	pinGreen.Low()
	pinViolet.Low()
	pinOrange.Low()
	pinBlue.Low()
	return nil
}

func (f *GpioFarm) Close() error {
	pin.Low()
	pin2.Low()
	return rpio.Close()
}

func (f *GpioFarm) OnMotor() error {
    fmt.Println("OnMotor")
    pwmPin.Mode(rpio.Pwm)
    pwmPin.Freq(6000)
    pwmPin.DutyCycle(6, 8)
    return nil
}
func (f *GpioFarm) OffMotor() error {
    fmt.Println("OffMotor")
    pwmPin.Mode(rpio.Pwm)
    pwmPin.Freq(6000)
    pwmPin.DutyCycle(0, 8)
	pwmPin.Output()
	pwmPin.Low()
    return nil
}
func (f *GpioFarm) WaterGreen() error {
    fmt.Println("Water green")
    pwmPin.DutyCycle(8, 8)
	pinGreen.High()
	time.Sleep(WATER_TIME * time.Second)
    pwmPin.DutyCycle(6, 8)
	pinGreen.Low()
    return nil
}
func (f *GpioFarm) WaterViolet() error {
    fmt.Println("Water violet")
    pwmPin.DutyCycle(8, 8)
	pinViolet.High()
	time.Sleep(WATER_TIME * time.Second)
    pwmPin.DutyCycle(6, 8)
	pinViolet.Low()
    return nil
}
func (f *GpioFarm) WaterOrange() error {
    fmt.Println("Water orange")
    pwmPin.DutyCycle(8, 8)
	pinOrange.High()
	time.Sleep(WATER_TIME * time.Second)
    pwmPin.DutyCycle(6, 8)
	pinOrange.Low()
    return nil
}
func (f *GpioFarm) WaterBlue() error {
    fmt.Println("Water blue")
    pwmPin.DutyCycle(8, 8)
	pinBlue.High()
	time.Sleep(WATER_TIME * time.Second)
    pwmPin.DutyCycle(6, 8)
	pinBlue.Low()
    return nil
}
func (f *GpioFarm) TemperatureAndHumidityRoom() (uint, uint, error) {
    ms, err := i2c.NewI2C(0x44, 1)
	if err != nil {
		return 0, 0, err
	}
    defer ms.Close()
    _, err = ms.WriteBytes([]byte{0x2400 >> 8 , 0x2400 & 0xFF})
    if err != nil {
		return 0, 0, err
    }
    buf1 := make([]byte, 6)
    _, err = ms.ReadBytes(buf1)
	if err != nil {
		return 0, 0, err
	}
    t := ((uint(buf1[0]) * 256.0 + uint(buf1[1])) * 175.0) / 65535.0 - 45.0
    h := ((uint(buf1[3]) * 256.0 + uint(buf1[4])) * 100.0) / 65535.0

    return t, h, nil
}
