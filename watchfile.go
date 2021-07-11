package watchFile

import (
	 "time"
	 "os"
	 "fmt"
)

type OperationType int

const (
	Create OperationType = iota
	Remove
	Modify
)


type FileChangeMessage struct {
	File string
	Error error || nil = nil
}


func watchFile (file string, options ...time.Duration) chan bool{
	var checkingInterval time.Duration = 1 * time.Second
	if len(options) > 0 {
    	checkingInterval = options[0]
  	}

	fileChangeNotification := make(chan bool)
	
	pulse := interval(checkingInterval)
	
	go func(){
		var lastModTime int64 = 0
		for{
			<-pulse
			file, err := os.Stat(file)

			if err != nil {
				fmt.Println(err)
			}

			modTime := file.ModTime().Unix()
			
			if (lastModTime !=0 && lastModTime < modTime){
				fileChangeNotification <- true
			}

			lastModTime = modTime

		}
	}()
	return fileChangeNotification
}

type mainProc = func()
func forever(fn mainProc) {
	go func(){
		for{
			fn()
		}
	}()
	done := make(chan bool)
	<-done
}

func interval(period time.Duration) chan bool{
	pulse := make(chan bool)
	go forever( func(){
		for{
			time.Sleep(period)
			pulse <- true
		}
	})
	return pulse
}

