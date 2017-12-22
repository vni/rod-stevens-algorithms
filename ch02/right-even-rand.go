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

func (r *rand_st) right_even_rand(min, max int) int {
    x := r.rand()
    return min + int(float32(x)/float32(r.M) * float32(max-min))
}

func main() {
    r := newRand(7, 5, 11, 0)
    const SIZE = 4
    hist := [SIZE]int{}
    for i := 0; i < 1000; i++ {
        x := r.right_even_rand(0, SIZE)
        hist[x]++
    }

    for k,v := range(hist) {
        fmt.Println("hist[k: ", k, "]: ", v)
    }
}
