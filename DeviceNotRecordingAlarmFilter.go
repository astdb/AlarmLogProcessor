package main

import (
	"fmt"
	"os"
	"regexp"
    "bufio"
    "log"
    "strings"
)

func main() {
	// read input filename from command line
	if len(os.Args) <= 1 {
		fmt.Println("Error: please enter the alarms file name on command line (e.g. > DeviceNotRecordingAlarmFilter.go alarms.txt)")
		return
	}

	// read alarms file
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

    r1, _ := regexp.Compile("[0-9]{2}/[0-9]{2}/[0-9]{4} [0-9]*:[0-9]{2} [A|P]M")
    r2, _ := regexp.Compile("[0-9]{5}")

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
        line := scanner.Text()
        date := strings.TrimSpace(r1.FindString(line))
        ext := strings.TrimSpace(r2.FindString(line))

        if len(date) > 0 && len(ext) > 0 {
            fmt.Printf("%s\t%s\n", r1.FindString(line), r2.FindString(line))
        }
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
