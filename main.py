import random
#from edhapi import api_handle
from .localbulk import get_values, scryjson
from .getdeck import get_deck

def exampledeck():
    from .getdeck import imported_deck
    imported_deck = """"""" + imported deck + """""""
    return imported_deck

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
    deckl, cardcount = get_deck(imported_deck)

    for i in range(len(deckl)):
        m_value = get_values(deckl[i])
        if m_value == None:
            card = "n"
        else:
            card = "".join(m_value)
        deck += [card] * cardcount[i]

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

    return [avg_blue, avg_red, avg_black, avg_green, avg_white, avg_cless, avg_rest, len(deck)]

if __name__ == "__main__":
    main(int(input("Enter number of simulations: ")))
