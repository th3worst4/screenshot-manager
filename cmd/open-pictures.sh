#!/usr/bin/env bash

IFS=">" read -ra BODY <<< "$1"

IFS="<" read -ra PATH <<< "${BODY[1]}"

/home/caioc/.config/swaync/screenshot-manager "$PATH"
