package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "birthDay_01.txt"
	newName, err := match(filename)
	if err != nil {
		fmt.Println("no match")
		os.Exit(1)
	}
	fmt.Println(newName)
}
func match(fileName string) (string, error) {
	pieces := strings.Split(fileName, ".")
	ext := pieces[len(pieces)-1]
	tmp := strings.Join(pieces[:len(pieces)-1], ".")
	pieces = strings.Split(tmp, "_")
	no, err := strconv.Atoi(pieces[len(pieces)-1])

	if err != nil {
		return "", fmt.Errorf("Not found")
	}
	name := strings.Join(pieces[:len(pieces)-1], "_")
	return fmt.Sprintf("%s-%d.%s", strings.Title(name), no, ext), nil
}
