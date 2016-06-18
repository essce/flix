import requests, json
from flask import Flask, render_template, request

app = Flask(__name__)

@app.route('/')
def main():
	print('yoooo')
	return render_template('query.html')

@app.route('/', methods=['POST'])
def query_console():
	_title = request.form['title']
	_url = getURL(_title)

	r = requests.get(_url)
	getData(r)

	return render_template('query.html')

def getURL(title):
	title = addSpace(title)
	__url = "http://api.tvmaze.com/singlesearch/shows?q="
	__url += title
	__url += "&embed=episodes"
	print("url: ", __url)
	return __url

def testData(data):
	for a in result["_embedded"]["episodes"]:
		print(a)


def getData(response):
	print("Response code: ", response.status_code)
	result = response.json()
	print(type(result))

	try:
		#print(result["_embedded"])
		for a in result["_embedded"]["episodes"]:
			print(a)
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


if __name__ == "__main__":
    app.run()


