package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io"
	"os"
	"strconv"
	"strings"
)

// ReadLinesV2 reads all lines of the file.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	reader := bufio.NewReader(file)
	for {
		// ReadString reads until the first occurrence of delim in the input,
		// returning a string containing the data up to and including the delimiter.
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			lines = append(lines, line)
			break
		}
		if err != nil {
			return lines, err
		}
		lines = append(lines, line[:len(line)-1])
	}
	return lines, nil
}

func stringToSeconds(str string) (int, error) {
	if len(str) == 0 {
		return 0, errors.New("invalid 0")
	}
	hms := strings.Split(str, ":")
	n := len(hms)

	if n == 1 {
		return strconv.Atoi(hms[0])
	} else if n == 2 {
		m, e0 := strconv.Atoi(hms[0])
		s, e1 := strconv.Atoi(hms[1])
		if e0 != nil || e1 != nil {
			return 0, errors.New("invalid 2")
		}
		return m*60 + s, nil
	} else if n == 3 {
		h, e0 := strconv.Atoi(hms[0])
		m, e1 := strconv.Atoi(hms[1])
		s, e2 := strconv.Atoi(hms[2])
		if e0 != nil || e1 != nil || e2 != nil {
			return 0, errors.New("invalid 3")
		}
		return h*3600 + m*60 + s, nil
	}

	return 0, errors.New("invalid--")
}

func intToTimeFormat(i int) string {
	h := i / 3600
	m := (i - h*3600) / 60
	s := i - h*3600 - m*60
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

type Index struct {
	Index  []string
	Limit  int
	Offset int
}

func NewIndex(limit, offset int) *Index {
	return &Index{
		Index:  make([]string, 0),
		Limit:  limit,
		Offset: offset,
	}
}

func (i *Index) AppendTimeStamp(ts string) {
	i.Index = append(i.Index, ts)
}

func (i *Index) AppendSec(s int) {
	if s < 3 {
		return
	}
	i.Index = append(i.Index, intToTimeFormat(s-3))
}

func (i *Index) SetLimit(l int) {
	i.Limit = l
}

func (i *Index) Template() {
	tplText := `#!/bin/sh
VIDEO_FILE=__FILE__
OUT=goal
LIMIT={{.Limit}}
mkdir -p ${OUT}
{{- range $i, $e := .Index}}
ffmpeg -i ${VIDEO_FILE} -ss  {{$e}} -t ${LIMIT} -an -vcodec copy ${OUT}/${VIDEO_FILE}_{{$i}}.mp4
{{- end }}
`
	tpl, err := template.New("bash").Parse(tplText)
	if err != nil {
		fmt.Printf("failed parse tpltext,err:%s\n", err.Error())
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, i)
	if err != nil {
		fmt.Printf("failed execute tpltext,err:%s\n", err.Error())
	}

	filePath := "export.sh"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	buf.WriteTo(file)

}

func main() {

	var limit int
	var offset int
	var videoIndex string

	// Example command: go run echo.go --port 9000 --multicore=true
	flag.IntVar(&limit, "limit", 5, "--limit 5")
	flag.IntVar(&offset, "offset", 5, "--offset 5")
	flag.StringVar(&videoIndex, "index", "video.index", "--index video.index")
	flag.Parse()

	lines, err := ReadLines(videoIndex)
	if err != nil {
		fmt.Println(err.Error())
	}

	indx := NewIndex(limit, offset)

	for _, line := range lines {
		// fmt.Println(line)
		s, e := stringToSeconds(strings.TrimSpace(line))
		if e != nil {
			fmt.Println(e.Error())
			continue
		}
		fmt.Println(s)
		indx.AppendSec(s)
		// fmt.Println(intToTimeFormat(s - 3))
	}

	indx.Template()

}
