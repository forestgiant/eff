# eff
Effulgent Drawing API
---
This API provides a way to easily create graphics programs in go.
The sdl package is a partial wrapper of sdl for go.  For a complete wrapper checkout go-sdl2 <https://github.com/veandco/go-sdl2>

###SDL setup
* OSX: `brew install sdl2{,_mixer,_image,_ttf}`
* Arch Linux: `sudo pacman -S sdl2{,_mixer,_image,_ttf}`
* Ubuntu/Debian: `sudo apt-get install libsdl2{,-mixer,-image,-ttf}-dev `
* Windows:
    1. Install mingw <http://www.mingw.org/>, ensure the `bin` folder is in the windows path
    2. Download the windows development sdl2 libraries:
        1. SDL2 <https://www.libsdl.org/download-2.0.php>
        2. SDL2_ttf <https://www.libsdl.org/projects/SDL_ttf/>
        3. SDL2_mixer <https://www.libsdl.org/projects/SDL_mixer/>
        4. SDL2_image <https://www.libsdl.org/projects/SDL_image/>
    3. Extract each tarball to the same directory (i.e. `c:\mingw_dev_lib`) and add the bin(i.e. `c:\mingw_dev_lib\bin`) folder to the PATH enviroment variable.  Currently only the 32bit(i686) libraries are supported.
    4. Update the cgo comment at the top of sdl.go to ensure that the include path and lib path match where you extracted the libraries
    5. When building set the `GOARCH=386` and `CGO_ENABLED=1`. Use the `SET` command if you are using the normal windows command line and not git-bash

### Keyboard control
* Press `f` to toggle fullscreen
* Press `q` to quit the program
