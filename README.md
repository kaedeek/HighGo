# High Go

このライブラリは画像を、**ESPCN_x4.pb** で高画質にするライブラリです。

## Usage
```Go
package main

import (
    "fmt"
    "os/exec"
    "github.com/kaedeek/HighGo"
)

func main() {
    // モデルファイルのパス
    modelPath := "model/ESPCN_x4.pb"
    inputImage := "input.jpg"
    outputImage := "output.jpg"

    // HighGoライブラリを使用して超解像処理を実行
    err := HighGo.SuperResolve(inputImage, modelPath, outputImage)
    if err != nil {
        fmt.Println("Error in super resolution:", err)
        return
    }

    fmt.Println("High-resolution image saved as", outputImage)
}
```

## Authors
- Owner and Developer: [kaedeek](https://github.com/kaedeek)