# Install dependencies 
* sudo pacman -Syu jq grim slurp wl-clipboard libnotify hyprpicker*

# Download script grimblast and add to him chmod +x ./grimblast

# Base command
* grimblast save active (Save the screenshot of active window)
* grimblast save area (Save the screenshot of a rectangular area selectable with mouse.)
* grimblast --notify copy active|area|output|screen (Take the screenshot of respective area and copy it to clipboard and notify the user.)
* grimblast --cursor copysave area (Copy the screenshot of the selected window area along with cursor to the clipboard. Also save the image to Pictures directory.)

# And if you want add to PATH
* sudo cp ./grimblast /usr/local/bin/
