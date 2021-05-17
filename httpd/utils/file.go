package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

type FilesModTime []string

func (files FilesModTime) Len() int {
	return len(files)
}

func (files FilesModTime) Swap(i, j int) {
	files[i], files[j] = files[j], files[i]
}

func (files FilesModTime) Less(i, j int) bool {
	fileI, err := os.Stat(files[i])
	if err != nil {
		fmt.Println(files[i], err)
	}
	fileJ, err := os.Stat(files[j])
	if err != nil {
		fmt.Println(files[j], err)
	}
	return fileI.ModTime().Unix() >= fileJ.ModTime().Unix()
}

func Bytes2File(data []byte, fileName string) {
	tmpFile := fileName + ".tmp"
	createFile(data, tmpFile)
	bakFile := fileName + ".bak"
	oldData, err := ioutil.ReadFile(fileName)
	if err == nil {
		// 先删除备份文件
		os.Remove(bakFile)
		//再创建新的备份文件
		createFile(oldData, bakFile)
	}
	// 删除原文件
	os.Remove(fileName)
	// 重命临时文件
	os.Rename(tmpFile, fileName)
}

func createFile(data []byte, fileName string) {
	err := EnsureDir(fileName)
	if err != nil {
		fmt.Println("EnsureDir error=%v ", err.Error())
	}
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("create file error ", err.Error())
	}
	file.Write(data)
	file.Sync()
	defer file.Close()
}

// File2String 读取文件内容
func File2String(filePath string) (data string, err error) {
	bf, err := file2String(filePath)
	if err != nil {
		return "", err
	}

	return string(bf), nil
}

func file2String(filePath string) (bf []byte, err error) {
	bf, err = ioutil.ReadFile(filePath)
	if err != nil {
		return []byte{}, err
	}

	return bf, nil
}

func File2Bytes(filePath string) (bf []byte, err error) {
	bf, err = ioutil.ReadFile(filePath)
	if err != nil {
		return []byte{}, err
	}

	return bf, nil
}

// CreateDir 创建文件夹
func CreateDir(dir string) (bool, error) {
	if err := os.MkdirAll(dir, os.FileMode(os.O_CREATE)); err != nil {
		return false, err
	}

	return true, nil
}

// CreateFile 创建文件
func CreateFile(fileFullName string) (bool, error) {
	parentDir := filepath.Dir(fileFullName)
	_, err := CreateDir(parentDir)
	if err != nil {
		return false, err
	}

	_, err = os.Create(fileFullName)
	if err != nil {
		return false, err
	}

	return true, nil
}

// ListFilesOrDir 返回当前目录下所有的文件或者目录
// Params: path 路径
// Return: listType 从ALL|DIR|FILE选取，分别代表返回所有，目录，文件
// Since: 2017/8/8
func ListFilesOrDir(path string, listType string) ([]string, error) {
	var pathSlice []string
	err := filepath.Walk(path, func(path2 string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			if listType == "DIR" || listType == "ALL" {
				pathSlice = append(pathSlice, path2)
			}
		} else if listType == "FILE" || listType == "ALL" {
			pathSlice = append(pathSlice, path2)
		}
		return nil
	})
	return pathSlice, err
}

func ListSubFilesOrDir(path string, listType string) ([]string, error) {
	var pathSlice []string
	err := filepath.Walk(path, func(path2 string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if path2 == path {
			return nil
		}
		if f.IsDir() {
			if listType == "DIR" || listType == "ALL" {
				pathSlice = append(pathSlice, path2)
			}
		} else if listType == "FILE" || listType == "ALL" {
			pathSlice = append(pathSlice, path2)
		}
		return nil
	})
	return pathSlice, err
}

