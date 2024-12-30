package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

func KmToMi(km float64) float64 {
    return km * 0.621371
}
func MiToKm(mi float64) float64 {
    return mi / 0.621371
}
func KgToLb(kg float64) float64 {
    return kg * 2.20462
}
func LbToKg(lb float64) float64 {
    return lb / 2.20462
}
func CToF(c float64) float64 {
    return c * 9/5 + 32
}
func FToC(f float64) float64 {
    return (f - 32) * 5/9
}

func main() {
    //Sprawdzanie 
    if len(os.Args) != 3 {
        fmt.Println(" _______________________________________________________________________________")
        fmt.Println("|Użycie: go run convert.go <wartość> <typ_konwersji>                            |")
        fmt.Println("|_______________________________________________________________________________|")
        fmt.Println("|Dostępne typy konwersji: km_na_mi, mi_na_km, kg_na_lb, lb_na_kg, c_na_f, f_na_c|")
        fmt.Println("|_______________________________________________________________________________|")
        fmt.Println("|Przykład użycia: go run convert.go 100 km_na_mi                                |")
        fmt.Println("|_______________________________________________________________________________|")
        return
    }

    //Pobierz argumenty
    valueStr := os.Args[1]
    conversionType := os.Args[2]

    //Przekształć wartość na float
    value, err := strconv.ParseFloat(valueStr, 64)
    if err!= nil {
        fmt.Printf("Błędna wartość: %s\n", valueStr)
        return
    }

    //Przekształć konwersję
    var result float64
    switch strings.ToLower(conversionType) {
    case "km_na_mi":
        result = KmToMi(value)
    case "mi_na_km":
        result = MiToKm(value)
    case "kg_na_lb":
        result = KgToLb(value)
    case "lb_na_kg":
        result = LbToKg(value)
    case "c_na_f":
        result = CToF(value)
    case "f_na_c":
        result = FToC(value)
    default:
        fmt.Println("Nieznany typ konwersji:", conversionType)
        fmt.Printf("Dostępne typy konwersji: km_na_mi, mi_na_km, kg_na_lb, lb_na_kg, c_na_f, f_na_c")
    }
    //Wyświetl wynik
    fmt.Printf("Wynik: %.2f\n", result)
}
