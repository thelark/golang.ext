package file

import (
	"bytes"
	"encoding/json"
	"image"
	"image/png"
	"io"
	"os"
	"strings"
	"thelark.cn/golang.ext/wrong"
)

/**
 * TODO: 文件操作扩展
 */

// TODO: 判断路径是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	wrong.Println(err)
	return false
}

// TODO: 创建目录
func MkDIR(path string) {
	if IsExist(path) {
		return
	}
	err := os.MkdirAll(path, 0777)
	wrong.Println(err)
}

// TODO: 将 []byte 保存到文件 => 覆盖性写入 ✔
func SaveFile(path string, data []byte) {
	if len(data) == 0 {
		return
	}
	f, err := os.Create(path) //创建文件
	if err != nil {
		wrong.Println(err)
		return
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		wrong.Println(err)
		return
	}
}

// TODO: 将 string 保存到文件 => 覆盖性写入 ✔
func SaveFileByString(path string, data string) {
	if len(data) == 0 {
		return
	}
	f, err := os.Create(path) //创建文件
	if err != nil {
		wrong.Println(err)
		return
	}
	defer f.Close()
	_, err = io.WriteString(f, data) //写入文件(字符串)
	if err != nil {
		wrong.Println(err)
		return
	}
}

// TODO: 将 interface{} 转换 Json 后保存到文件 => 覆盖性写入 ✔
func SaveFileByInterface(path string, data interface{}) {
	if data == nil {
		return
	}
	resJsonBytes, err := json.Marshal(data)
	if err != nil {
		wrong.Println(err)
	}
	SaveFile(path, resJsonBytes)
}

// TODO: 将文件删除
func DeleteFile(path string) {
	fI, err := os.Stat(path)
	if err != nil {
		wrong.Println(err)
		return
	}
	if fI.IsDir() {
		return
	}
	wrong.Println(os.Remove(path))
}

// TODO: 将 []byte 保存为图片
func SaveImage(path string, data []byte) {
	reader := bytes.NewReader(data)
	img, _, err := image.Decode(reader)
	if err != nil {
		wrong.Println(err)
		return
	}
	f, err := os.Create(path)
	if err != nil {
		wrong.Println(err)
		return
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		wrong.Println(err)
		return
	}
}

// TODO: 将 string 保存为图片
func SaveImageByString(path, data string) {
	readaer := strings.NewReader(data)
	img, _, err := image.Decode(readaer)
	if err != nil {
		wrong.Println(err)
		return
	}
	file, err := os.Create(path)
	if err != nil {
		wrong.Println(err)
		return
	}
	defer file.Close()
	err = png.Encode(file, img)
	if err != nil {
		wrong.Println(err)
		return
	}
}
