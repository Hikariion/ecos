package common

import (
	"ecos/utils/logger"
	"errors"
	"os"
)

// InitPath create path as an empty dir
func InitPath(path string) error {
	s, err := os.Stat(path) // path 是否存在
	if err != nil {
		if os.IsExist(err) { //
			return err // path 存在，且不是目录
		}
		logger.Infof("Path: %v not exist, create it", path)
		err = os.MkdirAll(path, 0777) // 目录不存在，创建空目录
		if err != nil {
			return err
		}
		return nil
	}
	if !s.IsDir() {
		return errors.New("path exist and not a dir")
	}
	return nil
}
