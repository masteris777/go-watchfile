package watchfile

import (
	"os"
	"time"
)

func Notify (file string, options ...time.Duration) (chan bool, chan error){
	var checkingInterval time.Duration = 1 * time.Second
	if len(options) > 0 { checkingInterval = options[0] }

	fileChangeNotificationCh := make(chan bool)
	fileReadingErrorNotificationCh := make(chan error)
	
	go func(){

		var lastModTime int64 = 0
		fileCheckingError := false
		
		for {

			time.Sleep(checkingInterval)

			file, err := os.Stat(file)
			
			if (err != nil) {

				if (!fileCheckingError) { 
					fileReadingErrorNotificationCh <- err
				}

			} else {

				modTime := file.ModTime().Unix()
				
				if (lastModTime !=0 && lastModTime < modTime){
					fileChangeNotificationCh <- true
				}

				lastModTime = modTime
			}

			fileCheckingError = (err != nil)

		}
	}()
	return fileChangeNotificationCh, fileReadingErrorNotificationCh
}
