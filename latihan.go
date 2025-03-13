package main
import "fmt"

func prima(masukan, i int) bool {
    if masukan < 2 { 
        return false
    }
    if i == masukan { 
        return true
    }
    if masukan%i == 0 { 
        return false
    }
    return prima(masukan, i+1)
}

func main() {
    var d1 int
    fmt.Scan(&d1)
    if prima(d1, 2) {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}