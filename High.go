package High

import (
	"fmt"
	"os/exec"
)

// SuperResolveは指定された画像を超解像化する関数です。
// inputImageは入力画像のパス、modelPathはモデルファイルのパス、outputImageは出力画像のパスです。
func SuperResolve(inputImage, modelPath, outputImage string) error {
	// Pythonスクリプトを実行するためのコマンドを作成
	cmd := exec.Command("python3", "body.py", inputImage, modelPath, outputImage)

	// コマンドの実行
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error executing Python script: %v", err)
	}
	return nil
}
