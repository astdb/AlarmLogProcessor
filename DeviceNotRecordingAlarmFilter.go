// reads in a text file containing alarm traces (not provided) and prints out the date and extension portions for each alarm.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	// read input filename from command line
	if len(os.Args) <= 1 {
		fmt.Println("Error: please enter the alarms file name on command line (e.g. > go run DeviceNotRecordingAlarmFilter.go alarms.txt)")
		return
	}

	// read alarms file
	inFileName := os.Args[1]
	file, err := os.Open(inFileName)
	logErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

    // detector for date and time e.g. "28/09/2016 9:00 AM"
	r1, _ := regexp.Compile("[0-9]{2}/[0-9]{2}/[0-9]{4} [0-9]*:[0-9]{2} [A|P]M")

    // detector for extension e.g. "34331"
	r2, _ := regexp.Compile("[0-9]{5}")

	// open output file
	outFileName := strings.Split(inFileName, ".")[0] + "-out.txt"
	f, err := os.Create(outFileName)
    logErr(err)
	defer f.Close()

	for scanner.Scan() {
		line := scanner.Text()
		date := strings.TrimSpace(r1.FindString(line))
		ext := strings.TrimSpace(r2.FindString(line))

		if len(date) > 0 && len(ext) > 0 {
			// write line to file 
			_, err := f.WriteString(r1.FindString(line) + "\t" + r2.FindString(line) + "\n")
			logErr(err)
		}
	}

	f.Sync()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func logErr(e error) {
    if e != nil {
		log.Fatal(e)
	}
}
