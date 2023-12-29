from math import sqrt
from time import perf_counter


def get_divisors(n):
    small_divs = [i for i in range(1, int(sqrt(n))+ 1) if n % i == 0]
    large_divs = [n // d for d in small_divs if n != d * d] # Ignore perfect squares, need for very large numbers
    return small_divs + large_divs

# Solution is slow, but we can speed it up by finding the solution for both parts at the same time.
# Doesn't pass the test, but it's gets the right answer.
def part1(input):
    target = int(input)
    p1, p2 = 0, 0
    i = target 
    while p1 == 0 or p2 == 0:
        i += 1
        divisors = get_divisors(i)
        if p1 == 0 and sum(divisors) * 10 >= target:
            p1 = i
        if p2 == 0 and sum(d for d in divisors if i / d <= 50) * 11 >= target:
            p2 = i
    return p1, p2
               
def main():
    input = open("input.txt", "r").read().splitlines()
    print("Don't worry, this will take a while...\n")
    
    t1 = perf_counter()
    part1_ans, part2_ans = part1(input[0])
    t2 = perf_counter()
    print(f"Part 1: {part1_ans}")
    print(f"Part 1 took {t2 - t1:.4f} seconds\n")
    
    print(f"Part 2: {part2_ans}")

if __name__ == "__main__":
    main()