package main
import "fmt"

type rand_st struct {
    A, B, M, X int
}

func newRand(a, b, m, x int) *rand_st {
    return &rand_st{
        a, b, m, x,
    }
}

func (r *rand_st) rand() int {
    r.X = (r.A*r.X + r.B) % r.M
    return r.X
}

func main() {
    r := newRand(7, 5, 11, 0)
    hist := [11]int{}
    for i := 0; i < 1000; i++ {
        x := r.rand()
        hist[x]++
    }

    for i,v := range(hist) {
        fmt.Println("hist[i: ", i, "]: ", v)
    }
}
