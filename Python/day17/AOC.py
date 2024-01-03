from itertools import combinations

def part1(input, target=25):
    containers = []
    for line in input:
        containers.append(int(line))
        
    combos = []
    for i in range(len(containers)):
        for combo in combinations(containers, i):
            if sum(combo) == target:
                combos.append(combo)
    return combos

def part2(input, target=25):
    combos = part1(input, target)
    min_len = min([len(combo) for combo in combos])
    min_len_combos = [combo for combo in combos if len(combo) == min_len]
    return len(min_len_combos)

def main():
    input = open("input.txt", "r").read().splitlines()
    part1_sum = len(part1(input, 150))
    print(f"Part 1: {part1_sum}")
    part2_sum = part2(input, 150)
    print(f"Part 2: {part2_sum}")

if __name__ == "__main__":
    main()