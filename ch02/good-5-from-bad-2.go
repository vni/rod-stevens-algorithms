package main
import ("fmt"; "time"; "math/rand")

// bad random source. Generates 0 in 10%, 1 in 90%
func bad2() int {
    if rand.Intn(10) == 0 {
        return 0
    }
    return 1
}

// use bad source to generate good random numbers [0; 2)
func good2() int {
    for {
        a := bad2()
        b := bad2()
        if a != b {
            return a
        }
    }
}

// rand 0..4. Use good random source
func good5() int {
    for {
        a := good2()
        b := good2()
        c := good2()
        /* If bad2 is used instead of good2 => the histogram is very bad
        a := bad2()
        b := bad2()
        c := bad2()
        */
        res := a*4 + b*2 + c
        if res < 5 {
            return res
        }
    }
}

func main() {
    rand.Seed(time.Now().Unix())

    bad2_hist := [2]int{}
    good2_hist := [2]int{}
    good5_hist := [5]int{}

    for i:=0; i<100000; i++ {
        x := bad2()
        y := good2()
        z := good5()

        bad2_hist[x]++
        good2_hist[y]++
        good5_hist[z]++
    }

    fmt.Printf("IN 100000 DROPS:\n")
    fmt.Printf("\nbad2_hist[0]: %v\n", bad2_hist[0])
    fmt.Printf("bad2_hist[1]: %v\n", bad2_hist[1])

    fmt.Printf("\ngood2_hist[0]: %v\n", good2_hist[0])
    fmt.Printf("good2_hist[1]: %v\n", good2_hist[1])

    fmt.Printf("\ngood5_hist[0]: %v\n", good5_hist[0])
    fmt.Printf("good5_hist[1]: %v\n", good5_hist[1])
    fmt.Printf("good5_hist[2]: %v\n", good5_hist[2])
    fmt.Printf("good5_hist[3]: %v\n", good5_hist[3])
    fmt.Printf("good5_hist[4]: %v\n", good5_hist[4])
}
