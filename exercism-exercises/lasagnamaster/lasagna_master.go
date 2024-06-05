package lasagnamaster

// Get preparation time based on number of layers and average prepartion time.
func PreparationTime(layers []string, minutes int) int {
	if minutes == 0 {
		minutes = 2
	}
	return len(layers) * minutes
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		} else if layers[i] == "sauce" {
			sauce += 0.2
		}
	}
	return
}

func AddSecretIngredient(list1, list2 []string) {
	var list1LastItem string = list1[len(list1)-1]
	var list2LastItemIndex int = len(list2) - 1
	list2[list2LastItemIndex] = list1LastItem
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	var result []float64 = make([]float64, len(quantities))
	copy(result, quantities)
	for i := 0; i < len(result); i++ {
		result[i] *= float64(portions) / 2
	}
	return result
}

func Run() {
	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	println("Preparation time (expected - 18):", PreparationTime(layers, 3))
	println("Preparation time (expected - 12):", PreparationTime(layers, 0))

	noodles, sauce := Quantities(layers)
	println("Quantities (expected - 100, 0.4):", noodles, "|", sauce)

	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella", "?"}
	AddSecretIngredient(friendsList, myList)
	println("My list: ")
	for i := 0; i < len(myList); i++ {
		print(myList[i], " ")
	}

	println("")

	quantities := []float64{1.2, 3.6, 10.5}
	scaledQuantities := ScaleRecipe(quantities, 1)
	print("Scaled quantities: ")
	for i := 0; i < len(scaledQuantities); i++ {
		print(scaledQuantities[i], " (orig_input: ", quantities[i], ") ", "| ")
	}
}
