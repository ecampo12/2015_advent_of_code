import re
import networkx as nx
from itertools import permutations

def parse_input(input):
    graph = nx.DiGraph()
    for line in input:
        x = line.split()
        p1 = x[0]
        p2 = x[-1][:-1] # remove the period
        weight = int(x[3])
        if x[2] == "lose":
            weight = -weight
        graph.add_edge(p1, p2, weight=weight)
    return graph

def part1(input, me=False):
    graph = parse_input(input)
    if me:
        graph.add_node("me")
        for node in graph.nodes():
            if node != "me":
                graph.add_edge("me", node, weight=0)
                graph.add_edge(node, "me", weight=0)
                
    happiness = []
    for path in permutations(graph.nodes()):
        happy = 0
        for i in range(len(path) - 1):
            happy += graph.get_edge_data(path[i], path[(i+1)])['weight']
            happy += graph.get_edge_data(path[(i+1)], path[i])['weight']
        happy += graph.get_edge_data(path[0], path[-1])['weight']
        happy += graph.get_edge_data(path[-1], path[0])['weight']
        happiness.append(happy)
    return max(happiness)

def part2(input):
    return part1(input, me=True)

def main():
    input = open("input.txt", "r").read().splitlines()
    part1_sum = part1(input)
    print(f"Part 1: {part1_sum}")
    part2_sum = part2(input)
    print(f"Part 2: {part2_sum}")

if __name__ == "__main__":
    main()