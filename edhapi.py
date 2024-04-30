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

def get_uri():
    response = requests.get('https://api.scryfall.com/bulk-data/oracle-cards')
    uri = response.json()['download_uri']
    print(uri)
    return uri


if __name__ == "__main__":
    get_uri()
