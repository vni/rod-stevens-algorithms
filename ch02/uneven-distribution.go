package main
import (
    "fmt"
    "time"
    "math/rand"
)

func d6() int {
    return 1 + rand.Intn(6);
}

func d12() int {
    return 1 + rand.Intn(12);
}

func main() {
    rand.Seed(time.Now().Unix())
    d6x2_hist := [13]int{}
    d12_hist := [13]int{}
    for i := 0; i < 10e5; i++ {
        d6x2_hist[d6()+d6()]++;
        d12_hist[d12()]++;
    }

    fmt.Println("d6x2_hist:")
    for i,v := range(d6x2_hist) {
        fmt.Printf("d6x2_hist[i:%d]: %v\n", i, v)
    }
    fmt.Println("\nd12_hist:")
    for i,v := range(d12_hist) {
        fmt.Printf("d12_hist[i:%d]: %v\n", i, v)
    }
}
