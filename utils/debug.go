package utils

import (
	"bufio"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
)

func parseFileAndLine(line string) (string, int) {
	// 解析格式：/path/to/file.go:123 +0x456
	parts := strings.Split(line, " ")
	if len(parts) > 0 {
		fileLine := parts[0]
		if lastColon := strings.LastIndex(fileLine, ":"); lastColon != -1 {
			filePath := fileLine[:lastColon]
			lineNumStr := fileLine[lastColon+1:]
			if lineNum, err := strconv.Atoi(lineNumStr); err == nil {
				return filePath, lineNum
			}
		}
	}
	return "", 0
}

func readCodeLine(filePath string, lineNum int) string {
	file, err := os.Open(filePath)
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentLine := 1

	for scanner.Scan() {
		if currentLine == lineNum {
			return scanner.Text()
		}
		currentLine++
	}

	return ""
}

func readCodeLineContext(filePath string, lineNum int) string {
	file, err := os.Open(filePath)
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentLine := 1

	for scanner.Scan() {
		if currentLine == lineNum {
			code := strings.TrimSpace(scanner.Text())

			// 移除過長的程式碼，只保留前 100 個字符
			if len(code) > 100 {
				code = code[:100] + "..."
			}

			// 如果是空行或只有註釋，嘗試找上一行
			if code == "" || strings.HasPrefix(code, "//") {
				// 可以擴展邏輯找前後幾行
			}

			return code
		}
		currentLine++
	}

	return ""
}

func FormatStackOneLineWithCode() (string, int, string) {
	stack := debug.Stack()
	lines := strings.Split(string(stack), "\n")

	for i, line := range lines {
		if strings.Contains(line, ".go:") &&
			!strings.Contains(line, "/runtime/") &&
			!strings.Contains(line, "middleware") &&
			!strings.Contains(line, "github.com/") &&
			!strings.Contains(line, "utils/mq.go") &&
			!strings.Contains(line, "utils/debug.go") &&
			!strings.Contains(line, "/debug/") {

			if i+1 < len(lines) {
				fileLine := strings.TrimSpace(lines[i])
				funcLine := strings.TrimSpace(lines[i-1])

				// // 解析文件和行號
				if filePath, lineNum := parseFileAndLine(fileLine); filePath != "" {
					return filePath, lineNum, funcLine
					// 	code := readCodeLineContext(filePath, lineNum)

					// 	result := fmt.Sprintf("%s | %s",
					// 		strings.TrimSpace(line),
					// 		funcLine)

					// 	if code != "" {
					// 		result += fmt.Sprintf(" | CODE: %s", code)
					// 	}

					// 	return result
					// } else {
					// 	return fmt.Sprintf("%s | %s",
					// 		strings.TrimSpace(line),
					// 		funcLine)
				}
				return "", 0, funcLine
			}
		}
	}
	return "", 0, ""
}
