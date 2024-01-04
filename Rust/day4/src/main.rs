use md5;
use std::time::Instant;

fn part1(input: &String) -> i32{
    let mut i = 0;
    loop {
        let key = format!("{}{}", input, i);
        let result_str = format!("{:x}", md5::compute(key));
        if result_str.starts_with("00000") {
            return i;
        }
        i += 1;
    }
}

fn part2(input: &String) -> i32{
    let mut i = 0;
    loop {
        let key = format!("{}{}", input, i);
        let result_str = format!("{:x}", md5::compute(key));
        if result_str.starts_with("000000") {
            return i;
        }
        i += 1;
    }
}

fn main() {
    let input = std::fs::read_to_string("./input.txt").expect("Error reading input file");
    let mut now = Instant::now();
    println!("Part 1: {}", part1(&input));
    println!("Time: {} millisecs", now.elapsed().as_millis());

    now = Instant::now();
    println!("Part 2: {}", part2(&input));
    println!("Time: {} secs", now.elapsed().as_secs());
}

#[cfg(test)]
mod tests {
    #[test]
    fn test_part1() {
        assert_eq!(super::part1(&String::from("abcdef")), 609043);
        assert_eq!(super::part1(&String::from("pqrstuv")), 1048970);
    }
}