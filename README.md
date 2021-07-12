# Watchfile
a Golang module to trigger a file changes

## Usage

```go
	
    fileChangeNotification, fileCheckError := watchfile.Notify("my-file.txt");

	for {
		select {
			case err := <- fileCheckError:
				fmt.Println(err)
			case <- fileChangeNotification:
				fmt.Println("file changed")
		}
	}
    
```

