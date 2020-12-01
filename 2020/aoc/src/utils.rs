use std::error::Error;
use std::fs;
use std::io::{BufRead, BufReader, Lines};
use std::path::{Path, PathBuf};

use reqwest::{header, Client};

pub fn input_dir() -> &'static Path {
    Path::new("input/")
}

pub fn read_file_lines(file_name: PathBuf) -> Result<Lines<BufReader<fs::File>>, Box<dyn Error>> {
    let reader = BufReader::new(fs::File::open(file_name)?);
    Ok(reader.lines())
}

pub async fn load_input(day: i32) -> Result<(), Box<dyn Error>> {
    let output_file = input_dir().join(format!("day{:02}", day));

    if output_file.exists() {
        return Ok(());
    }

    if !input_dir().exists() {
        fs::create_dir(input_dir())?;
    }

    let secrets = load_secrets()?;
    let mut headers = header::HeaderMap::new();
    let cookie = format!("session={}", secrets.session_token);
    let mut cookie_header = header::HeaderValue::from_str(cookie.as_str())?;
    cookie_header.set_sensitive(true);
    headers.insert(header::COOKIE, cookie_header);

    let client = Client::builder().default_headers(headers).build()?;

    let res = client
        .get(format!("https://adventofcode.com/2020/day/{}/input", day).as_str())
        .send()
        .await?;

    fs::write(output_file, res.text().await?)?;
    Ok(())
}

#[derive(serde::Deserialize)]
struct Secrets {
    session_token: String,
}

fn load_secrets() -> Result<Secrets, Box<dyn Error>> {
    println!("loading secrets...");
    let data = fs::read_to_string("secrets.json")?;
    let secrets: Secrets = serde_json::from_str(data.as_str())?;
    println!("secrets loaded");
    Ok(secrets)
}
