package logger

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	MsgChanLength     = 10000       // 日志消息管道最大长度
	MaxMsgBufLength   = 100         // 日志消息缓冲区最大长度
	FlushTimeInterval = time.Second // 日志刷盘周期
)

// 日志对象
type logger struct {
	f           *os.File      // 日志文件指针
	bufMessages []*message    // 存储每一次需要同步的消息，相当于缓冲区，刷盘就会被清空
	mesChan     chan *message // 该管道接收每一条日志消息
}

// 日志消息
type message struct {
	content     string    // 日志内容
	currentTime time.Time // 日志写入时间
}

// 初始化一个logger对象
func New(logPath string) *logger {
	// 打开一个文件，这里的模式为创建或者追加
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}

	logger := &logger{
		f:           f,
		bufMessages: make([]*message, 0, MaxMsgBufLength),
		mesChan:     make(chan *message, MsgChanLength),
	}

	// 监听日志buf刷盘情况
	go logger.listenFlush()
	return logger
}

// 格式化一下日志，让日志好看一点儿
func (l *logger) formatMsg(mes *message) string {
	builder := &strings.Builder{}
	builder.WriteString(mes.currentTime.Format("2006-01-02 15:04:05.999999"))
	builder.WriteString(" ")
	builder.WriteString(mes.content)
	builder.WriteString("\n")
	return builder.String()
}

// 将日志入队
func (l *logger) Info(content string) {
	l.mesChan <- &message{
		content:     content,
		currentTime: time.Now(),
	}
}

// 批量将buf内容刷新到日志文件
func (l *logger) batchFlush() (err error) {
	builder := strings.Builder{}
	for _, mes := range l.bufMessages {
		// 将所有的buffer内容全部拼接起来
		builder.WriteString(l.formatMsg(mes))
	}

	content := builder.String()

	if content == "" {
		return
	}

	// 重置bufMessages
	l.bufMessages = make([]*message, 0, MaxMsgBufLength)

	// 写入日志文件
	_, err = l.f.WriteString(content)
	if err != nil {
		fmt.Println("写入日志文件失败,", err)
		return
	}

	fmt.Println("成功写入日志文件，", time.Now().String())

	return
}

// 监听刷盘情况
func (l *logger) listenFlush() {
	// 这里，我们加一个结束信号，来优雅退出
	c := make(chan os.Signal, 1)
	// 监听信号
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// 注册定时器
	ticker := time.NewTicker(FlushTimeInterval)

	for {
		select {
		case mes := <-l.mesChan:
			l.bufMessages = append(l.bufMessages, mes)
			if len(l.bufMessages) == MaxMsgBufLength {
				fmt.Println("缓冲区日志到达上限，将日志刷盘")
				l.batchFlush()
			}
		case <-ticker.C:
			fmt.Println("每隔1s，将日志刷盘")
			l.batchFlush()
		case <-c:
			fmt.Println("收到结束信号，将日志刷盘")
			l.batchFlush()
			return
		}
	}
}
