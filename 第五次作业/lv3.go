package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"your-module-path/pool"
)

type FileSearchTask struct {
	filePath string
	keyword  string
}

func (t *FileSearchTask) Run() {

	file, err := os.Open(t.filePath)
	if err != nil {
		fmt.Printf("打开文件失败 %s: %v\n", t.filePath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if strings.Contains(line, t.keyword) {
			fmt.Printf("%s:%d: %s\n", t.filePath, lineNum, line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("读取文件失败 %s: %v\n", t.filePath, err)
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("用法: ./catch [需要检索的目录] [需要检索的关键词]")
		os.Exit(1)
	}
	rootDir := os.Args[1]
	keyword := os.Args[2]

	workerPool := pool.NewPool(5, 1000)
	defer workerPool.Wait()

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			task := &FileSearchTask{
				filePath: path,
				keyword:  keyword,
			}
			workerPool.Submit(task)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("遍历目录失败 %s: %v\n", rootDir, err)
		os.Exit(1)
	}
}
