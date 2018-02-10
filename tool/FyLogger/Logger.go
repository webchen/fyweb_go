package FyLogger

import (
	"goweb/config"
	"goweb/core"
	"goweb/tool"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/huandu/goroutine"
	//"github.com/huandu/goroutine"
)

/*
type fyFileLog struct {
	fileName string
	logPath  string
	level    uint
	log      *log.Logger
}

func newFyFileLog(fileName string, logPath string, level uint) (f *fyFileLog) {
	f = new(fyFileLog)
	f.fileName = fileName
	f.logPath = logPath
	f.level = level
	f.log = new(log.Logger)
	// 这里传递指针，主要是避免如果这里传递空字符串的话，logDir函数会处理
	fullFile := logDir(&f.fileName, &f.logPath)
	file, err := os.OpenFile(fullFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatalln("创建日志失败：", err)
	}
	f.log.SetFlags(log.Ldate | log.Lmicroseconds)
	f.log.SetOutput(file)
	return
}
*/

func access(fileName string, filePath string) (l *log.Logger) {
	fullFile := logDir(&filePath, &fileName, true)
	file, err := os.OpenFile(fullFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatalln("创建日志失败：", err)
	}
	l = new(log.Logger)
	l.SetFlags(log.Ldate | log.Lmicroseconds)

	ticker := time.NewTicker(time.Second * config.ACCESSLOGCHECKMINUTE)
	go func() {
		for _ = range ticker.C {
			info, _ := os.Stat(fullFile)
			// 日志文件大于100M，或隔日了，则日志自动分文件
			if info.Size() >= config.ACCESSLOGCHECKSIZE || info.ModTime().Format("20060102") != time.Now().Format("20060102") {
				mu := new(sync.Mutex)
				mu.Lock()
				defer mu.Unlock()
				file.Close()
				dir, _ := path.Split(fullFile)
				fileSuffix := path.Ext(fullFile)
				os.Rename(fullFile, dir+time.Now().Format("150405.0000")+fileSuffix)
				fullFile = logDir(&filePath, &fileName, true)
				file, _ = os.OpenFile(fullFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
				l.SetOutput(file)
			}
			/*
				fmt.Println(time.Now().Format("2006-01-02 15:04:05.0000"))
				fmt.Println(os.Getgid())
				fmt.Println(fileName)
			*/
		}
	}()
	l.SetOutput(file)
	return
}

func init() {
	//debugLogger.SetFlags(debugLogger.Flags() | log.Llongfile)
	//_, file, line, ok := runtime.Caller(99)

}

var debugLogger = access("debug.log", "log")

// DebugLog ，写DEBUG日志
func DebugLog(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	str := "file:" + file + "\tline:" + strconv.Itoa(line) + "\r\n\t"
	debugLogger.Println(str, args)

}

var accessLogger = access("access.log", "Access")

// AccessLog ，访问日志记录
func AccessLog(ctx *core.Context, startTime time.Time) {
	durTime := time.Now().Sub(startTime)
	// os.Getgid()
	rid := goroutine.GoroutineId()
	r := ctx.Req
	str := "routine_id:" + strconv.FormatInt(rid, 10) + " - " + tool.GetIp(r) + " - " + r.Method + " " + r.RequestURI + " " + r.Proto + " " + strconv.Itoa(ctx.GetStatusCode()) + " " + r.Header.Get("User-Agent") + " " + r.Header.Get("Accept-Language") + " " + strconv.FormatFloat(durTime.Seconds(), 'f', 4, 32) + "s"
	accessLogger.Println(str)

}

func logDir(logName *string, fileName *string, partDay bool) string {
	if strings.TrimSpace(*fileName) == "" {
		*fileName = "log.log"
	}
	if strings.TrimSpace(*logName) == "" {
		*logName = "Default"
	}
	dir, _ := os.Getwd()
	dir += config.LOGDIR + *logName + string(os.PathSeparator)
	if partDay {
		dir += time.Now().Format("20060102") + string(os.PathSeparator)
	}
	tool.MustCreateDir(dir)
	return dir + *fileName
}
