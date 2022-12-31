from flask import Flask,request
from rutermextract import TermExtractor
term_extractor = TermExtractor()

app = Flask(__name__)

@app.route('/', methods=['GET'])
def add_message():
    text = request.args.get('text')
    print(text)
    response = ""
    for word in term_extractor(text, limit=10, strings=True):
      response =response + word + ","
    response = response[:-1]
    print(response)
    return response


app.run(host='0.0.0.0', port=81)
