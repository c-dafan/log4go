package main

import (
	"time"
)

import l4g "github.com/c-dafan/log4go"

const (
	filename = "flw.log"
)

func main() {

	log := make(l4g.Logger)
	flw := l4g.NewFileLogWriter(filename, true)
	flw.SetFormat("[%D %T] [%L] (%S) %M")
	flw.SetRotate(true)
	flw.SetRotateSize(0)
	flw.SetRotateLines(2)
	flw.SetRotateDaily(false)
	flw.SetRotateMaxBackup(2)
	log.AddFilter("file", l4g.FINE, flw)

	// Log some experimental messages
	log.Finest("Everything is created now (notice that I will not be printing to the file)")
	for i := 1; i < 20; i++ {
		log.Info("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
	}
	log.Critical("Time to close out!")
	time.Sleep(time.Second)
	// Close the log
	log.Close()

	// Print what was logged to the file (yes, I know I'm skipping error checking)
	//fd, _ := os.Open(filename)
	//in := bufio.NewReader(fd)
	//fmt.Print("Messages logged to file were: (line numbers not included)\n")
	//for lineno := 1; ; lineno++ {
	//	line, err := in.ReadString('\n')
	//	if err == io.EOF {
	//		break
	//	}
	//	fmt.Printf("%3d:\t%s", lineno, line)
	//}
	//fd.Close()
	//
	//// Remove the file so it's not lying around
	//os.Remove(filename)
}
