#!/usr/bin/env python
# coding: utf-8

# In[11]:


import requests
import random 
import json
import re


# In[ ]:


from typing import Optional

from fastapi import FastAPI

app = FastAPI()


@app.get("/data")
def read_data():
    rannm = random.randint(1, 3075362)
    rancomp = requests.get('https://data.gov.au/data/api/3/action/datastore_search?resource_id=cb7e4eb5-ed46-4c6c-97a0-4532f4479b7d&limit=1&offset={}'.format(random.randint(1, 3075362))).json()
    res = re.sub(' +', ' ', rancomp['result']['records'][0]['Company Name'])

    
    currency = requests.get('https://api.currencyfreaks.com/latest?apikey=4f7765e522f84a37ab63672b887ed45e')
    person = requests.get("https://randomuser.me/api/").json()
    scamname = person['results'][0]['name']
    name = '{} {} {}'.format(scamname['title'], scamname['first'], scamname['last'])
    location = person['results'][0]['location']
    scamloc = "{}, {}".format(location['city'], location['country'])
    return ({"name" : name, "location" : scamloc,
        "company": res, "currency" : random.choice(list(currency.json()['rates'].keys()))})


# In[73]:


person = requests.get("https://randomuser.me/api/").json()


# In[78]:


person['results'][0]['location']


# In[ ]:





# In[ ]:





# In[ ]:





# In[ ]:





# In[ ]:




