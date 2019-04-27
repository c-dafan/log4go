package main

import "time"
import l4g "github.com/c-dafan/log4go"

func main() {
	// Load the configuration (isn't this easy?)
	l4g.LoadConfiguration("example.xml")
	// And now we're ready!
	l4g.Finest("This will only go to those of you really cool UDP kids!  If you change enabled=true.")
	for i := 1; i < 20; i++ {
		l4g.Debug("Oh no!  %d + %d = %d!", 2, 2, 2+2)
	}
	l4g.Info("About that time, eh chaps?")
	time.Sleep(time.Second)
	defer l4g.Close()
}
