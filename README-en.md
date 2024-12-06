# High Go
This library enhances image resolution using ESPCN_x4.pb, allowing for high-quality outputs.

## Usage
```Go
package main

import (
    "fmt"
    "github.com/kaedeek/HighGo"
)

func main() {
    // Path to the model file
    modelPath := "model/ESPCN_x4.pb"
    inputImage := "input.jpg"
    outputImage := "output.jpg"

    // Perform super-resolution using the HighGo library
    err := HighGo.SuperResolve(inputImage, modelPath, outputImage)
    if err != nil {
        fmt.Println("Error in super resolution:", err)
        return
    }

    fmt.Println("High-resolution image saved as", outputImage)
}
```


## Authors
Owner and Developer: [kaedeek](https://github.com/kaedeek)