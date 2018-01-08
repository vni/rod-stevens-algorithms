package main
import (
    "fmt"
    "os"
    "strconv"
)

func gcd(a, b int) int {
    for b != 0 {
        var r = a % b
        a = b
        b = r
    }
    return a
}

func main() {
    if len(os.Args) != 3 {
        fmt.Fprintf(os.Stderr, "Usage: %s num1 num2\n", os.Args[0]);
        return
    }

    num1, err1 := strconv.Atoi(os.Args[1])
    if err1 != nil {
        fmt.Fprintf(os.Stderr, "failed to strconv.Atoi args[1]\n")
        return
    }
    num2, err2 := strconv.Atoi(os.Args[2])
    if err2 != nil {
        fmt.Fprintf(os.Stderr, "failed to strconv.Atoi args[2]\n")
        return
    }

    fmt.Printf("gcd(%v, %v) is %v\n", num1, num2, gcd(num1, num2))
}
