package main

import "fmt"

//Declare Stuct here
type currency struct {
	currency     string
	exchangeRate float32
}

//Declare Map here
var currencyExchange map[string]currency

func main() {
	//Insert Main Code here

	// activity 1
	currencyExchange = make(map[string]currency)
	currencyExchange["USD"] = currency{"US Dollar", 1.1318}
	currencyExchange["JPY"] = currency{"Japanese yen", 121.05}
	currencyExchange["GBP"] = currency{"Pound sterling", 0.9063}
	currencyExchange["CNY"] = currency{"Chinese yuan renminbi", 7.9944}
	currencyExchange["SGD"] = currency{"Singapore dollar", 1.5743}
	currencyExchange["MYR"] = currency{"Malaysia ringgit", 4.839}

	fmt.Println(currencyExchange)
	fmt.Println(currencyExchange["SGD"].currency, currencyExchange["SGD"].exchangeRate)

	// activity 2
	var userCurrency string
	var userAmount float32
	var convCurrency string

	fmt.Println(("Please choose the currency you use: "))
	fmt.Scanln(&userCurrency)
	fmt.Println(("Please choose the amount: "))
	fmt.Scan(&userAmount)
	fmt.Println(("Please choose the currency you want to exchange to: "))
	fmt.Scanln(&convCurrency)

	result := (userAmount / currencyExchange[userCurrency].exchangeRate) * currencyExchange[convCurrency].exchangeRate
	fmt.Println(result)

}
