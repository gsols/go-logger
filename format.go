package logger

import "strconv"

var MaxParentDirs = 2

func countParentDirectories(filePath string, maxDirs int) int {
	parentsCounter := 0
	for i := len(filePath) - 1; i > 0; i-- {
		if filePath[i] == '/' {
			parentsCounter++
		}
		if parentsCounter == maxDirs+1 {
			return i + 1
		}
	}

	return 0
}

var CallerMarshalFunc = func(_ uintptr, file string, line int) string {
	file = file[countParentDirectories(file, MaxParentDirs):]
	return file + ":" + strconv.Itoa(line)
}
