import requests
import json

url = 'https://recyclemap.ru/index.php?option=com_greenmarkers&task=get_json&type=points&tmpl=component'

url_city = "https://recyclemap.ru/index.php?option=com_greenmarkers&task=get_json&type=cities&tmpl=component"

response = requests.post(url_city)
cities = response.json()
result = dict()
for city in cities:
	response = requests.post(url, data={'city': city['id'], 'layer':'0', 'gos':'0'})
	data = response.json()
	result = { **result , **data }

# response = requests.post(url, data={'city':'1', 'layer':'0', 'gos':'0'})
# data = response.json()
# f = open("cities.json", 'w', encoding = 'UTF-8')
# f = open("points.json", 'w', encoding = 'UTF-8')

with open('points.json', 'w', encoding='utf-8') as f:
    json.dump(result, f, ensure_ascii=False, indent=4)

