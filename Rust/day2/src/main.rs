fn part1(input: &String) -> i32 {
    input.split("\n").map(|line| {
        let dims: Vec<i32> = line.split("x").map(|x| x.parse::<i32>().unwrap()).collect();
        let sides = vec![dims[0] * dims[1], dims[1] * dims[2], dims[2] * dims[0]];
        let small = sides.iter().min().unwrap();
        sides.iter().map(|x| x * 2).sum::<i32>() + small
    }).sum()
}

fn part2(input: &String) -> i32 {
    input.split("\n").map(|line| {
        let mut dims: Vec<i32> = line.split("x").map(|x| x.parse::<i32>().unwrap()).collect();
        dims.sort();
        dims.iter().fold(1, |acc, x| acc * x) + 2 * dims[0] + 2 * dims[1]
    }).sum()
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
        assert_eq!(super::part1(&String::from("2x3x4")), 58);
        assert_eq!(super::part1(&String::from("1x1x10")), 43);
    }

    #[test]
    fn test_part2() {
        assert_eq!(super::part2(&String::from("2x3x4")), 34);
        assert_eq!(super::part2(&String::from("1x1x10")), 14);
    }
}
