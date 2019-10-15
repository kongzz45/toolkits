package files

import (
	"bufio"
	"io/ioutil"
	"os"
	"path"
)

// IsExist  目录或者文件是否存在
func IsExist(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil || os.IsExist(err)
}

// UnderDirFiles 目录下的文件
func UnderDirFiles(dir string) ([]string, error) {
	if !IsExist(dir) {
		return nil, nil
	}
	fs, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	if len(fs) == 0 {
		return nil, nil
	}
	var ret []string
	for i := 0; i < len(fs); i++ {
		if !fs[i].IsDir() {
			ret = append(ret, fs[i].Name())
		}
	}
	return ret, nil
}

// FModifyTime 文件修改时间
func FModifyTime(fpath string) int64 {
	if !IsExist(fpath) {
		return 0
	}
	fs, err := os.Stat(fpath)
	if err != nil {
		return 0
	}
	return fs.ModTime().Unix()
}

// WriteToBytes 字节内容写入到文件
func WriteToBytes(fpath string, content []byte) (int, error) {
	os.MkdirAll(path.Dir(fpath), os.ModePerm)
	fs, err := os.Create(fpath)
	if err != nil {
		return 0, err
	}
	defer fs.Close()
	return fs.Write(content)
}

// WriteStrSliceToFile 将slice写入到文件(覆盖), sep == "\n"
func WriteStrSliceToFile(fpath string, content []string, sep string) error {
	if !IsExist(fpath) {
		os.Create(fpath)
	}
	fs, err := os.OpenFile(fpath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer fs.Close()
	writer := bufio.NewWriter(fs)
	for _, line := range content {
		if _, err := writer.WriteString(line + sep); err != nil {
			return err
		}
	}
	return writer.Flush()
}
