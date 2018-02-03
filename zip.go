package main

import (
	"archive/zip"
	"io"
	"net/http"
	"strings"
)

func hander(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=sample.zip")

	// zip
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	// ファイル数だけ書き込み
	a, err := zipWriter.Create("a.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(a, strings.NewReader("1つめのファイルのテキスト"))

	b, err := zipWriter.Create("b.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(b, strings.NewReader("2つめのファイルのテキストです"))

}

func main() {
	http.HandleFunc("/", hander)
	http.ListenAndServe(":8080", nil)
}
