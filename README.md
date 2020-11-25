# Google-Search-Hotkey v3.0
Search Google using simple keyboard shortcut.

## Getting Started
If you have go installed,
```
go get github.com/krAshwin/Google-Search-Hotkey
```
and you are good to go!

You can also download the precompiled binaries from [Linux](https://github.com/krAshwin/Google-Search-Hotkey/blob/master/linux/Google-Search-Hotkey) and [Windows](https://github.com/krAshwin/Google-Search-Hotkey/blob/master/windows/Google-Search-Hotkey.exe) directories.
## Prerequisites
For linux - xclip, xsel or wl-clipboard, any one must be installed.
```shell
sudo apt install xclip
sudo apt install xsel
sudo apt install wl-clipboard
```
  
## Usage
I have migrated this code on Go, and have added some additional functionality.
### Release Notes
1. If a URL is identified in the clipboard or passed along with the tool, it directly navigates to the specified URL.
2. Support for Mac OS.

This tool when binded with the shortcut keys can be very much helpful in order to search stuff on Google. <br> <br>
For efficient use of this script, it is recommended to create keyboard shortcut for running the script. 
Copy the text you want to search for and press the specified hotkeys to directly search on Google. <br> <br>
This tool searches for the last copied text on internet using Google and is a major time saver.

## Creating Shortcuts
**Windows**
> Create a shortcut of the windows' version of tool and make a shortcut of the file on Desktop and specify the shortcut key in the Properties section. And maybe hide it if you don't need to see it often ;)

**Linux**
> Go to Settings → Keyboards and add command `Google-Search-Hotkey` and shortcut key in there.


## Author
**Kumar Ashwin** <br>
Email - **krashwin00@gmail.com**
