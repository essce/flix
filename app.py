import os
import requests, json
from random import randint
from flask import Flask, render_template, request
from Show import Episode

app = Flask(__name__)

#global variables#######
allEpisodes = []
#global variables#######

@app.route('/')
def main():
	return render_template('query.html')

@app.route('/', methods=['POST'])
def query_console(placeholder=None, showname=None, name=None, season=None, episode=None):

	try: 
		_title = request.form['title']
		if not _title:
			return render_template('query.html', placeholder="Input cannot be empty")
			
		_url = getURL(_title)

		r = requests.get(_url)
		ep = getData(r)

		showname = ep.showname
		name = ep.name
		season = ep.seasonNum
		episode = ep.episodenum

	except ValueError as e:
		print(e)
		return render_template('query.html', placeholder="Please enter a proper TV show name")

	return render_template('result.html', showname=showname, name=name, season=season, episode=episode)

def query_console(showname=None, name=None, season=None, episode=None):
	_title = request.form['title']

	if not _title:
		try: 
			raise ValueError("Search is empty.")
		except ValueError as err:
			print(err.args)
		finally:
			return render_template('query.html')

	_url = getURL(_title)

	r = requests.get(_url)
	ep = getData(r)
	return render_template('result.html', showname=showname, name=name, season=season, episode=episode)


@app.route('/reroll', methods=['GET'])
def reroll(showname=None, name=None, season=None, episode=None):
	
	ep = randomize()
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
	return __url

def getData(response):
	result = response.json()
	return __url

def getData(response):
	result = response.json()

	try:
		alleps = result["_embedded"]["episodes"]
		episodes = []

		for a in alleps: 				#for each episode
			epName = a["name"] 			#get name
			seasonNum = a["season"] 	#get season
			epNum = a["number"]			#get episode number

			episode = Episode(epName, seasonNum, epNum, result["name"]) 	#make an Episode obj
			episodes.append(episode)

		# Episodes list generated 
		global allEpisodes
		allEpisodes = episodes 

		return randomize()

	except ValueError as e:
		print(e)

def randomize():
	randomEp = randint(0, len(allEpisodes)-1)
	ep = allEpisodes[randomEp]
	return ep

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

