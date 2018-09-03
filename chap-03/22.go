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
	/*
			JSON構造体定義
			書式は：変数名[tab]型名[tab]`json:"キー名"`
		　# JSON構造体の深さが可変のものは別途定義する必要あり。
	*/
	Text  string `json:"text"`
	Title string `json:"title"`
}

func main() {
	wiki := []Wiki{}
	var check string
	re := regexp.MustCompile(".*Category:(.*)]]")

	/*
		対象ファイルを読み込み、JSONをParseし、ハッシュへ格納
	*/
	f, _ := os.Open("./jawiki-country.json")
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		b, err := r.ReadBytes('\n')
		var w Wiki
		if err == io.EOF {
			break
		}
		/*
			func Unmarshal(data []byte, v interface{}) error
			第一引数にはJSONのbyte列、第二引数にはJSONをマッピングしたい値（今回の場合、構造体）のポインタ
		*/
		json.Unmarshal([]byte(b), &w)
		wiki = append(wiki, w)

		/*
			作成したハッシュのvalueをrangeで取り出し、
			イギリスにマッチするハッシュを検出。
		*/
		for _, name := range wiki {
			if name.Title == "イギリス" {
				check = name.Text
			}
		}
	}
	/*
		func (re *Regexp) FindSubmatch(b []byte) [][]byte
	*/
	for _, s := range strings.Split(check, "\n") {
		ret := re.FindSubmatch([]byte(s))
		if ret != nil {
			fmt.Println(string(ret[1]))
		}
	}
}
