import re
import json

def part1(input):
    return sum(map(int, re.findall(r"-?\d+", input)))

def get_sum(obj):
    if isinstance(obj, int):
        return obj
    elif isinstance(obj, list):
        return sum([get_sum(item) for item in obj])
    elif isinstance(obj, dict):
        if "red" in obj.values():
            return 0
        else:
            return sum([get_sum(item) for item in obj.values()])
    return 0

def part2(input):
    obj = json.loads(input)
    return get_sum(obj)

def main():
    input = open("input.txt", "r").read().splitlines()
    part1_sum = part1(input[0])
    print(f"Part 1: {part1_sum}")
    part2_sum = part2(input[0])
    print(f"Part 2: {part2_sum}")

if __name__ == "__main__":
    main()