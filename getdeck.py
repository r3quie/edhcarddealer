#imported_deck = open("cache/deck.txt")
imported_deck = """1 Aarakocra Sneak\r
1 Access Tunnel\r
1 Aether Tunnel\r
1 Ancestral Vision\r
1 Aqueous Form\r
1 As Foretold\r
1 Baleful Mastery\r
1 Blackblade Reforged\r
1 Blasphemous Act\r
1 Bojuka Bog\r
1 Braids, Conjurer Adept\r
1 Cascade Bluffs\r
1 Caves of Chaos Adventurer\r
1 Chaos Warp\r
1 Command Tower\r
1 Counterspell\r
1 Court of Ire\r
1 Creeping Bloodsucker\r
1 Crumbling Necropolis\r
1 Darkwater Catacombs\r
1 Defabricate\r
1 Descent into Avernus\r
1 Dimir Signet\r
1 Dragonmaster Outcast\r
1 Dragonskull Summit\r
1 Drowned Catacomb\r
1 Endless Evil\r
1 Exotic Orchard\r
1 Feed the Swarm\r
1 Feywild Caretaker\r
1 Gate to the Aether\r
1 Indulgent Tormentor\r
1 Infernal Grasp\r
1 Inspiring Refrain\r
6 Island\r
1 Izzet Signet\r
1 Kumena's Awakening\r
1 Lizard Blades\r
1 Mechanized Production\r
1 Midnight Clock\r
4 Mountainv
1 Negate\r
1 Nightscape Familiar\r
1 Palace Siege\r
1 Passageway Seer\r
1 Path of Ancestry\r
1 Phyrexian Arena\r
1 Plargg and Nassari\r
1 Profane Tutor\r
1 Protection Racket\r
1 Rakdos Signet\r
1 Ravenloft Adventurer\r
1 Reliquary Tower\r
1 Replicating Ring\r
1 Rilsa Rael, Kingpin\r
1 Ring of Evos Isle\r
1 Ring of Valkas\r
1 Ring of Xathrid\r
1 Rogue's Passage\r
1 Rousing Refrain\r
1 Shivan Reef\r
1 Skyline Despot\r
1 Smoldering Marsh\r
1 Sphinx of the Second Sun\r
1 Star Whale\r
1 Stirring Bard\r
1 Stolen Strategy\r
1 Sulfur Falls\r
1 Sulfurous Springs\r
1 Sunken Hollow\r
2 Swamp\r
1 Swiftfoot Boots\r
1 Sword Coast Sailor\r
1 Talisman of Creativity\r
1 Talisman of Dominance\r
1 Talisman of Indulgence\r
1 Tavern Brawler\r
1 Temple of Deceit\r
1 Temple of Epiphany\r
1 Temple of Malice\r
1 Thassa, God of the Sea\r
1 The Ninth Doctor\r
1 The Tenth Doctor\r
1 Thopter Spy Network\r
1 Tomb of Horrors Adventurer\r
1 Twilight Prophet\r
1 Underground River\r
1 Vandalblast\r
1 Wheel of Fate\r
1 Whispersilk Cloak\r
1 Obeka, Splitter of Seconds"""


## THIS CODE IS SHIT
def get_deck(imported_deck):
    lines = imported_deck.split("\r\n")  # Split input into lines
    modified_lines = []  # To store modified lines
    removed_numbers = []  # To store removed numbers

    for line in lines:
        space_index = line.find(" ")

        if space_index == -1:
            modified_lines.append(line)
            continue

        removed_digits = ""
        for char in line[:space_index]:
            if char.isdigit():
                removed_digits += char

        removed_numbers.append(int(removed_digits))
        modified_lines.append(line[space_index+1:])

    return modified_lines, removed_numbers

if __name__ == "__main__":
    print(get_deck(imported_deck))
