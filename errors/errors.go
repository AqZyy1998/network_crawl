package errors

import (
	"bytes"
	"fmt"
)

// CrawlerError 爬虫错误接口
type CrawlerError interface {
	// Type 错误类型
	Type() ErrorType
	// Error 获取错误信息
	Error() string
}

// ErrorType error类型
type ErrorType string

// MyCrawlerError 爬虫错误的实例类型
type MyCrawlerError struct {
	// errType error类型
	errType ErrorType
	// errMsg error简要信息
	errMsg string
	// fullErrMsg error完整信息
	fullErrMsg string
}

const (
	// ERROR_TYPE_DOWNLOADER 下载器error
	ERROR_TYPE_DOWNLOADER ErrorType = "downloader error"
	// ERROR_TYPE_ANALYZER 分析器error
	ERROR_TYPE_ANALYZER ErrorType = "analyzer error"
	// ERROR_TYPE_PIPELINE 条目管理管道error
	ERROR_TYPE_PIPELINE ErrorType = "pipeline error"
	// ERROR_TYPE_SCHEDULER 调度器error
	ERROR_TYPE_SCHEDULER ErrorType = "scheduler error"
)

// NewCrawlerError 实例化error
func NewCrawlerError(errorType ErrorType, errMsg string) CrawlerError {
	return &MyCrawlerError{
		errType: errorType,
		errMsg: errMsg,
	}
}

// Type 获取错误类型
func (ce *MyCrawlerError) Type() ErrorType {
	return ce.errType
}

// Error 获取错误信息
func (ce *MyCrawlerError) Error() string {
	if ce.fullErrMsg == "" {
		ce.GenFullErrMsg()
	}
	return ce.fullErrMsg
}

// GenFullErrMsg 生成error完整信息
func (ce *MyCrawlerError) GenFullErrMsg() {
	var buffer bytes.Buffer
	buffer.WriteString("crawler error: ")
	if len(ce.errType) > 0 {
		buffer.WriteString(string(ce.errType))
		buffer.WriteString(": ")
	}
	buffer.WriteString(ce.errMsg)
	ce.fullErrMsg = fmt.Sprintf("%s", buffer.String())
	return
}