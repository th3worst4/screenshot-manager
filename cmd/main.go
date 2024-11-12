package main

import (
    "os"
    "path"
    "path/filepath"
    "fmt"
    //"strconv"
)

func main(){
    if len(os.Args) != 2 {
		panic("Invalid arguments")
    } 
    source := os.Args[1]
    
    spawn_path, err := filepath.Abs(path.Dir(source))
    if err != nil {
        panic(err)
    }
	pwd := CreateNode(spawn_path)
    
	file, err := os.Open(source);
    save_name := filepath.Base(source);
	if err != nil {
		panic(err)
	}
	defer file.Close()
    
	MainWindow(pwd, save_name, file)
    
    fmt.Println("something")
}
