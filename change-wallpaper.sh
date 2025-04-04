#!/bin/bash

# Find all image files (jpg, png, gif) in the Downloads directory
paths=$(find ~/Downloads/ -type f \( -iname "*.jpg" -o -iname "*.png" -o -iname "*.gif" \))

# Let the user select a wallpaper, but only show the filenames in rofi
WALLPAPER=$(echo "$paths" | xargs -n 1 basename | rofi -dmenu -p "Set wallpaper:" -theme ~/.config/rofi/cfg2.rasi)

# If the user selected a wallpaper, find the full path and set it using swww
if [ -n "$WALLPAPER" ]; then
    # Find the full path based on the selected filename
    FULL_PATH=$(find ~/Downloads/ -type f -iname "$WALLPAPER")

    # Set the wallpaper using the full path
    swww img "$FULL_PATH" --transition-type=random --transition-duration=3
fi
