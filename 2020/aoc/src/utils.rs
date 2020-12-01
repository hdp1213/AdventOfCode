use std::io::{BufRead, BufReader, Lines};
use std::fs::File;

pub fn read_file_lines(file_name: &str) -> Lines<BufReader<File>> {
  let reader = BufReader::new(File::open(file_name).expect("Cannot open file"));
  reader.lines()
}
