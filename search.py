# Automation using Python : Searching internet (Google) using simple keyboard shortcut.
# Author : Kumar Ashwin, krashwin00@gmail.com
# Version: 1.0, 15th Jan, 2020

import pyperclip
from selenium import webdriver

try:
    browser = webdriver.Firefox() #geckodriver(firefox) is available in the environment path
except:
    print("Firefox's driver is only specifed! Specify the driver of your choice of web browser.")

query = pyperclip.paste().replace(" ","+") #taking the last copied text and making it URL ready
search_url = "https://www.google.com/search?q="+query
browser.get(search_url)
