def part1(input):
    floor = 0
    for char in input:
        if char == '(':
            floor += 1
        elif char == ')':
            floor -= 1
    return floor

def part2(input):
    index = 0
    floor = 0
    for char in input:
        index += 1
        if char == '(':
            floor += 1
        elif char == ')':
            floor -= 1
        if floor < 0:
            return index
    return 0

def main():
    input = open("input.txt", "r").read().splitlines()
    part1_sum = part1(input[0])
    part2_sum = part2(input[0])
    print(f"Part 1: {part1_sum}")
    print(f"Part 2: {part2_sum}")

if __name__ == "__main__":
    main()