func ListDirectFilesOrDir(path string, listType string) ([]string, error) {
	var pathSlice []string
	files, err := filepath.Glob(filepath.Join(path, "*"))
	if listType == "ALL" {
		return files, err
	}
	if listType == "DIR" {
		for _, file := range files {
			f, err := os.Stat(file)
			if err != nil {
				continue
			}
			if f.IsDir() {
				pathSlice = append(pathSlice, file)
			}
		}
	}
	if listType == "FILE" {
		for _, file := range files {
			f, err := os.Stat(file)
			if err != nil {
				continue
			}
			if !f.IsDir() {
				pathSlice = append(pathSlice, file)
			}
		}
	}
	return pathSlice, err
}

func FileSize(file *os.File) (int64, error) {
	f, err := file.Stat()
	if err != nil {
		return -1, err
	}
	return f.Size(), nil

}

// get filepath base name
func Basename(fp string) string {
	return path.Base(fp)
}

// get filepath dir name
func FileDir(fp string) string {
	return path.Dir(fp)
}

func InsureDir(fp string) error {
	if IsExist(fp) {
		return nil
	}
	return os.MkdirAll(fp, os.ModePerm)
}

// EnsureDir dir if not exist
// 如果path为"./tmp/test/test.txt"则只会创建./tmp/test/目录
// Since: 2017/8/8
func EnsureDir(fp string) error {
	dir := FileDir(fp)
	return os.MkdirAll(dir, os.ModePerm)
}

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func IsExist(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil || os.IsExist(err)
}

// Offset2FileName 格式化位20个字符长度的string
// Since: 2017/8/9
func Offset2FileName(offset int64) string {
	fileName := strconv.FormatInt(offset, 10)
	var byteList = make([]byte, 20)
	i := 0
	for i < 20 {
		byteList[i] = 0 + '0'
		i++
	}
	byteList = append(byteList, fileName...)
	index := len(byteList) - 20

	return string(byteList[index : index+20])
}

func ReadFileBytes(filePth string, byteSize int64) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buf := make([]byte, byteSize)
	bfRd := bufio.NewReader(f)
	_, err = bfRd.Read(buf)
	return buf, err
}

func GetSize(path string) (int64, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	fileSize := fileInfo.Size()
	return fileSize, nil
}

func GetLineOffsetDesc(filename string, n int) (offset int64, err error) {
	var defaultBufSize int64 = 1
	f, e := os.Stat(filename)
	if e == nil {
		size := f.Size()
		var fi *os.File
		fi, err = os.Open(filename)
		if err == nil {
			b := make([]byte, defaultBufSize)
			sz := int64(defaultBufSize)
			nn := n
			bTail := bytes.NewBuffer([]byte{})
			istart := size
			flag := true
			for flag {
				if istart < defaultBufSize {
					sz = istart
					istart = 0
				} else {
					istart -= sz
				}
				offset, err = fi.Seek(istart, io.SeekStart)
				if err == nil {
					mm, e := fi.Read(b)
					if e == nil && mm > 0 {
						j := mm
						for i := mm - 1; i >= 0; i-- {
							if b[i] == '\n' {
								bLine := bytes.NewBuffer([]byte{})
								bLine.Write(b[i+1 : j])
								j = i
								if bTail.Len() > 0 {
									bLine.Write(bTail.Bytes())
									bTail.Reset()

								}
								if (nn == n && bLine.Len() > 0) || nn < n {
									nn--

								}
								if nn == 0 {
									flag = false
									break

								}

							}

						}
						if flag && j > 0 {
							if istart == 0 {
								bLine := bytes.NewBuffer([]byte{})
								bLine.Write(b[:j])
								if bTail.Len() > 0 {
									bLine.Write(bTail.Bytes())
									bTail.Reset()

								}
								flag = false
							} else {
								bb := make([]byte, bTail.Len())
								copy(bb, bTail.Bytes())
								bTail.Reset()
								bTail.Write(b[:j])
								bTail.Write(bb)
							}
						}
					} else {
						flag = false
						break
					}
				} else {
					flag = false
					break
				}
			}
		}
		defer fi.Close()
	}
	return
}
