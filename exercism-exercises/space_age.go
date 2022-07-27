package main

type Planet string

func Age(seconds float64, planet Planet) float64 {
	earthYears := (((seconds / 60) / 60) / 24) / 365.25
	switch planet {
	case "Mercury":
		return earthYears / 0.2408467
	case "Venus":
		return earthYears / 0.61519726
	case "Mars":
		return earthYears / 1.8808158
	case "Jupiter":
		return earthYears / 11.862615
	case "Saturn":
		return earthYears / 29.447498
	case "Uranus":
		return earthYears / 84.016846
	case "Neptune":
		return earthYears / 164.79132
	case "Earth":
		return earthYears
	default:
		return -1
	}
}

func main() {
	var mercury Planet = "Mercury"
	println("Mercury (expected - 280.88):", Age(2134835688.0, mercury))
	println("Earth (expected - 31.69):", Age(1000000000.0, "Earth"))
}
