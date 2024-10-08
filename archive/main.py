import random
#from edhapi import api_handle
from localbulk import get_values
from getdeck import get_deck
from mulligan import riffle, arrange

def simdeal(deck):
    hand = random.sample(deck, 7)

    blue = red = black = green = white = cless = 0
    for ii in range(len(hand)):
        if "U" in hand[ii]:
            blue += 1
        if "B" in hand[ii]:
            black += 1
        if "R" in hand[ii]:
            red += 1
        if "G" in hand[ii]:
            green += 1
        if "W" in hand[ii]:
            white += 1
        if "C" in hand[ii]:
            cless += 1

    rest = hand.count("n")

    return blue, red, black, green, white, cless, rest

def main(numofsim, imported_deck):
    deck = []
    deck_identity = []
    deckl, cardcount = get_deck(imported_deck)
    commander_count = 0

    for i in range(len(deckl)):
        m_value, card_identity, commander_mana = get_values(deckl[i])

        if commander_mana is True:
            commander_count += cardcount[i]
            continue
        if m_value == None:
            card = "n"
        else:
            card = "".join(m_value)
        deck += [card] * cardcount[i]
        deck_identity += [color for color in card_identity if color not in deck_identity]

    deck_identity = "".join(deck_identity)
    
    deck += [deck_identity] * commander_count
 
    totalblue = totalred = totalblack = totalgreen = totalwhite = totalcless = totalrest = 0
    for _ in range(numofsim):
        blue, red, black, green, white, cless, rest = simdeal(deck)
        totalblue += blue
        totalred += red
        totalblack += black
        totalgreen += green
        totalwhite += white
        totalcless += cless
        totalrest += rest

    avg_blue = str(round(totalblue / numofsim, 2))
    avg_red = str(round(totalred / numofsim, 2))
    avg_black = str(round(totalblack / numofsim, 2))
    avg_green = str(round(totalgreen / numofsim, 2))
    avg_white = str(round(totalwhite / numofsim, 2))
    avg_cless = str(round(totalcless / numofsim, 2))
    avg_rest = str(round((7 - totalrest / numofsim), 2))

    return [avg_blue, avg_red, avg_black, avg_green, avg_white, avg_cless, avg_rest, len(deck), deck_identity]

############# NEW SHUFFLE FUNCTION #############

def true_simdeal(deck):
    hand = deck[:7]
    blue = red = black = green = white = cless = 0
    for ii in range(len(hand)):
        if "U" in hand[ii]:
            blue += 1
        if "B" in hand[ii]:
            black += 1
        if "R" in hand[ii]:
            red += 1
        if "G" in hand[ii]:
            green += 1
        if "W" in hand[ii]:
            white += 1
        if "C" in hand[ii]:
            cless += 1

    rest = hand.count("n")

    return blue, red, black, green, white, cless, rest

def make_true_deck(imported_deck):
    deckl, cardcount = get_deck(imported_deck)
    deck = []
    for x, y in zip(deckl, cardcount):
        deck += [x] * y 
    random.shuffle(deck)
    return deck

def shuffle_mana_main(imported_deck):
    deck = []
    deck_identity = []
    img_uris = []
    deckl = make_true_deck(imported_deck)

    for i in range(len(deckl)):
        m_value, card_identity, commander_mana, img_uri = get_values(deckl[i], img=True)

        if commander_mana is True:
            card = "co"
        if m_value is None:
            card = "n"
        else:
            card = "".join(m_value)
        deck += [card]
        deck_identity += [color for color in card_identity if color not in deck_identity]
        img_uris.append(img_uri)

    deck_identity = "".join(deck_identity)
    
    deck[:] = [x if x != "co" else deck_identity for x in deck] # replace all "co" with deck_identity

    blue, red, black, green, white, cless, rest = true_simdeal(deck)


    return [blue, red, black, green, white, cless, 7-rest, len(deck), deck_identity, img_uris[:7]]

if __name__ == "__main__":
    from getdeck import imported_deck
    import time
    start = time.time()
    print(main(1000000, imported_deck))
    print((time.time() - start)*1000, "ms")
    #print(make_true_deck(imported_deck))
    #print(shuffle_mana_main(imported_deck))
