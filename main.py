import random
#from edhapi import api_handle
from localbulk import get_values, scryjson
from getdeck import get_deck, imported_deck

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

def main(numofsim):
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

    avg_blue = totalblue / numofsim
    avg_red = totalred / numofsim
    avg_black = totalblack / numofsim
    avg_green = totalgreen / numofsim
    avg_white = totalwhite / numofsim
    avg_cless = totalcless / numofsim
    avg_rest = totalrest / numofsim

    print("Avg number of blue mana cards:", avg_blue)
    print("Avg number of red mana cards:", avg_red)
    print("Avg number of black mana cards:", avg_black)
    print("Avg number of green mana cards:", avg_green)
    print("Avg number of white mana cards:", avg_white)
    print("Avg number of colorless mana cards:", avg_cless)
    print("Avg number of mana generating cards cards:", 7-avg_rest)

if __name__ == "__main__":
    main(int(input("Enter number of simulations: ")))