package logger

import "strconv"

var MaxPathNumber = 2

func countParentDirectories(filePath string, maxPath int) int {
	parentsCounter := 0
	for i := len(filePath) - 1; i > 0; i-- {
		if filePath[i] == '/' {
			parentsCounter++
		}
		if parentsCounter == maxPath {
			return i + 1
		}
	}

	return 0
}

var CallerMarshalFunc = func(_ uintptr, file string, line int) string {
	file = file[countParentDirectories(file, MaxPathNumber):]
	return file + ":" + strconv.Itoa(line)
}
