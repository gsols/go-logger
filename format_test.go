package logger_test

import (
	"testing"

	"github.com/gsols/go-logger"
)

func TestCallerMarshalFunc(t *testing.T) {
	tests := []struct {
		name string
		file string
		line int
		max  int
		want string
	}{
		{
			name: "Test with no parent",
			file: "file.txt",
			line: 10,
			max:  0,
			want: "file.txt:10",
		},
		{
			name: "Test with one parent directory",
			file: "dir/file.txt",
			line: 20,
			max:  1,
			want: "dir/file.txt:20",
		},
		{
			name: "Test with two parent directories",
			file: "dir/subdir/file.txt",
			line: 30,
			max:  2,
			want: "dir/subdir/file.txt:30",
		},
		{
			name: "Test with three parent directories and maxDirs set to 2",
			file: "dir/subdir/subsubdir/file.txt",
			line: 40,
			max:  2,
			want: "subdir/subsubdir/file.txt:40",
		},
		{
			name: "Test with four parent directories and maxDirs set to 3",
			file: "dir/subdir/subsubdir/subsubsubdir/file.txt",
			line: 40,
			max:  3,
			want: "subdir/subsubdir/subsubsubdir/file.txt:40",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger.MaxPathNumber = tt.max
			if got := logger.CallerMarshalFunc(0, tt.file, tt.line); got != tt.want {
				t.Errorf("CallerMarshalFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
