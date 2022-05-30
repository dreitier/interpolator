package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("No input file given")
	}


	inputFileName := os.Args[1]
	var outputFileName string

	if len(os.Args) > 2 {
		outputFileName = os.Args[2]
		log.Printf("Replacing variable references in file %s, output file is %s", inputFileName, outputFileName)
	} else {
		outputFileName = inputFileName
		log.Printf("Replacing variable references in file %s in place", inputFileName)
	}

	info, err := os.Stat(inputFileName)
	origPerms := info.Mode().Perm()

	newContent := readAndReplace(inputFileName)

	err = ioutil.WriteFile(outputFileName, []byte(newContent), origPerms)
	if err != nil {
		log.Fatalln(err)
	}

}

func readAndReplace(fileName string) string {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalln("Failed to open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	var output []string
	for scanner.Scan() {
		line := scanner.Text()

		r, _ := regexp.Compile(".*(\\${\\w*}).*")
		result := r.FindStringSubmatch(line)

		if len(result) == 2 {
			varName := result[1][2 : len(result[1])-1]
			osVar, found := os.LookupEnv(varName)
			if found {
				line = strings.Replace(result[0], result[1], osVar, -1)
				log.Printf("Replacing placeholder '%s' with environment variable\n", result[1])
			} else {
				log.Printf("Environment variable '%s' was not set or empty\n", result[1])
			}

		}
		output = append(output, line)
	}
	return strings.Join(output, "\n")
}
