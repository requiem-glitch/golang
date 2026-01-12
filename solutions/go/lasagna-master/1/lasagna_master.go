package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTimePerLayer int) (totalTime int) {
    if avgTimePerLayer == 0 {
        avgTimePerLayer = 2
    }
    totalTime = len(layers) * avgTimePerLayer
    return
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
    for i := 0; i < len(layers); i++ {
        if layers[i] == "noodles" {
            noodles++
        } else if layers[i] == "sauce" {
            sauce++
        }
    } 
    noodles *= 50
    sauce *= 0.2
    return
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsLayers, myLayers []string) {
    myLayers[len(myLayers) - 1] = friendsLayers[len(friendsLayers) - 1]
    // myLayers = append(myLayers, friendsLayers[len(friendsLayers) - 1])
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, targetPortions int) (amountsNeeded []float64) {
    for i := 0; i < len(amounts); i++ {
        amountsNeeded = append(amountsNeeded, amounts[i] * (float64(targetPortions) / 2))
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
