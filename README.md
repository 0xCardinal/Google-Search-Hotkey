# Google-Search-Hotkey
Automation using Python : Search Google using simple keyboard shortcut.

## Getting Started
Clone the repository and specify a shortcut key for running the following command,
```shell 
 python3 /path/to/search.py
```

## Prerequisites
`python3` must be installed in your system.
<br />

And following Python modules are required and must be installed,
* pyperclip - used for accessing the clipboard
  ```shell 
  pip3 install pyperclip 
  ```

* webbrowser - opens a new tab in existing browser (if open) otherwise opens a new window in your default browser
  
  > You don't have to install it specifically, it is already present in Python's standard library.
  
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

## Author
**Kumar Ashwin** <br>
Email - **krashwin00@gmail.com**
