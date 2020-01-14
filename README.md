# Google-Search-Hotkey
Automation using Python : Search Google using simple keyboard shortcut.

## Getting Started
Clone the repository and specify a shortcut key for running the following command,
```shell 
 python3 /path/to/search.py
```
`/path/to/the/directory/` is the directory location, that needs to be uncluttered.

## Prerequisites
`python3` must be installed in your system.

And following Python modules are required and must be installed,
*pyperclip - used for accessing the clipboard
  ```shell pip3 install pyperclip ```

*selenium
  ```shell pip3 install selenium ```
  Selenium also requires web drivers to get control over web browsers.
    Chrome:	https://sites.google.com/a/chromium.org/chromedriver/downloads
    Edge:	https://developer.microsoft.com/en-us/microsoft-edge/tools/webdriver/
    Firefox:	https://github.com/mozilla/geckodriver/releases
    Safari:	https://webkit.org/blog/6900/webdriver-support-in-safari-10/ 
    
   Set environment paths for respective drivers.
   Note: if chromedriver is not in your path, you’ll need to add it here
   ```python driver = webdriver.Chrome(r'C:\Python\chromedriver.exe')```
  

## Author
Kumar Ashwin
Email - krashwin00@gmail.com
