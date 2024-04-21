import requests # pip3 install requests
import json

######################################## UNDER CONSTRUCTION ########################################

def api_handle(card_id):
    try:
        response = requests.get('https://api.scryfall.com/cards/' + card_id + '?format=json')
    
        # Check if the request was successful (status code 200)
        if response.status_code == 200:
            # Try to access the 'produced_mana' key in the JSON response
            try:
                produced_mana = response.json()['produced_mana']
                print(produced_mana)
            except KeyError:
                pass
        else:
            pass

    except requests.RequestException as e:
        pass

filej = open('cache/oracle-cards-20240419090153.json', 'r')

card_preid = "03415c42-086e-4a2e-9be8-5cdcde83f134"
jsondata = filej.read()
obj = json.loads(jsondata)
print(obj['produced_mana'])

def json_handle(card_id, path):
    json.loads()
    produced_mana = path.json()[card_id][produced_mana]
    print(produced_mana)

json_handle(card_id, path)
card_id = "f1e6e3a2-5145-4898-990e-150e090f5206"
if __name__ == "__main__":
    api_handle(card_id)