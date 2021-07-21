use std::env;
use regex::Regex;
use std::process::{Stdio};
use execute::{Execute, shell};

fn run(command: &String) {
    let mut to_run = shell(&command);
    to_run.stdout(Stdio::piped());
    let output = to_run.execute_output().unwrap();
    println!("{}", String::from_utf8(output.stdout).unwrap());
}

fn main() {
    let args: Vec<String> = env::args().skip(1).collect();
    let command = args.join(" ");

    let re = Regex::new(r"[sudo]{4}").unwrap();

    if re.is_match(&command) {
        run(&command.replace("sudo ", ""))
    } else {
        run(&command)
    }
}
