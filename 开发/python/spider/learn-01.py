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

def show_info(soup):
    list = soup.find(class_='Box-body px-5 pb-5').find_all('li')
    for item in list:
        item_name = item.find('a').string
        item_links = item.find('a').get('href')
        print('<<' + item_name + '>> ' item_links)

def main():
    url = 'https://github.com/wistbean/learn_python3_spider'
    html = is_response(url)
    soup = BeautifulSoup(html, 'lxml')
    show_info(soup)

if __name__ == '__main__':
    main()
