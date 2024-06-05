package nfs

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		battery:      100,
		distance:     0,
		batteryDrain: batteryDrain,
	}
}

type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery > 0 && car.battery >= car.batteryDrain {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	if car.batteryDrain <= 0 {
		return true
	}
	var drivableDistance int
	for {
		car = Drive(car)
		println("Battery:", car.battery, "|", "Distance:", car.distance)
		if car.battery < car.batteryDrain {
			drivableDistance = car.distance
			break
		}
	}
	println("Car is able to drive:", drivableDistance)
	if drivableDistance >= track.distance {
		return true
	} else {
		return false
	}
}

func Run() {
	car := NewCar(5, 2)
	track := NewTrack(100)

	println("Can finish (expected - true):", CanFinish(car, track))
}
