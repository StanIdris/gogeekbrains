package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите числа через пробел, для примера: 5 4 1 3 6 2")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		x := scanner.Text()
		y := strings.Split(x, " ")

		var ints []int

		for _, s := range y {
			num, err := strconv.Atoi(s)
			if err == nil {
				ints = append(ints, num)
			} else if err != nil {
				fmt.Println(err)
				break
			}
		}

		sort.Ints(ints)
		fmt.Println(ints)
		break
	}
}
