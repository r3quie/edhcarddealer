import json

with open("cache/oracle-cards-20240419090153.json", "r", encoding="utf8") as f:
    data =json.load(f)


######################################## UNDER CONSTRUCTION ########################################

cardname = "Replicating Ring"

def get_values(cardname):
    for i in data:
        if i["name"] == cardname:
            try:
                print(i["produced_mana"])
            except KeyError:
                pass
            break

