package main
import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().Unix())
    var xs = [50]int{}
    fmt.Println("Original array:")
    for i,_ := range(xs) {
        xs[i] = i
        fmt.Printf(" %v ", xs[i])
    }
    fmt.Println()

    // shuffle the array
    for i,_ := range(xs) {
        var tmp = xs[i]
        var j = i + rand.Intn(len(xs)-i)
        xs[i] = xs[j]
        xs[j] = tmp
    }

    fmt.Println("Shuffled array:")
    for i,_ := range(xs) {
        fmt.Printf(" %v ", xs[i])
    }
    fmt.Println()
}
