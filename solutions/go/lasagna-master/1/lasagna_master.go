package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPreparationTime int) int {
	if avgPreparationTime == 0 {
		return len(layers) * 2
	}
	return len(layers) * avgPreparationTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	var countSauce, countNoodles int
	for _, val := range layers {
		if val == "sauce" {
			countSauce++
		} else if val == "noodles" {
			countNoodles++
		}
	}
	noodles = countNoodles * 50
	sauce = float64(countSauce) * 0.2
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) (scaledQuantities []float64) {
	var extaPortions float64 = float64(portions) / 2
	for _, val := range quantities {
		scaledQuantities = append(scaledQuantities, val*extaPortions)
	}
	return
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
