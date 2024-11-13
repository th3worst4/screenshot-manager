package main

import (
	"fmt"
	"os"
    "io/fs"
	//"fmt"
	"path"
	//"path/filepath"
	rlgui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var historyIndex uint8 = 0
var scrollIndex int32 = 0
var editSaveName bool = false
var show_warning int = 0

func RenderContentList(pwd *node){
    dir_entries, err := os.ReadDir((*pwd).value)
    if err != nil {
        panic(err)
    }
    fmt.Println(dir_entries)
}

func MainWindow(pwd *node, save_name string, source *os.File){
	tmp := new(node)
	
    rl.InitWindow(552, 280, "hyprshot-gui")
    defer rl.CloseWindow()
    
    rl.SetTargetFPS(60)

    gui_execution:
    for !rl.WindowShouldClose(){
        rl.BeginDrawing()
        rl.ClearBackground(rl.Color{255, 255, 255, 255})

        if rlgui.Button(rl.Rectangle{10, 10, 120, 24}, "#121#") {
            tmp = CreateNode(path.Dir((*pwd).value))
            pwd = InsertNewEntry(&pwd, tmp)
        }

//        switch rlgui.ToggleGroup(rl.Rectangle{10, 10, 38.5, 24 }, "#118#;#119#;#121#", -1) {
//            case 0:
//                if historyIndex != 0 {
//                }
//            case 1:
//            case 2:
//                tmp = CreateNode(path.Dir((*pwd).value))
//				pwd = InsertNewEntry(&pwd, tmp)
//        }   
        
        rlgui.TextBox(rl.Rectangle{140, 10, 402, 24}, &(*pwd).value, 128, false)
        
        if rlgui.TextBox(rl.Rectangle{10, 280, 346, 24}, &save_name, 128, editSaveName) {
            editSaveName = !editSaveName
        }

        switch rlgui.ListView(rl.Rectangle{10, 40, 120, 184}, "Home;Documents;Downloads;Music;Pictures;Videos", &scrollIndex, -1) {
            case 0:
                tmp = CreateNode(Home_Var)
				pwd = InsertNewEntry(&pwd, tmp)
            case 1:
                tmp = CreateNode(Documents)
				pwd = InsertNewEntry(&pwd, tmp)
            case 2:
                tmp = CreateNode(Downloads)
				pwd = InsertNewEntry(&pwd, tmp)
            case 3:
                tmp = CreateNode(Music)
				pwd = InsertNewEntry(&pwd, tmp)
            case 4:
                tmp = CreateNode(Pictures)
				pwd = InsertNewEntry(&pwd, tmp)
            case 5:
                tmp = CreateNode(Videos)
				pwd = InsertNewEntry(&pwd, tmp)
        } 

        RenderContentList(pwd)

        save_path := path.Join((*pwd).value, save_name)
		if rlgui.Button(rl.Rectangle{366, 280, 88, 24}, "Save") {
            exists, err := PathExists(save_path)
            if err != nil {
                panic("Error occurred while checking path existence")
            }
            if exists{
                show_warning = 1
            }else {
                if fs.ValidPath(save_path){
                    SaveFile(source, save_path)
                    break
                }else{
                    show_warning = 2
                }
            }
		}
        if show_warning == 1 {
            result := rlgui.MessageBox(rl.Rectangle{151, 90, 250, 100}, "#187#Warning", "Filename already exists. Overwrite it?", "Yes;No")
            if result == 1 {
                SaveFile(source, save_path)
                break gui_execution
            }
            if result == 0 {
                show_warning = 0
            }
        }
        if show_warning == 2{
            result := rlgui.MessageBox(rl.Rectangle{151, 90, 250, 100}, "#187#Warning", "Invalid file name", "Ok")
            if result >= 0 {
                show_warning = 0
            }

        }

		if rlgui.Button(rl.Rectangle{454, 280, 88, 24}, "Cancel") {
			break gui_execution
		}
        rl.EndDrawing()
    }
}

