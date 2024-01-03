import re

def parse_input(input):
    reindeer_data = []
    for line in input:
        speed, fly_time, rest_time = map(int, re.findall(r"(\d+)", line))
        reindeer_data.append((speed, fly_time, rest_time))        
    return reindeer_data

def part1(input, race_time=1000):
    reideer_data = parse_input(input)
    times = []
    for reindeer in reideer_data:
        speed, fly_time, rest_time = reindeer
        time = 0
        distance = 0
        while time < race_time:
            if time % (fly_time + rest_time) < fly_time:
                distance += speed
            time += 1
        times.append(distance)
    return times

def part2(input, race_time=1000):
    reindeer_data = parse_input(input)
    points = [0] * len(reindeer_data)
    distances = [0]*len(reindeer_data)
    time = 0
    while time < race_time:
        for i, reindeer in enumerate(reindeer_data):
            speed, fly_time, rest_time = reindeer
            distance = distances[i]
            if time % (fly_time + rest_time) < fly_time:
                distance += speed
            distances[i] = distance
        for i, distance in enumerate(distances):
            if distance == max(distances):
                points[i] += 1
        time += 1
    return points

def main():
    input = open("input.txt", "r").read().splitlines()
    part1_sum = max(part1(input, 2503))
    print(f"Part 1: {part1_sum}")
    part2_sum = max(part2(input, 2503))
    print(f"Part 2: {part2_sum}")

if __name__ == "__main__":
    main()