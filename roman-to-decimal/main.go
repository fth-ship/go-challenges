package main

import "fmt"

func hasValue ( expression string, value string ) bool {
    out := false

    for i, _ := range expression {
        l := i + len( value )
        if l <= len( expression ) && value == expression[ i : l ] {
            out = true
            break
        }
    }

    return out
}

func match ( value string ) int {
    out := 0

    tokens := map [ string ] int {
        "IX": 2,
        "IV": 2,
        "XL": 20,
        "XC": 20,
        "CD": 200,
        "CM": 200,
    }

    for k, v := range tokens {
        if hasValue( value, k ) {
            out = v
            break
        }
    }

    return out
}

func toDecimal ( value string ) int {
    out := 0

    tokens := map [string] int {
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }

    for _, v := range value {
        out += tokens[ string(v) ]
    }

    out -= match( value )

    return out
}

func main () {
    fmt.Println( toDecimal("I") )
    fmt.Println( toDecimal("II") )
    fmt.Println( toDecimal("III") )
    fmt.Println( toDecimal("IV") )
    fmt.Println( toDecimal("V") )
    fmt.Println( toDecimal("VI") )
    fmt.Println( toDecimal("VII") )
    fmt.Println( toDecimal("VIII") )
    fmt.Println( toDecimal("IX") )
    fmt.Println( toDecimal("X") )
}
