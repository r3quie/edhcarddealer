import json
from pathlib import Path
#import os

#print(os.getcwd()) #prints working pwd, needed to se correct relative path in open fnc

try:
    with open(str(Path.home()) + "/site/mtg/mana_calc/edhcarddealer/cache/oracle-cards.json", "r", encoding="utf8") as f:
        scryjson =json.load(f)
except FileNotFoundError:
    with open(str(Path.home()) + "/gitrepos/edhcarddealer/cache/oracle-cards.json", "r", encoding="utf8") as f:
        scryjson = json.load(f)

def get_values(cardname):
    commander_mana = False
    for i in scryjson:
        if i["name"] == cardname:
            try:
                produced_mana = i["produced_mana"]
            except KeyError:
                produced_mana = None

            try:
                color_identity = i["color_identity"]
            except KeyError:
                color_identity = None

            try:
                if "dd one mana of any color in your commander's color identity" in i["oracle_text"]:
                    commander_mana = True
            except KeyError:
                pass

            return produced_mana, color_identity, commander_mana


if __name__ == "__main__":
    print(get_values("Obeka, Splitter of Seconds"))
