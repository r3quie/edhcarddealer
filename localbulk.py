import json

with open("cache/oracle-cards-20240419090153.json", "r", encoding="utf8") as f:
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