package main

import (
	"bufio"
	"errors"
	"fmt"
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

func generateBash(indexs []string, length int) string {
	tplText := `
#!/bin/sh
VIDEO_FILE=VID_20230913_210108.mp4
OUT=goal
INDEX=(__INDEX__)
LIMIT=__LIMIT__
mkdir -p ${OUT}
for i in ${INDEX[*]}
do
	ffmpeg -i ${VIDEO_FILE} -ss  00:02:38 -t 5 -an -vcodec copy ${OUT}/${VIDEO_FILE}_01.mp4
done
`

	return s

}

type Index struct {
	Index string
	Limit int
}

func main() {

	videoIndex := "/Users/rotaercw/Downloads/video.index"

	lines, err := ReadLines(videoIndex)
	if err != nil {
		fmt.Println(err.Error())
	}

	indx := Index{}

	for _, line := range lines {
		fmt.Println(line)
		s, e := stringToSeconds(strings.TrimSpace(line))
		if e != nil {
			fmt.Println(e.Error())
			continue
		}
		fmt.Println(intToTimeFormat(s - 3))
	}

	indx.Index = 

}
