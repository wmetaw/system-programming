package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go func() {
			fmt.Printf("Accept %v\n", conn.RemoteAddr())

			// Accept後のソケットで何度も応答を返すためにループ
			for {
				// タイムアウトを設定
				conn.SetReadDeadline(time.Now().Add(4 * time.Second))

				// リクエストを読み込む
				request, err := http.ReadRequest(bufio.NewReader(conn))

				// タイムアウト or ソケットクローズ時は終了。それ以外はエラー
				if err != nil {
					neterr, ok := err.(net.Error) // 型アサーションを使用してエラーがnet.Error型であるかどうかを確認
					if ok && neterr.Timeout() {
						fmt.Println("timeout")
						break
					} else if err == io.EOF {
						break
					}
					panic(err)
				}

				// リクエストを表示
				dump, err := httputil.DumpRequest(request, true)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(dump))

				// レスポンスの書き込み(HTTP/1.1かつContentLengthの設定が必要)
				content := "Done"
				response := http.Response{
					StatusCode:    200,
					ProtoMajor:    1,
					ProtoMinor:    1,
					ContentLength: int64(len(content)),
					Body:          ioutil.NopCloser(strings.NewReader("Done")),
				}
				response.Write(conn)
			}
			conn.Close()
		}()
	}
}
