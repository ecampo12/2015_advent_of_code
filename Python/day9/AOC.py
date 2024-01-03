import networkx as nx
import itertools

def parse_input(input):
    graph = nx.Graph()
    for line in input:
        start , _, to, _, distance = line.split()
        graph.add_edge(start, to, weight=int(distance))
    return graph

def path_lengths(input):
    graph = parse_input(input)
    distances = []
    for path in itertools.permutations(graph.nodes()):
        distance = 0
        for i in range(len(path) - 1):
            distance += graph.get_edge_data(path[i], path[i+1])['weight']
        distances.append(distance)
    return distances
    
def part1(input):
    return min(path_lengths(input))

def part2(input):
    return max(path_lengths(input))

def main():
    input = open("input.txt", "r").read().splitlines()
    part1_sum = part1(input)
    part2_sum = part2(input)
    print(f"Part 1: {part1_sum}")
    print(f"Part 2: {part2_sum}")

if __name__ == "__main__":
    main()