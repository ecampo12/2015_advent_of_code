use regex::Regex;

fn part1(input: &String) -> usize {
    let hex_num = Regex::new(r"\\x[0-9a-f]{2}").unwrap();

    input
        .split("\n")
        .map(|line| {
            let mut line_mod = line.replace(r#"\""#, "X").replace(r#"\\"#, "X");
            line_mod = hex_num.replace_all(&line_mod, "X").to_string();
            let l1 = line.len();
            let l2 = line_mod.len() - 2;
            l1 - l2
        })
        .into_iter()
        .sum()
}

fn part2(input: &String) -> usize {
    input.split("\n").fold(0, |acc, f| {
        acc + 2 + f.matches(r#"""#).count() + f.matches(r#"\"#).count()
    })
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
        let test_input =
            std::fs::read_to_string("./test_input.txt").expect("Error reading test input file");
        assert_eq!(super::part1(&test_input), 12);
    }

    #[test]
    fn test_part2() {
        let test_input =
            std::fs::read_to_string("./test_input.txt").expect("Error reading test input file");
        assert_eq!(super::part2(&test_input), 19);
    }
}
