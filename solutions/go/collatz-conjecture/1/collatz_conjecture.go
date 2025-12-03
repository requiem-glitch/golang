package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	cnt := 0
    for cnt = 0; n != 1; cnt++ {
        if cnt >= 200 {
            return 0, errors.New("sosat;")
        }
        if n % 2 == 0 {
        	n /= 2
    	} else {
        	n = n * 3 + 1
    	}
    }
    return cnt, nil
    // panic("Please implement the CollatzConjecture function")
}
