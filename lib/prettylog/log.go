package prettylog

import (
	"fmt"
	"log"
)

const (
	Reset   = "\033[0m"
	Info    = "\033[34m"
	Warning = "\033[33m"
	Error   = "\033[31m"
)

type PrettyLog struct {
	// LogLevel defines the minimum level of messages to log.
	infoLog  *log.Logger
	debugLog *log.Logger
	errorLog *log.Logger
}

func NewPrettyLog() *PrettyLog {
	return &PrettyLog{
		infoLog:  log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		debugLog: log.New(log.Writer(), "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLog: log.New(log.Writer(), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (p *PrettyLog) Info(msg string) {
	p.infoLog.SetPrefix(Info + p.infoLog.Prefix() + Reset)
	p.infoLog.Println(msg)
}
func (p *PrettyLog) Infof(msg string, args ...any) {
	p.infoLog.SetPrefix(Info + p.infoLog.Prefix() + Reset)
	formattedMsg := fmt.Sprintf(msg, args...)
	_ = p.infoLog.Output(2, formattedMsg)
}

func (p *PrettyLog) Debug(msg string) {
	p.debugLog.SetPrefix(Info + p.debugLog.Prefix() + Reset)
	_ = p.infoLog.Output(2, msg)
}

func (p *PrettyLog) Debugf(msg string, args ...any) {
	p.debugLog.SetPrefix(Info + p.debugLog.Prefix() + Reset)
	formattedMsg := fmt.Sprintf(msg, args...)
	_ = p.infoLog.Output(2, formattedMsg)
}

func (p *PrettyLog) Error(msg string) {
	p.errorLog.SetPrefix(Error + p.errorLog.Prefix() + Reset)
	_ = p.infoLog.Output(2, msg)

}
func (p *PrettyLog) Errorf(msg string, args ...any) {
	p.errorLog.SetPrefix(Error + p.errorLog.Prefix() + Reset)
	formattedMsg := fmt.Sprintf(msg, args...)
	_ = p.infoLog.Output(2, formattedMsg)
}

func (p *PrettyLog) Warning(msg string) {
	p.infoLog.SetPrefix(Warning + p.infoLog.Prefix() + Reset)
	_ = p.infoLog.Output(2, msg)
}
func (p *PrettyLog) Warningf(msg string, args ...any) {
	p.infoLog.SetPrefix(Warning + p.infoLog.Prefix() + Reset)
	formattedMsg := fmt.Sprintf(msg, args...)
	_ = p.infoLog.Output(2, formattedMsg)
}
