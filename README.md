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

* selenium
  ```shell 
  pip3 install selenium 
  ```
  Selenium also requires web drivers to get control over web browsers. <br/>WebDrivers can be downloaded from the below-mentioned links,
   > Chrome:	https://sites.google.com/a/chromium.org/chromedriver/downloads <br />
   > Edge:	https://developer.microsoft.com/en-us/microsoft-edge/tools/webdriver/ <br />
   > Firefox:	https://github.com/mozilla/geckodriver/releases <br />
   > Safari:	https://webkit.org/blog/6900/webdriver-support-in-safari-10/ <br />
    
   Set environment paths for respective drivers. <br />
   E.g.: if chromedriver is not in your path, you’ll need to add it here
     ```python 
     driver = webdriver.Chrome(r'C:\Python\chromedriver.exe')
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

## Author
**Kumar Ashwin** <br>
Email - **krashwin00@gmail.com**
