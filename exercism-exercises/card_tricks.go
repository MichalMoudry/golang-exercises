package main

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	println(index, len(slice))
	if index >= len(slice) || index < 0 {
		return -1
	}
	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index >= len(slice) {
		return append(slice, value)
	}
	slice[index] = value
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, value ...int) []int {
	var res []int = value
	return append(res, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index > len(slice) || index < 0 {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	cards := FavoriteCards()
	println("Favourite cards:", cards)
	println("Get item:", GetItem(cards, 3))
	println("Set item:", SetItem(cards, 0, 504)[0])
	println("Prepend items:", PrependItems(cards, 100)[1])
	println("--- All cards ---")
	for i := 0; i < len(cards); i++ {
		println(cards[i])
	}
	println("Get item:", GetItem([]int{5, 2, 10, 6, 8, 7, 0, 9}, 8))
	println("Set item:", SetItem([]int{5, 2, 10, 6, 8, 7, 0, 9}, -1, 8))
}
