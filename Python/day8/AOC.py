def part1(input):
    return sum([len(line) - len(eval(line)) for line in input])

def part2(input):
    # orig_sum = sum([len(line) for line in input])
    new_sum = sum([2 + line.count('"') + line.count('\\') for line in input])
    return new_sum #- orig_sum

def main():
    input = open("input.txt", "r").read().splitlines()
    part1_sum = part1(input)
    part2_sum = part2(input)
    print(f"Part 1: {part1_sum}")
    print(f"Part 2: {part2_sum}")

if __name__ == "__main__":
    main()