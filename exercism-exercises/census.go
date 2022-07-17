package main

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	street, _ := r.Address["street"]
	if r.Name == "" || r.Address == nil || street == "" {
		return false
	}
	return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	var result int = 0
	for i := 0; i < len(residents); i++ {
		if residents[i].HasRequiredInfo() {
			result++
		}
	}
	return result
}

func main() {
	name := "Matthew Sanabria"
	age := 29
	address := map[string]string{"street": "Main St."}
	newResident := NewResident(name, age, address)

	println(newResident.Name, newResident.Age, len(newResident.Address))

	println("Has required information (expected - true):", newResident.HasRequiredInfo())

	newResident.Delete()
	println(newResident.Name, newResident.Age, len(newResident.Address))
	println("Has required information (expected - false):", newResident.HasRequiredInfo())

	newResident2 := NewResident("Hossein", 30, map[string]string{"street": ""})
	println(newResident2.HasRequiredInfo())
}
