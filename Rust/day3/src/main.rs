use std::collections::{HashSet, HashMap};
use bevy_math::IVec2;

fn part1(input: &String) -> i32 {
    let directions: HashMap<char, IVec2> =  HashMap::from([
        ('>', IVec2::new(1, 0)),
        ('<', IVec2::new(-1, 0)),
        ('^', IVec2::new(0, 1)),
        ('v', IVec2::new(0, -1)),
    ]);
    let mut start = IVec2::new(0, 0);
    let mut visited = HashSet::new();
    visited.insert(start);
    for c in input.chars() {
        let dir = directions.get(&c).unwrap();
        start += *dir;
        visited.insert(start);
    }
    return visited.len() as i32
}

fn part2(input: &String) -> i32 {
    let directions: HashMap<char, IVec2> =  HashMap::from([
        ('>', IVec2::new(1, 0)),
        ('<', IVec2::new(-1, 0)),
        ('^', IVec2::new(0, 1)),
        ('v', IVec2::new(0, -1)),
    ]);

    let mut santa = IVec2::new(0, 0);
    let mut robo = IVec2::new(0, 0);

    let mut santa_visited = HashSet::new();
    let mut robo_visited = HashSet::new();

    santa_visited.insert(santa);
    robo_visited.insert(robo);

    for (i, c) in input.chars().enumerate() {
        let dir = directions.get(&c).unwrap();
        if i % 2 == 0 {
            santa += *dir;
            santa_visited.insert(santa);
        } else {
            robo += *dir;
            robo_visited.insert(robo);
        }
    }
    return santa_visited.len() as i32 + robo_visited.len() as i32 - santa_visited.intersection(&robo_visited).count() as i32
}

fn main() {
    let input = std::fs::read_to_string("./input.txt").expect("Error reading input file");
    println!("Part 1: {}", part1(&input));
    println!("Part 2: {}", part2(&input));
}

#[cfg(test)]
mod tests {
    #[test]
    fn test_part1() {
        assert_eq!(super::part1(&String::from(">")), 2);
        assert_eq!(super::part1(&String::from("^>v<")), 4);
        assert_eq!(super::part1(&String::from("^v^v^v^v^v")), 2);
    }

    #[test]
    fn test_part2() {
        assert_eq!(super::part2(&String::from("^v")), 3);
        assert_eq!(super::part2(&String::from("^>v<")), 3);
        assert_eq!(super::part2(&String::from("^v^v^v^v^v")), 11);
    }
}