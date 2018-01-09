package log

import (
	"fmt"
	"os"
	"path/filepath"
)

func __get_rel_go_path_from_abs_path__(filePath string) (string, error) {
	return __get_rel_path_from_abs_path__(filePath, fmt.Sprintf("%s/%s", os.Getenv("GOPATH"), "src"))
}
func __get_rel_path_from_abs_path__(filePath, RefPath string) (string, error) {
	return filepath.Rel(RefPath, filePath)
}