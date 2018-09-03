package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type Wiki struct {
	Text  string `json:"text"`
	Title string `json:"title"`
}

func main() {
	wiki := []Wiki{}
	var check string
	re := regexp.MustCompile(`^.*(File|ファイル):([^\|]*)\|.*`)

	f, _ := os.Open("./jawiki-country.json")
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		b, err := r.ReadBytes('\n')
		var w Wiki
		if err == io.EOF {
			break
		}
		json.Unmarshal([]byte(b), &w)
		wiki = append(wiki, w)

		for _, name := range wiki {
			if name.Title == "イギリス" {
				check = name.Text
			}
		}
	}

	for _, s := range strings.Split(check, "\n") {
		ret := re.FindSubmatch([]byte(s))
		if ret != nil {
			fmt.Println(string(ret[2]))
		}
	}
}
