package common

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

// StringInput : input string from console
func StringInput(msg string) (string, error) {
	fmt.Printf("%s: ", msg)
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	t := input.Text()
	if err := input.Err(); err != nil {
		return "", err
	}
	return t, nil
}

// IntegerInput : input integer from console
func IntegerInput(msg string) (int, error) {
	fmt.Printf("%s: ", msg)
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	t := input.Text()
	if err := input.Err(); err != nil {
		return 0, err
	}
	v, err := strconv.Atoi(t)
	if err != nil {
		return 0, err
	}
	return v, nil
}

//ReadFileContent - Read file content
func ReadFileContent(path string) ([]byte, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Read File error: " + err.Error())
	}
	return dat, nil
}

//ClearInputScreen -
func ClearInputScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
