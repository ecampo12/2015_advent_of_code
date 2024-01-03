import re
from itertools import product
from time import perf_counter

def parse_input(input):
    ingredients = []
    for line in input:
        name = line.split(": ")[0]
        cap, dur, flavor, texture, calories = map(int, re.findall(r"(-?\d+)", line))
        ingredients.append((name, cap, dur, flavor, texture, calories))
    return ingredients

# Slow, brute force solution. I could probably do this with linear algebra, but I'm not sure how.
def part1(input, cal_count=False):
    ingredients = parse_input(input)
    total_picks = 100
    prods = []
    for counts in product(range(total_picks + 1), repeat=len(ingredients)):
        if sum(counts) == total_picks:
            totals = [0] * 5
            for ingredient, count in zip(ingredients, counts):
                _, cap, dur, flavor, texture, calories = ingredient
                totals[0] += cap * count
                totals[1] += dur * count
                totals[2] += flavor * count
                totals[3] += texture * count
                totals[4] += calories * count
            if cal_count and totals[4] != 500:
                continue
            prod = 1
            for i in range(len(totals) - 1):
                prod *= max(0, totals[i])
            prods.append(prod)
    return max(prods)

def part2(input):
    return part1(input, True)

def main():
    input = open("input.txt", "r").read().splitlines()
    t1 = perf_counter()
    part1_sum = part1(input)
    t2 = perf_counter()
    print(f"Part 1: {part1_sum}")
    print(f'Part 1 took {t2 - t1:.4f} seconds')
    
    t1 = perf_counter()
    part2_sum = part2(input)
    t1 = perf_counter()
    print(f"Part 2: {part2_sum}")
    print(f'Part 2 took {t1 - t2:.4f} seconds')

if __name__ == "__main__":
    main()