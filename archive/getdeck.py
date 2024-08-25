imported_deck = open("cache/deck.txt").read()

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
