import json
#import os

#print(os.getcwd()) #prints working pwd, needed to se correct relative path in open fnc
with open("/home/superadmin/site/mtg/mana_calc/edhcarddealer/cache/oracle-cards.json", "r", encoding="utf8") as f:
    scryjson =json.load(f)

def get_values(cardname):
    for i in scryjson:
        if i["name"] == cardname:
            try:
                return(i["produced_mana"])
            except KeyError:
                pass
            break

if __name__ == "__main__":
    print(get_values("Replicating Ring"))
