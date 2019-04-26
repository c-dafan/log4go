
forked from [alecthomas/log4go](https://github.com/alecthomas/log4go)

>  [quick start](#QuickStart)
> 
>  [user guide](#UserGuide)



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

### 