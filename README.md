
forked from [alecthomas/log4go](https://github.com/alecthomas/log4go)

>  [quick start](#QuickStart)
> 
>  [user guide](#UserGuide)

[中文](http://c-dafan.github.io/04/26/log4go用户指导/)

Installation:
- Run `go install log4go.googlecode.com/hg`

Usage:
- Add the following import:
import l4g "log4go.googlecode.com/hg"

Acknowledgements:
- pomack
  For providing awesome patches to bring log4go up to the latest Go spec


## QuickStart

### installation

```bash
go get github.com/alecthomas/log4go
```

### import

```go
import l4g "github.com/alecthomas/log4go"
```

### get logger

#### Global logger

```go
l4g.Info("hello world")
defer l4g.Close()
```

#### NewDefaultLogger

```go
log := l4g.NewDefaultLogger(l4g.INFO)
log.Info("hello world")
defer log.Close()
```

#### l4g.Logger

```go
log := make(l4g.Logger)
defer log.Close()
log.AddFilter("stdout", l4g.DEBUG, l4g.NewConsoleLogWriter())
log.Info("hello world")
```

## output log

```go
l4g.Finest()
l4g.Fine()
l4g.Debug()
l4g.Trace()
l4g.Info()
l4g.Warning()
l4g.Error()
l4g.Critical()

```

## UserGuide

### Level

```go
	FINEST
	FINE
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	CRITICAL
```

### Writer

#### ConsoleLogWriter


#### FileLogWriter


#### configFile

```xml
<logging>
  <filter enabled="true">
    <tag>stdout</tag>
    <type>console</type>
    <!-- level is (:?FINEST|FINE|DEBUG|TRACE|INFO|WARNING|ERROR) -->
    <level>DEBUG</level>
  </filter>
  <filter enabled="true">
    <tag>file</tag>
    <type>file</type>
    <level>FINEST</level>
    <property name="filename">test{}.log</property>
    <!--
       %T - Time (15:04:05 MST)
       %t - Time (15:04)
       %D - Date (2006/01/02)
       %d - Date (01/02/06)
       %L - Level (FNST, FINE, DEBG, TRAC, WARN, EROR, CRIT)
       %S - Source
       %M - Message
       It ignores unknown format strings (and removes them)
       Recommended: "[%D %T] [%L] (%S) %M"
    -->
    <property name="format">[%D %T] [%L] (%S) %M</property>
    <property name="rotate">false</property> <!-- true enables log rotation, otherwise append -->
    <property name="maxsize">0M</property> <!-- \d+[KMG]? Suffixes are in terms of 2**10 -->
    <property name="maxlines">0K</property> <!-- \d+[KMG]? Suffixes are in terms of thousands -->
    <property name="maxbackup">999</property> 
	<property name="daily">true</property> <!-- Automatically rotates when a log message is written after midnight -->
  </filter>
</logging>
```

```json
[
  {
    "enabled": true,
    "Tag": "stdout",
    "Type": "console",
    "Level": "DEBUG"
  },
  {
    "enabled": true,
    "Tag": "file",
    "Type": "file",
    "Level": "DEBUG",
    "filename": "test.log",
    "format": "[%D %T] [%L] (%S) %M",
    "rotate": true,
    "maxsize": "0M",
    "maxlines": "3",
    "maxbackup": 3,
    "daily": false
  }
]
```

[XMLConfigurationExample.go](https://github.com/c-dafan/log4go/blob/master/examples/XMLConfigurationExample.go)

[JsonConfigExample.go](https://github.com/c-dafan/log4go/blob/master/examples/JsonConfigExample.go)

## issue

if the main thread ends, the logger thread will also end and part of logs will be lose.Suggest write sleep in the end of the code.

```go
defer func(){
	time.Sleep(time.Second)
}
```