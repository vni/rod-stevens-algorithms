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

// [min; max)
func (r *rand_st) right_even_rand_2(min, max int) int {
    var x int
    for {
        x = r.rand()
        if min <= x && x < max {
            break;
        }
    }
    return x
}

func main() {
    r := newRand(7, 5, 11, 0)
    const SIZE = 4
    hist := [5]int{}
    for i := 0; i < 1000; i++ {
        x := r.right_even_rand_2(2, 5)
        hist[x]++
    }

    for i := 2; i < 5; i++ {
        fmt.Println("hist[k: ", i, "]: ", hist[i])
    }
}
