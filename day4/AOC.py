from hashlib import md5

def part1(input, start=5):
    count = 0
    while True:
        count += 1
        key = input + str(count)
        hash = md5(key.encode()).hexdigest()
        if hash.startswith("0"*start):
            return count

def part2(input):
    return part1(input, 6)

def main():
    input = open("input.txt", "r").read().splitlines()
    part1_sum = part1(input[0])
    part2_sum = part2(input[0])
    print(f"Part 1: {part1_sum}")
    print(f"Part 2: {part2_sum}")

if __name__ == "__main__":
    main()