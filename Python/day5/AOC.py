import re

def part1(input):
    nice = 0
    for line in input:
        if len(re.findall(r'(.)\1{1,}', line)) < 1:
            continue
        if len(re.findall(r'[a|e|i|o|u]', line)) < 3:
            continue
        if len(re.findall(r'ab|cd|pq|xy', line)) > 0:
            continue
        nice += 1
    return nice

def part2(input):
    nice = 0
    for line in input:
        if len(re.findall(r'(.{2}).*?(\1)', line)) < 1:
            continue
        if len(re.findall(r'(.).\1', line)) < 1:
            continue
        nice += 1
    return nice

def main():
    input = open("input.txt", "r").read().splitlines()
    part1_sum = part1(input)
    part2_sum = part2(input)
    print(f"Part 1: {part1_sum}")
    print(f"Part 2: {part2_sum}")

if __name__ == "__main__":
    main()