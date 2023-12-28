import re
from time import perf_counter

def parse_input(input):
    instructions = []
    for line in input:
        if "turn" in line:
            if "on" in line:
                inst = "on"
            else:
                inst = "off"
        else:
            inst = "toggle"
        ranges = re.findall(r'(\d+),(\d+) through (\d+),(\d+)', line)
        a, b, c, d = [list(map(int,x)) for x in ranges][0]
        instructions.append((inst, a, b, c, d))
    return instructions

def part1(input):
    lights = [[0 for _ in range(1000)] for _ in range(1000)]
    for inst, a, b, c, d in input:
        for row in range(a, c+1):
            for col in range(b, d+1):
                if inst == "on":
                    lights[row][col] = 1
                elif inst == "off":
                    lights[row][col] = 0
                else:
                    lights[row][col] = 1 - lights[row][col]
    return sum([sum(x) for x in lights])

def part2(input):
    lights = [[0 for _ in range(1000)] for _ in range(1000)]
    for inst, a, b, c, d in input:
        for row in range(a, c+1):
            for col in range(b, d+1):
                if inst == "on":
                    lights[row][col] += 1
                elif inst == "off":
                    lights[row][col] = max(lights[row][col] - 1, 0)
                else:
                    lights[row][col] += 2
    return sum([sum(x) for x in lights])

# No real way to speed this up, maybe use numpy?
def main():
    input = open("input.txt", "r").read().splitlines()
    t1 = perf_counter()
    part1_sum = part1(parse_input(input))
    t2 = perf_counter()
    print(f"Part 1: {part1_sum}")
    print(f'Part 1 took {t2-t1} seconds')
    
    t1 = perf_counter()
    part2_sum = part2(parse_input(input))
    t2 = perf_counter()
    print(f"Part 2: {part2_sum}")
    print(f'Part 2 took {t2-t1} seconds')

if __name__ == "__main__":
    main()