package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Confirm(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s [y/n]: ", s)

	response, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	response = strings.ToLower(strings.TrimSpace(response))

	if response == "y" || response == "yes" {
		return true
	}
	return false
}
