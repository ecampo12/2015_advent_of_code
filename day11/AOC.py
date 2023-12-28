import re
from time import perf_counter

# Password requirements:
# 1. Must include one increasing straight of at least three letters
# 2. Must not contain the letters i, o, or l
# 3. Must contain at least two different, non-overlapping pairs of letters
def password_validity(password):
    # This regex is very long and ugly, but it works
    if re.search(r"abc|bcd|cde|def|efg|fgh|ghi|hij|ijk|jkl|klm|lmn|mno|nop|opq|pqr|qrs|rst|stu|tuv|uvw|vwx|wxy|xyz", password) == None:
        return False
    if re.search(r"[iol]", password) != None:
        return False
    if re.search(r"(.)\1.*(.)\2", password) == None:
        return False
    return True

# Slow, but gets the job done. I takes ~ 9-10 seconds to run the tests
# Can be optimized by skipping paasswords that contain i, o, or l
# Implemented it, speed up about 2 seconds on the tests, no change real input
def next_password(password):
    p = list(password)
    if p[-1] == "z":
        p[-1] = "a"
        index = -2
        while p[index] == "z":
            p[index] = "a"
            index -= 1
        p[index] = chr(ord(p[index]) + 1)
    else:
        p[-1] = chr(ord(p[-1]) + 1)
    password = "".join(p)
    return password

def part1(input):
    password = next_password(input)
    while not password_validity(password):
        password = next_password(password)
    return password

def part2(input):
    return part1(input)

def main():
    input = open("input.txt", "r").read().splitlines()
    part1_sum = part1(input[0])
    print(f"Part 1: {part1_sum}")
    
    t1 = perf_counter()
    part2_sum = part2(part1_sum)
    t2 = perf_counter()
    print(f"Part 2: {part2_sum}")
    print(f"It part 2 took {t2 - t1} seconds") # ~ 1 second, tests take ~ 9-10 seconds

if __name__ == "__main__":
    main()