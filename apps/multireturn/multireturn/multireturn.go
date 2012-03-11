package multireturn

import (
    "fmt"
    "http"
)

func init() {
    http.HandleFunc("/", handler)
}


func handler(w http.ResponseWriter, r *http.Request) {
    //　正常の割り算
    div, ok := divide_ex(16, 3)
    fmt.Fprintf(w, "正常:%f, %t\n", div, ok)

    // 割られる数が0
    div, ok = divide_ex(0, 3)
    fmt.Fprintf(w, "割られる数が0: %f, %t\n", div, ok)
    
    // 割る数が0
    div, ok = divide_ex(16, 0)
    fmt.Fprintf(w, "割る数が0: %f, %t\n", div, ok)
}

func divide_ex(a int, b int) (float32, bool) {
    var result float32 = 0.0
    var ok bool = false;

    if a > 0 && b > 0 {
        result = float32(a) / float32(b)
        ok = true
    } else if b > 0 {
        // 割られる数が0
        ok = true
    } else {
        // 割る数が0
        ok = false
    }

    return result, ok
}
