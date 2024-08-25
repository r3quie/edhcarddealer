# import

def riffle(deck): # returns a list of indexes that are the result of a perfect riffle shuffle
    indexes = range(len(deck))
    return [x for t in zip(indexes[len(indexes)//2:], indexes[:-len(indexes)//2]) for x in t]

def arrange(deck): # rearanges the deck according to the indexes
    return [deck[i] for i in riffle(deck)]

def put_hand_on_bottom(deck):
    for _ in range(7):
        deck.append(deck.pop(0))
    return deck


if __name__ == "__main__":
    deck = [*range(100)]
    print(riffle(deck))
    print(arrange(deck))
    print(put_hand_on_bottom(deck))