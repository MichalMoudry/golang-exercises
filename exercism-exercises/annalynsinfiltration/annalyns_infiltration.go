package annalynsinfiltration

// CanFastAttack can be executed only when the knight is sleeping
func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

// CanSpy can be executed if at least one of the characters is awake
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	if knightIsAwake || archerIsAwake || prisonerIsAwake {
		return true
	} else {
		return false
	}
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	if prisonerIsAwake && !archerIsAwake {
		return true
	} else {
		return false
	}
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if petDogIsPresent && !archerIsAwake {
		return true
	} else {
		if prisonerIsAwake && !knightIsAwake && !archerIsAwake {
			return true
		} else {
			return false
		}
	}
}

func Run() {
	println("Can fast attack:", CanFastAttack(true))
	println("Can fast attack:", CanFastAttack(false))
	println("Can spy:", CanSpy(false, true, false))
	println("Can signal prisoner:", CanSignalPrisoner(false, true))
	println("Can free prisoner (expected - false):", CanFreePrisoner(false, true, false, false))
	println("Can free prisoner (expected - true):", CanFreePrisoner(false, false, true, false))
	println("Can free prisoner (expected - false):", CanFreePrisoner(false, true, true, false))
}
