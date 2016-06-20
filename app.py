import os
import requests, json
from random import randint
from flask import Flask, render_template, request
from Show import Episode

app = Flask(__name__)

@app.route('/')
def main():
	print('yoooo')
	return render_template('query.html')

@app.route('/', methods=['POST'])
def query_console(showname=None, name=None, season=None, episode=None):
	_title = request.form['title']
	_url = getURL(_title)

	r = requests.get(_url)
	ep = getData(r)

	showname = ep.showname
	name = ep.name
	season = ep.seasonNum
	episode = ep.episodenum

	return render_template('result.html', showname=showname, name=name, season=season, episode=episode)

def getURL(title):
	title = addSpace(title)
	__url = "http://api.tvmaze.com/singlesearch/shows?q="
	__url += title
	__url += "&embed=episodes"
	print("url: ", __url)
	return __url

def getData(response):
	print("Response code: ", response.status_code)
	result = response.json()
	print(type(result))

	try:
		alleps = result["_embedded"]["episodes"]
		episodes = []

		for a in alleps: #for each episode
			epName = a["name"] 			#get name
			seasonNum = a["season"] 	#get season
			epNum = a["number"]			#get episode number

			episode = Episode(epName, seasonNum, epNum, result["name"]) 	#make an Episode obj
			episodes.append(episode)

		randomEp = randint(0, len(episodes)-1)
		getEp = episodes[randomEp]
		return getEp

	except ValueError as e:
		print(e)

def addSpace(title):
	newTitle=""
	for letter in title: 
		if letter == " ":
			newTitle += "%20"
		else:
			newTitle += letter	

	return newTitle


port = os.getenv('PORT', '5000')
if __name__ == "__main__":
    app.run(host='0.0.0.0', port=int(port))

