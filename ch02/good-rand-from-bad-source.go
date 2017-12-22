package main
import "fmt"
import "math/rand"
import "time"

// bad coin random source
// gives 0 in 10% and 1 in 90%
func bad_coin() int {
    if rand.Intn(10) < 1 {
        return 0
    } else {
        return 1
    }
}

// returns 0 in 50% of cases
// returns 1 in another 50% of cases
func good_rand_from_bad_coin() int {
    for {
        a := bad_coin()
        b := bad_coin()
        if a != b {
            return a;
        }
    }
}

func main() {
    rand.Seed(time.Now().Unix())

    bchist := [2]int{}
    grhist := [2]int{}

    for i := 0; i < 1000000; i++ {
        x := bad_coin()
        bchist[x]++
    }
    fmt.Printf("bad coint in 1000000 drops:\n")
    fmt.Printf("bad coin hist: [0]: %v\n", bchist[0])
    fmt.Printf("bad coin hist: [1]: %v\n", bchist[1])

    for i := 0; i < 1000000; i++ {
        x := good_rand_from_bad_coin()
        grhist[x]++
    }
    fmt.Printf("\n\ngood rand in 1000000 drops:\n");
    fmt.Printf("good rand hist: [0]: %v\n", grhist[0])
    fmt.Printf("good rand hist: [1]: %v\n", grhist[1])
}
