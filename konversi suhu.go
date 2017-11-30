// konversi suhu oleh : Hendra Permana
package main

import "fmt"

func main() {
	var val float64
	var pilih int
	fmt.Println("Konversi suhu\n[1] Celcius ke Fahrenheit\n[2] Celcius ke Reamur\n[3] Fahrenheit ke Celsius\n[4] Fahrenheit ke Reamur\n[5] Reamur ke Celcius\n[6] Reamur ke Fahrenheit\n[7] Celcius ke Kelvin\n[8] Fahrenheit ke Kelvin\n[9] Kelvin ke Fahrenheit\n[10] Kelvin ke Celcius\n[11] Kelvin ke reamur\n[12] reamur ke Kelvin")
	fmt.Scanf("%d", &pilih)
	if pilih == 1{
		fmt.Println("Masukan nilai Celcius :")
		fmt.Scanf("%f", &val)
		celciusToFahrenheit(val)
	} else if pilih == 2{
		fmt.Println("Masukan nilai celcius :")
		fmt.Scanf("%f", &val)
		celciusToReamur(val)
	} else if pilih == 3{
		fmt.Println("Masukan nilai fahrenheit :")
		fmt.Scanf("%f", &val)
		fahrenheitToCelcius(val)
	} else if pilih == 4{
		fmt.Println("Masukan nilai fahrenheit :")
		fmt.Scanf("%f", &val)
		fahrenheitToReamur(val)
	} else if pilih == 5{
		fmt.Println("Masukan nilai Reamur :")
		fmt.Scanf("%f", &val)
		reamurToCelcius(val)
	} else if pilih == 6{
		fmt.Println("Masukan nilai Reamur :")
		fmt.Scanf("%f", &val)
		reamurToFahrenheit(val)
	} else if pilih == 7{
		fmt.Println("Masukan nilai Celcius :")
		fmt.Scanf("%f", &val)
		celciusToKelvin(val)
	} else if pilih == 8{
		fmt.Println("Masukan nilai Fahrenheit :")
		fmt.Scanf("%f", &val)
		fahrenheitToKelvin(val)
	} else if pilih == 9{
		fmt.Println("Masukan nilai kelvin :")
		fmt.Scanf("%f", &val)
		kelvinToFahrenheit(val)
	} else if pilih == 10{
		fmt.Println("Masukan nilai kelvin :")
		fmt.Scanf("%f", &val)
		kelvinToCelcius(val)
	} else if pilih == 11{
		fmt.Println("Masukan nilai kelvin :")
		fmt.Scanf("%f", &val)
		kelvinToReamur(val)
	} else if pilih == 12{
		fmt.Println("Masukan nilai reamur :")
		fmt.Scanf("%f", &val)
		reamurToKelvin(val)
	}


}

func celciusToFahrenheit(c float64){
	f := (c / 5 * 9) + 32
	fmt.Println(c, "celcius =",f,"Fahrenheit")
}

func celciusToReamur(c float64){
	r := c / 5 * 4
	fmt.Println(c, "Celcius =", r,"Reamur")
}

func fahrenheitToCelcius(f float64){
	c := (f - 32) / 9 * 5
	fmt.Println(f, "fahrenheit =", c,"Celcius")
}

func fahrenheitToReamur(f float64){
	r := (f - 32) / 9 * 4
	fmt.Println(f, "fahrenheit =", r,"Reamur")
}

func reamurToCelcius(r float64){
	c := r / 4 * 5
	fmt.Println(r, "reamur =", c,"Celcius")
}

func reamurToFahrenheit(r float64){
	f := (r / 4 * 9) + 32
	fmt.Println(r, "reamur =", f,"fahrenheit")
}

func celciusToKelvin(c float64){
	k := c + 273
	fmt.Println(c, "celcius =", k,"kelvin")
}

func fahrenheitToKelvin(f float64){
	k := (f - 32) / (9 * 5) + 273
	fmt.Println(f, "fahrenheit =", k,"kelvin")
}

func kelvinToFahrenheit(k float64){
	f := (k - 273) / (5 * 9) + 32
	fmt.Println(k, "kelvin =", f,"fahrenheit")
}

func kelvinToCelcius(k float64){
	c := k -273
	fmt.Println(k, "kelvin =", c,"celcius")
}

func kelvinToReamur(k float64){
	r := (k - 273) / 5 * 4
	fmt.Println(k, "kelvin =", r,"reamur")
}

func reamurToKelvin(r float64){
	k := (r / 4 * 5) + 273
	fmt.Println(r, "reamur =", k,"kelvin")
}
