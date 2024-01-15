package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("asin_codes.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	asin_codes := []string{}

	for scanner.Scan() {
		asin_codes = append(asin_codes, strings.TrimSpace(scanner.Text()))
	}

	url := "https://www.amazon.com/gp/aws/cart/add.html?AssociateTag=Associate+Tag&"
	for i, asin_code := range asin_codes {
		url += fmt.Sprintf("ASIN.%d=%s&Quantity.%d=1&", i+1, asin_code, i+1)
	}

	fmt.Println(strings.TrimSuffix(url, "&"))
}
