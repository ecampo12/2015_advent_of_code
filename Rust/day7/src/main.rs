use std::{collections::HashMap, vec};
use regex::Regex;

fn part1(input: &String, value: u16) -> HashMap<&str, u16> {
    let mut wires :HashMap<&str, u16>= HashMap::new();
    let mut lines = input.split("\n").collect::<Vec<&str>>();
    let re = Regex::new(r"AND|OR|LSHIFT|RSHIFT").unwrap();
    loop{
        let mut wait: Vec<&str> = vec![];
        for line in &lines{
            if line.contains("NOT"){
                let l = line.split(" ").collect::<Vec<&str>>();
                let src = l[1];
                let wire = l[3];
                if wires.contains_key(&src){
                    wires.insert(wire, !wires.get(&src).unwrap());
                } else {
                    wait.push(line);
                }
            }
            else if re.is_match(line){
                let l = line.split(" ").collect::<Vec<&str>>();
                let (src1, op, src2, wire) = (l[0], l[1], l[2], l[4]);
                if !src1.parse::<u16>().is_ok() && !wires.contains_key(src1) || !src2.parse::<u16>().is_ok() && !wires.contains_key(&src2) {
                    wait.push(line);
                } else {
                    let s1 = if src1.parse::<u16>().is_ok() {src1.parse::<u16>().unwrap()} else {*wires.get(&src1).unwrap()};
                    let s2 = if src2.parse::<u16>().is_ok() {src2.parse::<u16>().unwrap()} else {*wires.get(&src2).unwrap()};
                    let result = match op {
                        "AND" => s1 & s2,
                        "OR" => s1 | s2,
                        "LSHIFT" => s1 << s2,
                        "RSHIFT" => s1 >> s2,
                        _ => panic!("Invalid OP"),
                    };
                    wires.insert(wire, result);
                }
            }
            else {
                let l = line.split(" ").collect::<Vec<&str>>();
                let (src, wire) = (l[0], l[2]);
                if value != 0 && wire == "b" {
                    wires.insert(wire, value);
                } else if src.parse::<u16>().is_ok() {
                    wires.insert(wire, src.parse::<u16>().unwrap());
                } else if wires.contains_key(&src) {
                    wires.insert(wire, *wires.get(&src).unwrap());
                } else {
                    wait.push(line);
                }
            }
        }
        if wait.len() == 0 && lines.len() == 0 {
            return wires;
        }
        else {
            lines = wait.iter().map(|x| *x).collect();
        }
    }
}

fn main() {
    let input = std::fs::read_to_string("./input.txt").expect("Error reading input file");
    let binding = part1(&input, 0);
    let part1_result = binding.get(&"a").unwrap();
    println!("Part 1: {}", part1_result);
    println!("Part 2: {}", part1(&input, *part1_result).get(&"a").unwrap());
}

#[cfg(test)]
mod tests {
    #[test]
    fn test_part1() {
        let test_input =
            std::fs::read_to_string("./test_input.txt").expect("Error reading test input file");
        let values = std::collections::HashMap::from([
            ("d", 72),
            ("e", 507),
            ("f", 492),
            ("g", 114),
            ("h", 65412),
            ("i", 65079),
            ("x", 123),
            ("y", 456),
        ]);
        assert_eq!(super::part1(&test_input, 0), values);
    }
}
