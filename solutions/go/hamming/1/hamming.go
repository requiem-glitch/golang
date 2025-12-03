package hamming

import "errors"

func Distance(a, b string) (int, error) {
	cnt := 0
    if len(a) != len(b) {
        return 0, errors.New("lens of strings must be similiar")
    }
    for i, _ := range a {
        if a[i] != b[i] {
            cnt++
        }
    }
    return cnt, nil
    // panic("Implement the Distance function")
}
