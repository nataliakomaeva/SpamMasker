package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите свою строку:")
	a := bufio.NewScanner(os.Stdin)
	a.Scan()
	fmt.Println(Spam(a.Text()))
}

func Spam(s string) string {
	b := []byte(s)

	for i := 0; i < len(b)-6; i++ {
		if string(b[i:i+7]) == "http://" { // маскируемм ссылку после http....
			for j := i + 7; j < len(b); j++ {
				if b[j] == ' ' {
					i = j // перескакиваем, чтобы не проверять символы внутри ссылки
					break
				}
				b[j] = '*'
			}
		}
	}
	return string(b)
}
