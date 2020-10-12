# _*_ coding: utf-8 _*_

import requests
from bs4 import BeautifulSoup

def is_response(url):
    try:
        response = requests.get(url)
        if response.status_code == 200:
            return response.text
        else:
            print('Please check the url!')
    except requests.RequestException:
        return None

...

def main():
    url = 'http://xxx' + str(var)
    html = is_response(url)
    soup = BeautifulSoup(html, 'lxml')
    pass

if __name__ == '__main__':
    main()
