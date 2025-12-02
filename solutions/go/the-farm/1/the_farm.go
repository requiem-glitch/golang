package thefarm

import (
    "fmt"
    "errors"
)

type InvalidCowsError struct{
    numCows int
    mes string
}

func (e *InvalidCowsError) Error() string{
    return fmt.Sprintf("%d cows are invalid: %s", e.numCows, e.mes)
}

// TODO: define the 'DivideFood' function
func DivideFood(fd FodderCalculator, numCows int) (float64, error) {
    i, err := fd.FodderAmount(numCows)
    if err != nil {
        return 0.0, err
    }
    l, err := fd.FatteningFactor()
    if err != nil {
        return 0.0, err
    }
    return i*l/float64(numCows), nil
}
// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fd FodderCalculator, numCows int) (float64, error) {
    if numCows > 0 {
        i, err := DivideFood(fd, numCows)
        return i, err
    }
    return 0.0, errors.New("invalid number of cows")
}
// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numCows int) error {
    if numCows < 0 {
        return &InvalidCowsError{
            numCows: numCows,
            mes: "there are no negative cows",
        }
    }
    if numCows == 0 {
        return &InvalidCowsError{
            numCows: numCows,
            mes: "no cows don't need food",
        }
    }
    return nil
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
