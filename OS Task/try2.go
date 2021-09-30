package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/robfig/cron"
)

func main() {

	// root := "C:\\go-workspace\\OS Task\\root\\"
	// chkfiles(root)

	timefunc()
	for {
		time.Sleep(time.Second)
	}

}

func checkfiles(root string) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".csv" {
			files = append(files, path)
			return nil
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
		fmt.Println(readCsvFile(file))
	}
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	return records
}

func timefunc() {
	root := "C:\\go-workspace\\OS Task\\root\\"
	c := cron.New()
	c.AddFunc("@every 5s", func() {
		log.Printf("5sec")
		checkfiles(root)
	})
	c.Start()
}
