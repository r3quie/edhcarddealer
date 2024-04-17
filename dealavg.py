import random

def main():
    numofsim = int(input("Enter number of simulations: \n"))
    numlands = 33
    nummanarocks = 8
    deck = ["l"] * numlands + ["r"] * nummanarocks + ["c"] * (100 - numlands - nummanarocks)

    def simdeal():
        hand = random.sample(deck, 7)
        lands = hand.count("l")
        rocks = hand.count("r")
        rest = hand.count("c")
        return lands, rocks, rest

    totallands = totalrocks = totalrest = 0
    for _ in range(numofsim):
        lands, rocks, rest = simdeal()
        totallands += lands
        totalrocks += rocks
        totalrest += rest

    avg_lands = totallands / numofsim
    avg_rocks = totalrocks / numofsim
    avg_mana = (totallands + totalrocks) / numofsim

    print("Avg number of lands:", avg_lands)
    print("Avg number of manarocks:", avg_rocks)
    print("Avg number of mana cards:", avg_mana)

if __name__ == "__main__":
    main()