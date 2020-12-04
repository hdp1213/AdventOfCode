use std::collections::HashSet;
use std::error::Error;

use crate::utils;

pub async fn solve() -> Result<(), Box<dyn Error>> {
    let map_lines = utils::start_day(3, &(|l| l)).await?;

    let num_trees = part1(&map_lines)?;
    println!("number of trees for first pass: {}", num_trees);

    let product = part2(&map_lines)?;
    println!("product of routes: {}", product);

    Ok(())
}

struct TreeMap {
    width: usize,
    height: usize,
    trees: HashSet<(usize, usize)>,
}

impl TreeMap {
    fn new(input: &Vec<String>) -> Result<TreeMap, &'static str> {
        let first_line = input.iter().next().unwrap();

        let height = input.len();
        let width = first_line.len();
        let mut trees = HashSet::new();

        for (i, line) in input.iter().enumerate() {
            for (j, c) in line.char_indices() {
                if c == '#' {
                    trees.insert((i, j));
                }
            }
        }

        Ok(TreeMap {
            width,
            height,
            trees,
        })
    }

    fn get_tree_collisions(&self, slope: (usize, usize)) -> usize {
        let mut num_collisions: usize = 0;
        let mut num_steps: usize = 0;

        let mut i: usize = 0;
        let mut j: usize;

        while i < self.height {
            i = slope.0 * num_steps;
            j = (slope.1 * num_steps) % self.width;

            if self.trees.contains(&(i, j)) {
                num_collisions += 1;
            }

            num_steps += 1;
        }

        return num_collisions;
    }
}

fn part1(input: &Vec<String>) -> Result<usize, Box<dyn Error>> {
    let map = TreeMap::new(input)?;
    let mut num_trees: usize = 0;

    for i in 0..map.height {
        let j = (3 * i) % map.width;

        if map.trees.contains(&(i, j)) {
            num_trees += 1;
        }
    }

    return Ok(num_trees);
}

fn part2(input: &Vec<String>) -> Result<usize, Box<dyn Error>> {
    let map = TreeMap::new(input)?;
    let slopes = vec![(1, 1), (1, 3), (1, 5), (1, 7), (2, 1)];
    let mut product = 1;

    for slope in slopes {
        product *= map.get_tree_collisions(slope);
    }

    Ok(product)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn new_tree_map() {
        let input = vec![".##.", "....", "#.#."]
            .iter()
            .map(|&s| String::from(s))
            .collect();

        let map = TreeMap::new(&input).unwrap();

        assert_eq!(map.height, 3);
        assert_eq!(map.width, 4);
        assert!(map.trees.contains(&(0, 1)));
        assert!(map.trees.contains(&(0, 2)));
        assert!(map.trees.contains(&(2, 0)));
        assert!(map.trees.contains(&(2, 2)));
    }

    #[test]
    fn tree_map_collisions() {
        let input = vec![
            "..##.......",
            "#...#...#..",
            ".#....#..#.",
            "..#.#...#.#",
            ".#...##..#.",
            "..#.##.....",
            ".#.#.#....#",
            ".#........#",
            "#.##...#...",
            "#...##....#",
            ".#..#...#.#",
        ]
        .iter()
        .map(|&s| String::from(s))
        .collect();

        let map = TreeMap::new(&input).unwrap();

        let collisions = map.get_tree_collisions((1, 3));
        assert_eq!(collisions, 7);
    }

    #[test]
    fn part1_tree_path() {
        let input = vec![
            "..##.......",
            "#...#...#..",
            ".#....#..#.",
            "..#.#...#.#",
            ".#...##..#.",
            "..#.##.....",
            ".#.#.#....#",
            ".#........#",
            "#.##...#...",
            "#...##....#",
            ".#..#...#.#",
        ]
        .iter()
        .map(|&s| String::from(s))
        .collect();

        let num_trees = part1(&input).unwrap();

        assert_eq!(num_trees, 7);
    }

    #[test]
    fn part2_product() {
        let input = vec![
            "..##.......",
            "#...#...#..",
            ".#....#..#.",
            "..#.#...#.#",
            ".#...##..#.",
            "..#.##.....",
            ".#.#.#....#",
            ".#........#",
            "#.##...#...",
            "#...##....#",
            ".#..#...#.#",
        ]
        .iter()
        .map(|&s| String::from(s))
        .collect();

        let product = part2(&input).unwrap();

        assert_eq!(product, 336);
    }
}
