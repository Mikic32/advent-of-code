package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func aoc2() {
	content, err := os.ReadFile("elfs.txt")
	if err != nil {
		fmt.Println("error loading file:", err)
		return
	}
	contentStr := string(content)
	splitedContent := strings.Split(contentStr, "\n\n")
	summedContent := make([]int, len(splitedContent))

	for _, cals := range splitedContent {
		splitedContent = strings.Split(cals, "\n")
		sumOfOneElf := 0
		for _, item := range splitedContent {
			if value, err := strconv.Atoi(item); err == nil {
				sumOfOneElf += value
			}
		}
		summedContent = append(summedContent, sumOfOneElf)
	}
	sort.Slice(summedContent, func(i, j int) bool {
		return summedContent[i] > summedContent[j]
	})
	sumOfTop3 := summedContent[0] + summedContent[1] + summedContent[2]
	fmt.Println(sumOfTop3)
}
