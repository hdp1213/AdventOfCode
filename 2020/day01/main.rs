use std::io::{BufRead, BufReader, Lines};
use std::fs::File;
use std::collections::HashSet;

// parse numbers from string
// split numbers into two groups: ones <= to half the total, and ones > half the total
// store the difference of number and total in first group, store only the number in second group
// if two numbers exist s.t. they sum to total, the intersection of these two sets will give the pair

fn read_file(file_name: &str) -> Lines<BufReader<File>> {
  let reader = BufReader::new(File::open(file_name).expect("Cannot open file"));
  reader.lines()
}

fn main() {
  let total = 2020;
  let lines = read_file("input/day01");

  let half_total = total / 2;
  let mut lower_numbers = HashSet::new();
  let mut higher_numbers = HashSet::new();

  for line in lines {
    if !line.is_ok() {
      continue;
    }

    let parse_result = line.unwrap().parse::<i32>();
    if !parse_result.is_ok() {
      continue;
    }

    let number = parse_result.unwrap();

    if number <= half_total {
      lower_numbers.insert(total - number);
    }
    else if number > half_total {
      higher_numbers.insert(number);
    }
  }

  for number in lower_numbers.intersection(&higher_numbers) {
    println!("{}", number * (total - number));
  }
}
