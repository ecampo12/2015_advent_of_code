def look_and_say(input):
    count = 0
    look_at = input[0]
    output = ""
    for char in input:
        if look_at == char:
            count += 1
        else:
            output += str(count) + look_at
            look_at = char
            count = 1
    output += str(count) + look_at
    return output

def part1(input):
    for _ in range(40):
        input = look_and_say(input)
    return len(input)

def part2(input):
    for _ in range(50):
        input = look_and_say(input)
    return len(input)

def main():
    input = open("input.txt", "r").read().splitlines()
    part1_sum = part1(input[0])
    part2_sum = part2(input[0])
    print(f"Part 1: {part1_sum}")
    print(f"Part 2: {part2_sum}")

if __name__ == "__main__":
    main()