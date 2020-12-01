use std::collections::HashSet;
use std::error::Error;

use crate::utils;

// parse numbers from string
// split numbers into two groups: ones <= to half the total, and ones > half the total
// store the difference of number and total in first group, store only the number in second group
// if two numbers exist s.t. they sum to total, the intersection of these two sets will give the pair

pub fn solve_part1() -> Result<(), Box<dyn Error>> {
    let total = 2020;
    let lines = utils::read_file_lines("input/day01")?;

    let half_total = total / 2;
    let mut lower_numbers = HashSet::new();
    let mut higher_numbers = HashSet::new();

    for line in lines {
        let number = line?.parse::<i32>()?;

        if number <= half_total {
            lower_numbers.insert(total - number);
        } else if number > half_total {
            higher_numbers.insert(number);
        }
    }

    for number in lower_numbers.intersection(&higher_numbers) {
        println!("{}", number * (total - number));
    }

    Ok(())
}
