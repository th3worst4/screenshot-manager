# screenshot-manager
Simple File Manager application built in Go using the [RayGUI](https://github.com/raysan5/raygui) library made specifically to move screenshots.

# Motivation
Recently, I moved on from KDE Plasma to [Hyprland](https://hyprland.org/), and even though Hyprland provides the [Hyprshot](https://github.com/Gustash/hyprshot) application and its integration with
[swaync](https://github.com/ErikReider/SwayNotificationCenter) is pretty good; it was lacking the ability to manipulate screenshots as on [Spectacle](https://github.com/KDE/spectacle). So I built my
own application to provide this support.

The purpose of it was to allow the user to change the name and directory of a screenshot by simply clicking on the swaync notification.

# Building
There is a Makefile to build the project.
```
make build
```
This will create a `build/` directory with an executable.

# swaync integration
The following snippet presents my swaync configuration to integrate this project with it
```json
"scripts": {
        "click-on-print": {
                "exec": "sh -c '/home/caioc/.config/swaync/open-pictures.sh \"$SWAYNC_BODY\"'",
                "app-name": "Hyprshot",
                "run-on": "action" 
        }
}
```
The `open-pictures.sh` bash file is also provide within this repository.
