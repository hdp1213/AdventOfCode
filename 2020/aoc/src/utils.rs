use std::error::Error;
use std::fs::File;
use std::io::{BufRead, BufReader, Lines};

pub fn read_file_lines(file_name: &str) -> Result<Lines<BufReader<File>>, Box<dyn Error>> {
    let reader = BufReader::new(File::open(file_name)?);
    Ok(reader.lines())
}
