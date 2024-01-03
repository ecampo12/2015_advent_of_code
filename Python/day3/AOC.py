directions = {
        '^': (0, 1),
        '>': (1, 0),
        'v': (0, -1),
        '<': (-1, 0)
    }

def part1(input):
    house = (0, 0)
    visited = set()
    
    visited.add(house)
    for dir in input:
        house = tuple(map(sum, zip(house, directions[dir])))
        visited.add(house)
    return len(visited)

def part2(input):
    house, r_house = (0, 0), (0, 0)
    santa, robo_santa  = set(), set()
    
    santa.add(house)
    robo_santa.add(r_house)
    for i, dir in enumerate(input):
        if i % 2 == 0:
            house = tuple(map(sum, zip(house, directions[dir])))
            santa.add(house)
        else:
            r_house = tuple(map(sum, zip(r_house, directions[dir])))
            robo_santa.add(r_house)
    return len(santa) + len(robo_santa) - len(santa.intersection(robo_santa))


def main():
    input = open("input.txt", "r").read().splitlines()
    part1_sum = part1(input[0])
    part2_sum = part2(input[0])
    print(f"Part 1: {part1_sum}")
    print(f"Part 2: {part2_sum}")

if __name__ == "__main__":
    main()