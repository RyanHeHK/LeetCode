package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}

func GetSalt() string {
	salt := ""
	env := "live"
	if strings.ToUpper(env) == "LIVE" {
		salt = "u9Tp0BGBSWmu4XZmuWKzGR7ZYXUzLXF2hMvttnvKQIdm5SYidYJeyDBirx1t"
	} else if strings.ToUpper(env) == "STAGING" {
		salt = "f4aDm2MgKI9LPuhCqYEQbStC5wVQB29ImUgNUlD1heA7OJoB5Y0yN4m3SgsU"
	} else if strings.ToUpper(env) == "UAT" {
		salt = "e2Btqpt7EW0YBZa0zMiS0BtYmfpaCgK3Vvf7blw2HeFm7sWWIaTfySvD8qNh"
	} else {
		salt = "TsTGTpqQqkWAo5y74aIafr9JX6od7FT6gnm3iaEqQCSKS4VMAuADAYc7j9Gu"
	}
	return salt
}

func myAppend(a []int) {
	a = append(a, 100)
}

type baseRequest struct {
	Timestamp int64  `json:"ts"`
	Token     string `json:"token"`
}

type Req struct {
	baseRequest
	AddDomesticPayeeRequest
}

type AddDomesticPayeeRequest struct {
	MerchantID        int64  `json:"merchant_id"`
	PayeeType         string `json:"payee_type"`
	BankAccountName   string `json:"bank_account_name"`
	BankAccountNo     string `json:"bank_account_no"`
	BankNameID        uint32 `json:"bank_name_id"`
	BankName          string `json:"bank_name"`
	NUCCCode          string `json:"nucc_code"`
	EPCCBankCode      string `json:"epcc_bank_code"`
	CityCode          string `json:"city_code"`
	SubBankName       string `json:"sub_bank_name"`
	BankRegion        string `json:"bank_region"`
	CompanyCertFileID string `json:"company_cert_file_id"`
}

//func main() {
//	ch := make(chan string)
//
//	go func() { // calculate goroutine
//		fmt.Println("calculate goroutine starts calculating")
//		time.Sleep(time.Second) // Heavy calculation
//		fmt.Println("calculate goroutine ends calculating")
//
//		ch <- "FINISH" // goroutine 執行會在此被迫等待
//
//		fmt.Println("calculate goroutine finished")
//	}()
//
//	time.Sleep(2 * time.Second) // 使 main 比 goroutine 慢
//	fmt.Println(<-ch)
//	time.Sleep(time.Second)
//	fmt.Println("main goroutine finished")
//}

//func main() {
//	ch := make(chan string)
//
//	go func() {
//		fmt.Println("calculate goroutine starts calculating")
//		time.Sleep(time.Second) // Heavy calculation
//		fmt.Println("calculate goroutine ends calculating")
//
//		ch <- "FINISH"
//
//		fmt.Println("calculate goroutine finished")
//	}()
//
//	fmt.Println("main goroutine is waiting for channel to receive value")
//	fmt.Println(<-ch) // goroutine 執行會在此被迫等待
//	fmt.Println("main goroutine finished")
//}
