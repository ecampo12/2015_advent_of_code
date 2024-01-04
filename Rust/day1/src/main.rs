fn part1(input: &String) -> i32 {
    let mut floor = 0;
    for c in input.chars() {
        match c {
            '(' => floor += 1,
            ')' => floor -= 1,
            _ => (),
        }
    }
    return floor;
}

fn part2(input: &String) -> i32 {
    let mut floor = 0;
    for (i, c) in input.chars().enumerate() {
        match c {
            '(' => floor += 1,
            ')' => floor -= 1,
            _ => (),
        }
        if floor == -1 {
            return (i + 1) as i32;
        }
    }
    return -1;
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
        assert_eq!(super::part1(&String::from("(())")), 0);
        assert_eq!(super::part1(&String::from("(()(()(")), 3);
        assert_eq!(super::part1(&String::from("))(((((")), 3);
        assert_eq!(super::part1(&String::from("())")), -1);
        assert_eq!(super::part1(&String::from(")())())")), -3);
    }

    #[test]
    fn test_part2() {
        assert_eq!(super::part2(&String::from(")")), 1);
        assert_eq!(super::part2(&String::from("()())")), 5);
    }
}
