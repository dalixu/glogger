package glogger

import "log"

//GLogger 定义logger接口
type GLogger interface {
	Trace(v ...interface{})
	Tracef(format string, v ...interface{})
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
}

var defaultEmptyLogger = EmptyLogger{}

//NewEmpty 返回空实现
func NewEmpty() GLogger {
	return &defaultEmptyLogger
}

//NewStd 返回标准库实现
func NewStd() GLogger {
	return &StdLogger{}
}

//Factory logFactory
type Factory interface {
	GetLogger(name string) GLogger
}

var factory Factory = EmptyLoggerFactory{}

//SetFactoryImpl 设置一个factory
func SetFactoryImpl(fy Factory) {
	factory = fy
}

//GetFactory 得到1个Facotry
func GetFactory() Factory {
	return factory
}

//EmptyLoggerFactory 空LoggerFactory
type EmptyLoggerFactory struct {
}

//GetLogger implement Factory
func (factory EmptyLoggerFactory) GetLogger(name string) GLogger {
	return NewEmpty()
}

//EmptyLogger 空实现
type EmptyLogger struct {
}

//Trace 打印Trace日志
func (eb *EmptyLogger) Trace(v ...interface{}) {
}

//Tracef 打印Trace日志
func (eb *EmptyLogger) Tracef(format string, v ...interface{}) {
}

//Debug Debug
func (eb *EmptyLogger) Debug(v ...interface{}) {
}

//Debugf 打印Debugf日志
func (eb *EmptyLogger) Debugf(format string, v ...interface{}) {
}

//Info 打印Info日志
func (eb *EmptyLogger) Info(v ...interface{}) {
}

//Infof 打印Infof日志
func (eb *EmptyLogger) Infof(format string, v ...interface{}) {
}

//Warn 打印Warn日志
func (eb *EmptyLogger) Warn(v ...interface{}) {
}

//Warnf 打印Warnf日志
func (eb *EmptyLogger) Warnf(format string, v ...interface{}) {
}

//Error 打印Error日志
func (eb *EmptyLogger) Error(v ...interface{}) {
}

//Errorf 打印Errorf日志
func (eb *EmptyLogger) Errorf(format string, v ...interface{}) {
}

//Fatal 打印Fatal日志
func (eb *EmptyLogger) Fatal(v ...interface{}) {
}

//Fatalf 打印Fatalf日志
func (eb *EmptyLogger) Fatalf(format string, v ...interface{}) {
}

//StdLoggerFactory 标准库实现的Factory
type StdLoggerFactory struct {
}

//GetLogger implement Factory
func (factory StdLoggerFactory) GetLogger(name string) GLogger {
	return NewStd()
}

//StdLogger 默认的标准库logger实现
type StdLogger struct {
}

//Trace 打印Trace日志
func (eb *StdLogger) Trace(v ...interface{}) {
	log.Println(v...)
}

//Tracef 打印Trace日志
func (eb *StdLogger) Tracef(format string, v ...interface{}) {
	log.Printf(format, v...)
}

//Debug Debug
func (eb *StdLogger) Debug(v ...interface{}) {
	log.Println(v...)
}

//Debugf 打印Debugf日志
func (eb *StdLogger) Debugf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

//Info 打印Info日志
func (eb *StdLogger) Info(v ...interface{}) {
	log.Println(v...)
}

//Infof 打印Infof日志
func (eb *StdLogger) Infof(format string, v ...interface{}) {
	log.Printf(format, v...)
}

//Warn 打印Warn日志
func (eb *StdLogger) Warn(v ...interface{}) {
	log.Println(v...)
}

//Warnf 打印Warnf日志
func (eb *StdLogger) Warnf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

//Error 打印Error日志
func (eb *StdLogger) Error(v ...interface{}) {
	log.Println(v...)
}

//Errorf 打印Errorf日志
func (eb *StdLogger) Errorf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

//Fatal 打印Fatal日志
func (eb *StdLogger) Fatal(v ...interface{}) {
	log.Println(v...)
}

//Fatalf 打印Fatalf日志
func (eb *StdLogger) Fatalf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
