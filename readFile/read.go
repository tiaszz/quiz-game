package readfile

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

func ReadCsvFile(file string) {
	r := csv.NewReader(strings.NewReader(file))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}
}
