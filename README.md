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
<br />

And following Python modules are required and must be installed,
* pyperclip - used for accessing the clipboard
  ```shell 
  pip3 install pyperclip 
  ```
  
## Usage
This tool when binded with the shortcut keys can be very much helpful in order to search stuff on Google. <br> <br>
For efficient use of this script, it is recommended to create keyboard shortcut for running the script. 
Copy `Ctrl+C` the text you want to search for and press the specified hotkeys to directly search on Google. <br> <br>
This tool searches for the last copied text on internet using Google and is a major time saver.

## Creating Shortcuts
**Windows**
> Create a batch file (MyScript.bat) with `python3 /path/to/search.py` in it
and make a shortcut of the batch file on Desktop and specify the shortcut key in the Properties section.

![picture alt](https://i.stack.imgur.com/eMpiM.png "Bat Shortcut Properties")

**Linux**
> Go to Settings → Keyboards and add command `python3 /path/to/search.py` and shortcut key in there.

## Troubleshooting
**Linux**
If `Pyperclip could not find a copy/paste mechanism for your system.` error occurs, then install the following module and re-run the tool.
```shell
sudo apt install xsel
```

## Author
**Kumar Ashwin** <br>
Email - **krashwin00@gmail.com**
