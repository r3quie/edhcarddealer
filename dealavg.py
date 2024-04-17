import random

numofsim = input("enter number of simulations: \n")
numofsim = int(numofsim)
numlands = 33
nummanarocks = 8
deck = []

while len(deck) < numlands:
    deck.append("l")

while len(deck) - numlands < nummanarocks:
    deck.append("r")

while len(deck) < 100:
    deck.append("c")

print(deck)
print(len(deck))

def simdeal():
    hand = []
    i = 0
    
    while i < 7:
        hand.append(random.choice(deck))
        i += 1
    # print(hand)
    lands = 0
    rocks = 0
    rest = 0

    while len(hand) > 0:
        j = hand.pop(0)
        if j == "l":
            lands += 1
        elif j == "r":
            rocks += 1
        else:
            rest += 1

    return [lands, rocks, rest]

totallands = 0
totalrocks = 0
totalrest = 0
k = 0
while k < numofsim:
    sim = simdeal()
    totallands += int(sim.pop(0))
    totalrocks += int(sim.pop(0))
    totalrest += int(sim.pop(0))
    k += 1
print(k)
print("Avg number of lands: " + str(totallands/numofsim))
print("Avg number of manarocks: " + str(totalrocks/numofsim))
mana = totallands + totalrocks
print("Avg number of mana cards: " + str(mana/numofsim))