package golib

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func LoadENV(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if len(scanner.Text()) > 0 { // this check eliminates empty line consideration
			// what if line start with a space?
			if scanner.Text()[0] == 32 {
				return errors.New("line can not start with space")
			}

			// line start with #?
			if scanner.Text()[0] == 35 { // 35 means # (comment)
				continue
			}

			items := strings.Split(scanner.Text(), "=")
			for index := range items {
				if index == 1 {
					err := os.Setenv(items[0], items[1])
					if err != nil {
						return err
					}
				}

				if index > 1 { // env value has =
					value := os.Getenv(items[0])
					err := os.Setenv(items[0], value+"="+items[index])
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}
