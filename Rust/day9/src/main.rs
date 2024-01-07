use itertools::Itertools;
use std::collections::HashMap;

// Yes, for part 1 I could implement Dijkstra's algorithm, but I know that I need to find the longetst path for part 2, so I'm just going to implement a brute force solution for both parts.
fn all_path_lengths(input: &String) -> Vec<usize> {
    let mut paths = Vec::<usize>::new();
    let mut nodes = HashMap::<String, HashMap<String, usize>>::new();
    input.lines().for_each(|line| {
        let split = line.split(" ").collect::<Vec<&str>>();
        let from = split[0].to_string();
        let to = split[2].to_string();
        let distance = split[4].parse::<usize>().unwrap();
        nodes
            .entry(from.clone())
            .or_insert(HashMap::new())
            .insert(to.clone(), distance);
        nodes
            .entry(to.clone())
            .or_insert(HashMap::new())
            .insert(from, distance);
    });

    let node_labels = nodes.keys().map(|k| k.clone()).collect::<Vec<String>>();
    let permutations = node_labels.iter().permutations(node_labels.len());
    for permutation in permutations {
        let mut path_length = 0;
        for i in 0..permutation.len() - 1 {
            let from = &permutation[i];
            let to = &permutation[i + 1];
            path_length += nodes.get(*from).unwrap().get(*to).unwrap();
        }
        paths.push(path_length);
    }
    return paths;
}

fn main() {
    let input = std::fs::read_to_string("./input.txt").expect("Error reading input file");
    let paths = all_path_lengths(&input);
    println!("Part 1: {}", paths.iter().min().unwrap());
    println!("Part 2: {}", paths.iter().max().unwrap());
}

#[cfg(test)]
mod tests {
    #[test]
    fn test_part1() {
        let test_input =
            std::fs::read_to_string("./test_input.txt").expect("Error reading test input file");
        assert_eq!(
            super::all_path_lengths(&test_input)
                .into_iter()
                .min()
                .unwrap(),
            605
        );
    }

    #[test]
    fn test_part2() {
        let test_input =
            std::fs::read_to_string("./test_input.txt").expect("Error reading test input file");
        assert_eq!(
            super::all_path_lengths(&test_input)
                .into_iter()
                .max()
                .unwrap(),
            982
        );
    }
}
