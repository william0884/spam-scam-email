#!/usr/bin/env python
# coding: utf-8

# In[2]:


import requests
import random 
import json
import re
import configparser


# In[5]:


config = configparser.ConfigParser()
config.sections()

#config.read('config.ini')
config.read('config.ini')

config.sections()


# In[6]:


key = config['DEFAULT']['apikey']


# In[1]:


from typing import Optional

from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()

origins = [
    "http://localhost.tiangolo.com",
    "https://localhost.tiangolo.com",
    "http://localhost",
    "http://localhost:9887"
    "http://172.24.252.203:9887",
    "*",
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

@app.get("/data")
def read_data():
    
    rannm = random.randint(1, 3075362)
    rancomp = requests.get('https://data.gov.au/data/api/3/action/datastore_search?resource_id=cb7e4eb5-ed46-4c6c-97a0-4532f4479b7d&limit=1&offset={}'.format(random.randint(1, 3075362))).json()
    res = re.sub(' +', ' ', rancomp['result']['records'][0]['Company Name'])


    currency = requests.get('https://api.currencyfreaks.com/latest?apikey={}'.format(key))
    currencyvalue = random.choice(list(currency.json()['rates'].keys()))
    dollaramount = random.randint(1,100)
    dayransom = random.randint(1,10)

    person = requests.get("https://randomuser.me/api/").json()
    scamname = person['results'][0]['name']
    name = '{} {} {}'.format(scamname['title'], scamname['first'], scamname['last'])
    location = person['results'][0]['location']
    scamloc = "{}, {}".format(location['city'], location['country'])

    webhist = requests.get('http://localhost:5600/api/0/buckets/aw-watcher-web-chrome/events?limit=1').json()

    timestamp = webhist[0]['timestamp']
    duration = webhist[0]['duration']
    webhisty = webhist[0]['data']['url']
    bodytext = """Hello, My name is {} and I am located in {}. My company is {}. On {} you visited the url {}. You must send {} {} via this link: http://example.com within {} days of opening this email. If you do not I will tell your friends and family that you visited the site. Be kind, {}""".format(name, scamloc, res, timestamp, webhisty, dollaramount, currencyvalue, dayransom, name)
    
    return(bodytext)


# In[ ]:




