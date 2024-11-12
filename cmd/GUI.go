package main

import(
    "os"
    //"fmt"
	"path"
    //"path/filepath"
	rl "github.com/gen2brain/raylib-go/raylib"
	rlgui "github.com/gen2brain/raylib-go/raygui"
)

var historyIndex uint8 = 0
var scrollIndex int32 = 0
var editSaveName bool = false

func MainWindow(pwd *node, save_name string, source *os.File){
	tmp := new(node)
	
    rl.InitWindow(552, 280, "hyprshot-gui")
    defer rl.CloseWindow()
    
    rl.SetTargetFPS(60)

    for !rl.WindowShouldClose(){
        rl.BeginDrawing()
        rl.ClearBackground(rl.Color{255, 255, 255, 255})
        /*
            Home
            Documents
            Downloads
            Music
            Pictures
            Videos
        */ 
        switch rlgui.ToggleGroup(rl.Rectangle{10, 10, 38.5, 24 }, "#118#;#119#;#121#", -1) {
        /*
            lacks full implementation
        */
            case 0:
                if historyIndex != 0 {
                }
            //case 1:
            case 2:
                tmp = CreateNode(path.Dir((*pwd).value))
				pwd = InsertNewEntry(&pwd, tmp)
        }   
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

		if rlgui.Button(rl.Rectangle{366, 280, 88, 24}, "Save") {
            SaveFile(source, path.Join((*pwd).value, save_name))
            break
		}
		if rlgui.Button(rl.Rectangle{454, 280, 88, 24}, "Cancel") {
			break
		}
        rl.EndDrawing()
    }
}
