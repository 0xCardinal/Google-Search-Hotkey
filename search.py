#!/usr/bin/env python3

# Automation using Python : Searching internet (Google) using simple keyboard shortcut.
# Author : Kumar Ashwin, krashwin00@gmail.com
# Version: 2.0, 29th May, 2020

import pyperclip, webbrowser

query = pyperclip.paste().replace(" ","+") #taking the last copied text and making it URL ready
search_url = "https://www.google.com/search?q="+query

webbrowser.open_new_tab(search_url)
