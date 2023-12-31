from itertools import combinations
from functools import reduce
from time import perf_counter

# Assumptions:
#   1. the sum of the weights is divisible by 3
#   2. It's possible to split the weights into 3 equally weighted groups
#   3. There are no weights that make up exactly 1/3 of the total weight
#   4. There at leats 2 weights that make up 1/3 of the total weight
#   5. No more than half of the weights make up 1/3 of the total weight
def part1(input):
    input = [int(x) for x in input]
    g1 = []
    g2 = []
    for i in range(2, len(input)//2):
        for comb in combinations(input, i):
            if sum(comb) == sum(input)//3:
                g1.append(comb)
                break
    return min([reduce(lambda x, y: x*y, x) for x in g1])

def part2(input):
    input = [int(x) for x in input]
    g1 = []
    for i in range(2, len(input)//2):
        for comb in combinations(input, i):
            if sum(comb) == sum(input)//4:
                g1.append(comb)
                break
    for g in g1:
        print(f'Group: {g} QE: {reduce(lambda x, y: x*y, g)}')
    return min([reduce(lambda x, y: x*y, x) for x in g1])

def main():
    input = open("input.txt", "r").read().splitlines()
    t1 = perf_counter()
    part1_ans = part1(input)
    t2 = perf_counter()
    print(f"Part 1: {part1_ans}")
    print(f'Part 1 time: {t2-t1:.4f}\n')
    
    t1 = perf_counter()
    part2_ans = part2(input)
    t2 = perf_counter()
    print(f"Part 2: {part2_ans}")
    print(f'Part 2 time: {t2-t1:.4f}')

if __name__ == "__main__":
    main()