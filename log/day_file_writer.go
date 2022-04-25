package log

import (
	"github.com/pkg/errors"
	"io"
	"os"
	"path"
	"sync"
	"time"
)

type dayFileWriter struct {
	fileName       string
	lastYearDay    int
	outputFile     *os.File
	fileSwitchLock *sync.Mutex
}

func (w *dayFileWriter) Write(byteArray []byte) (n int, err error) {
	if nil == byteArray ||
		len(byteArray) <= 0 {
		return 0, nil
	}

	// 获取输出文件
	outputFile, err := w.getOutputFile()

	if nil != err {
		return 0, err
	}

	// 将日志打印到屏幕和写入文件
	_, _ = os.Stdout.Write(byteArray)
	_, _ = outputFile.Write(byteArray)

	return len(byteArray), nil
}

// 获取输出文件,
// 每天创建一个新的日志文件
func (w *dayFileWriter) getOutputFile() (io.Writer, error) {
	yearDay := time.Now().YearDay()

	if w.lastYearDay == yearDay &&
		nil != w.outputFile {
		return w.outputFile, nil
	}

	w.fileSwitchLock.Lock()
	defer w.fileSwitchLock.Unlock()

	if w.lastYearDay == yearDay &&
		nil != w.outputFile {
		return w.outputFile, nil
	}

	w.lastYearDay = yearDay

	// 构建日志目录
	err := os.MkdirAll(path.Dir(w.fileName), os.ModePerm)

	if nil != err {
		return nil, errors.Wrap(err, "创建目录失败")
	}

	// 定义新的日志文件
	newDailyFile := w.fileName + "." + time.Now().Format("20060102")

	outputFile, err := os.OpenFile(
		newDailyFile,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)

	if nil != err ||
		nil == outputFile {
		return nil, errors.Errorf("打开文件 %s 失败! err = %v", newDailyFile, err)
	}

	if nil != w.outputFile {
		// 关闭原来的文件
		_ = w.outputFile.Close()
	}

	w.outputFile = outputFile
	return outputFile, nil
}
