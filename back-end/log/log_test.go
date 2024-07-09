package log

import (
	"testing"
	"os"
	"bufio"
	"regexp"
)

func getLines(fpath string) []string {
	var lines []string
	f, err := os.Open(fpath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	i := 0
	for s.Scan() {
		line := s.Text()
		lines = append(lines, line)
		i += 1
		if i == 2 { // in fail case, max would be 2
			break
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return lines
}

func TestLog(t *testing.T) {
	msg := "something"
	fpath := "test_log.log"
	if _, err := os.Stat(fpath); err == nil {
		os.Remove(fpath)
	}
	InitLogEngine(fpath, "file", WARN)
	L.Log(WARN, msg)
	L.Log(DEBUG, msg) //assert this one didn't go through

	logLines := getLines(fpath)
	if len(logLines) != 1 {
		t.Error("Log file should have 1 message")
	}
	line1 := logLines[0]
	// TODO find regex for "2006-01-02 15:04:05"
	messageFormatRegex := `^\[([A-Z]+)\]\[(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})\]: (.*)`
	re := regexp.MustCompile(messageFormatRegex)

	matches := re.FindStringSubmatch(line1)
	if matches == nil {
		t.Error("Line 1 of log does not match message pattern")
	}
}