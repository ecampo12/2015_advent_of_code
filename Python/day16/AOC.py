import re
import dataclasses

@dataclasses.dataclass
class Sue:
    id: int
    children: int = None
    cats: int = None
    samoyeds: int = None
    pomeranians: int = None
    akitas: int = None
    vizslas: int = None
    goldfish: int = None
    trees: int = None
    cars: int = None
    perfumes: int = None

def parse_input(input):
    aunt_sue = []
    for i, line in enumerate(input):
        sue = Sue(i + 1)
        for match in re.findall(r"(\w+): (\d+)", line):
            setattr(sue, match[0], int(match[1]))
        aunt_sue.append(sue)
    return aunt_sue

def part1(input):
    aunt_sue = parse_input(input)
    search_cats = {"children": 3, 
                   "cats": 7, 
                   "samoyeds": 2, 
                   "pomeranians": 3, 
                   "akitas": 0, 
                   "vizslas": 0, 
                   "goldfish": 5, 
                   "trees": 3, 
                   "cars": 2, 
                   "perfumes": 1}
    for sue in aunt_sue:
        for key, value in search_cats.items():
            if getattr(sue, key) != None and getattr(sue, key) != value:
                break
        else:
            return sue.id
    return None

def part2(input):
    aunt_sue = parse_input(input)
    search_cats = {"children": 3, 
                   "samoyeds": 2, 
                   "akitas": 0, 
                   "vizslas": 0, 
                   "cars": 2, 
                   "perfumes": 1}
    for sue in aunt_sue:
        for key, value in search_cats.items():
            if getattr(sue, key) != None and getattr(sue, key) != value:
                break
        else:
            if sue.cats != None and sue.cats <= 7:
                continue
            if sue.trees != None and sue.trees <= 3:
                continue
            if sue.pomeranians != None and sue.pomeranians >= 3:
                continue
            if sue.goldfish != None and sue.goldfish >= 5:
                continue
            return sue.id
    return None

def main():
    input = open("input.txt", "r").read().splitlines()
    part1_sum = part1(input)
    print(f"Part 1: {part1_sum}")
    part2_sum = part2(input)
    print(f"Part 2: {part2_sum}")

if __name__ == "__main__":
    main()