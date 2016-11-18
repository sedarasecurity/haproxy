package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	b, _ := ioutil.ReadFile("events.txt")
	r := bufio.NewReader(bytes.NewReader(b))
	var lines []string

	for {
		bl, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		if len(bl) == 0 {
			continue
		}

		lines = append(lines, string(bl))
	}

	for _, x := range lines {
		fmt.Printf("INSERT INTO `plugin_sid` (`plugin_id`,`sid`,`category_id`,`subcategory_id`,`reliability`,`priority`,`name`) VALUES (81875, %s, 13, 109, 1, 1, '%s');\n", x[:3], x[4:])
	}
}
