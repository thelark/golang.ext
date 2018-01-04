package image

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"image/gif"
	"io"
	"log"
	"os"
	"strings"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

/**
 * Save
 * @param path 路径
 * @param data 数据
 * @param suffix 后缀名
 * @param _type 数据类型
 */
func Save(path string, data interface{}, suffix string, _type ...string) error {
	var err error
	var reader io.Reader
	var img image.Image

	if len(_type) == 0 {
		switch data.(type) {
		case string:
			reader = strings.NewReader(data.(string))
		case []byte:
			reader = bytes.NewReader(data.([]byte))
		}
		img, _, err = image.Decode(reader)
		if err != nil {
			return err
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if len(_type) == 1 {
		switch _type[0] {
		case "base64":
			dist, _ := base64.StdEncoding.DecodeString(fmt.Sprintf("%s", data))
			file.Write(dist)
			return nil
		}
	}
	funcMap[suffix](file, img)
	return nil
}

var funcMap = map[string]func(writer io.Writer, img image.Image){
	".png": func(writer io.Writer, img image.Image) {
		if err := png.Encode(writer, img); err != nil {
			log.Println(err)
		}
	},
	".jpg": func(writer io.Writer, img image.Image) {
		if err := jpeg.Encode(writer, img, &jpeg.Options{100}); err != nil {
			log.Println(err)
		}
	},
	".gif": func(writer io.Writer, img image.Image) {
		if err := gif.Encode(writer, img, &gif.Options{NumColors: 256}); err != nil {
			log.Println(err)
		}
	},
}
