package main

import "fmt"

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

func (car *Car) ToString() string {
	return fmt.Sprintf(
		"{speed: %d, batteryDrain: %d, battery: %d, distance: %d}",
		car.speed, car.batteryDrain, car.battery, car.distance)
}

func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		battery:      100,
		distance:     0,
		batteryDrain: batteryDrain,
	}
}

func (car *Car) Drive() {
	if car.battery > 0 && car.battery >= car.batteryDrain {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
}

func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

func (car *Car) CanFinish(trackDistance int) bool {
	if car.batteryDrain <= 0 {
		return true
	}
	var drivableDistance int
	for {
		car.Drive()
		if car.battery < car.batteryDrain {
			drivableDistance = car.distance
			break
		}
	}
	if drivableDistance >= trackDistance {
		return true
	} else {
		return false
	}
}

func main() {
	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)
	//car.Drive()
	trackDistance := 100
	println(car.ToString())
	println(car.DisplayDistance())
	println(car.DisplayBattery())
	println("Can car finish (expected - true):", car.CanFinish(trackDistance))
}
