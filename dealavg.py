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

def mainthree():
    numofsim = int(input("Enter number of simulations: \n"))
    
    numulands = int(input("Enter number of purely blue lands: "))
    numblands = int(input("Enter number of purely black lands: "))
    numbulands = int(input("Enter number of purely blue+black lands: "))
    numrlands = int(input("Enter number of purely red lands: "))
    numrulands = int(input("Enter number of purely red+blue lands: "))
    numrblands = int(input("Enter number of purely red+black lands: "))
    numrbulands = int(input("Enter number of any/or all color lands: "))

    numurocks = int(input("Enter number of purely blue rocks: "))
    numbrocks = int(input("Enter number of purely black rocks: "))
    numburocks = int(input("Enter number of purely blue+black rocks: "))
    numrrocks = int(input("Enter number of purely red rocks: "))
    numrurocks = int(input("Enter number of purely red+blue rocks: "))
    numrbrocks = int(input("Enter number of purely red+black rocks: "))
    deck = ["lu"] * numulands + ["lb"] * numblands + ["lbu"] * numbulands + ["lr"] * numrlands + ["lru"] * numrulands + ["lrb"] * numrblands + ["lrbu"] * numrbulands + ["Ru"] * numurocks + ["Rb"] * numbrocks + ["Rbu"] * numburocks + ["Rr"] * numrrocks + ["Rru"] * numrurocks + ["Rrb"] * numrbrocks
    deck += ["c"] * (100 - len(deck))

    def simdeal():
        hand = random.sample(deck, 7)
        print(hand)
        ulands = hand.count("lu")
        blands = hand.count("lb")
        rlands = hand.count("lr")
        bulands = hand.count("lbu")
        rulands = hand.count("lru")
        rblands = hand.count("lrb")
        rbulands = hand.count("lrbu")
        
        urocks = hand.count("Ru")
        brocks = hand.count("Rb")
        rrocks = hand.count("Rr")
        burocks = hand.count("Rbu")
        rurocks = hand.count("Rru")
        rbrocks = hand.count("Rrb")

        blue = red = black = 0
        for i in range(len(hand)):
            if "u" in hand[i]:
                blue += 1
            elif "b" in hand[i]:
                black += 1
            elif "r" in hand[i]:
                red += 1  
        
        rest = hand.count("c")
        # return ulands, blands, rlands, bulands, rulands, rblands, rbulands, urocks, brocks, rrocks, burocks, rurocks, rbrocks, rest, blue, red, black
        return blue, red, black
################################################################################## UNFINISHED UNDER THIS LINE ##################################################################################
    totalblue = totalred = totalblack = 0
    for _ in range(numofsim):
        blue, red, black = simdeal()
        totalblue += blue
        totalred += red
        totalblack += black

    avg_blue = totalblue / numofsim
    avg_red = totalred / numofsim
    avg_black = totalblack / numofsim

    print("Avg number of blue mana cards:", avg_blue)
    print("Avg number of red mana cards:", avg_red)
    print("Avg number of black mana cards:", avg_black)

if __name__ == "__main__":
    if input("(m)anavalue or (c)olor mana value?\n") == "c":
        mainthree()
    else:
        main()