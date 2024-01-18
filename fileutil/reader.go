package fileutil

import (
	"bufio"
	"os"
)

type FileReader struct {
	file   *os.File
	reader *bufio.Reader
}

func NewReader(path string) FileReader {
	f, err := os.Open(path)
	if err != nil {
		panic("Unable to open file")
	}
	reader := bufio.NewReader(f)
	return FileReader{
		file:   f,
		reader: reader,
	}
}

func (r *FileReader) ReadLine() (string, error) {
	return r.reader.ReadString('\n')
}

func (r *FileReader) Close() {
	r.file.Close()
}
