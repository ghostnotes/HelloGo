package controls

import (
    "fmt"
    "http"
)

func init() {
   http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    intValue := 5
    
    if intValue > 5 {
        // 判定結果がtrueのとき実行
        fmt.Fprintf(w, "5より大きい")
    } else {
        fmt.Fprintf(w, "5以下")
    }

}