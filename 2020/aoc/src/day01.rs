use std::collections::{HashMap, HashSet};
use std::error::Error;

use crate::utils;

// parse numbers from string
// split numbers into two groups: ones <= to half the total, and ones > half the total
// store the difference of number and total in first group, store only the number in second group
// if two numbers exist s.t. they sum to total, the intersection of these two sets will give the pair

pub async fn solve() -> Result<(), Box<dyn Error>> {
    let day: i32 = 1;
    utils::load_input(day).await?;
    let input_file = utils::input_dir().join(format!("day{:02}", day));

    println!("loading input for day 01...");
    let lines = utils::read_file_lines(input_file)?;
    let total = 2020;

    let mut numbers = Vec::new();

    for line in lines {
        numbers.push(line?.parse::<i32>()?);
    }

    part1(&numbers, total)?;
    part2(&numbers, total)?;

    Ok(())
}

fn part1(input: &Vec<i32>, total: i32) -> Result<(), Box<dyn Error>> {
    let half_total = total / 2;

    let mut lower_numbers: HashSet<i32> = HashSet::new();
    let mut higher_numbers: HashSet<i32> = HashSet::new();

    for &number in input {
        if number <= half_total {
            lower_numbers.insert(total - number);
        } else if number > half_total {
            higher_numbers.insert(number);
        }
    }

    for &number in lower_numbers.intersection(&higher_numbers) {
        let (a, b) = (number, total - number);
        println!("{} * {} = {}", a, b, a * b);
    }

    Ok(())
}

fn part2(input: &Vec<i32>, total: i32) -> Result<(), Box<dyn Error>> {
    let half_total = total / 2;

    let mut lower_numbers: HashSet<i32> = HashSet::new();
    let mut higher_numbers: HashSet<i32> = HashSet::new();
    let mut tuples: HashMap<i32, (i32, i32, i32)> = HashMap::new();

    for i in 0..input.len() {
        let a = input[i];

        // Make sure to add all numbers individually to the HashSets
        if a <= half_total {
            lower_numbers.insert(total - a);
        } else {
            higher_numbers.insert(a);
        }

        for j in i..input.len() {
            let b = input[j];

            // Also add the pair's sum
            let sub_sum = a + b;
            tuples.insert(sub_sum, (a, b, total - sub_sum));

            if sub_sum <= half_total {
                lower_numbers.insert(total - sub_sum);
            } else if sub_sum < total {
                higher_numbers.insert(sub_sum);
            }
        }
    }

    for &number in lower_numbers.intersection(&higher_numbers) {
        if tuples.contains_key(&number) {
            let (a, b, c) = tuples[&number];
            println!("{} * {} * {} = {}", a, b, c, a * b * c);
        }
    }

    Ok(())
}
