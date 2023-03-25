use std::io::Write;
use std::net::{TcpListener, TcpStream};
use rand::Rng;
use serde::{Serialize, Deserialize};
use chrono::{Utc, DateTime};

#[derive(Serialize, Deserialize)]
struct Response {
    time: String,
    password: String,
}

fn handle_client(mut stream: TcpStream) {
    let now: DateTime<Utc> = Utc::now();
    let time = now.to_rfc3339();

    let password: String = rand::thread_rng()
        .sample_iter(&rand::distributions::Alphanumeric)
        .take(12)
        .map(char::from)
        .collect();

    let response = Response {
        time: time,
        password: password,
    };

    let response_json = serde_json::to_string(&response).unwrap();
    let response = format!("HTTP/1.1 200 OK\r\nContent-Type: application/json\r\nContent-Length: {}\r\n\r\n{}", response_json.len(), response_json);

    stream.write(response.as_bytes()).unwrap();
    stream.flush().unwrap();
}


fn main() {
    let listener = TcpListener::bind("0.0.0.0:8080")
        .expect("Failed to start server. Try a different port");

    println!("Listening on http://0.0.0.0:8080...");

    for stream in listener.incoming() {
        match stream {
            Ok(stream) => handle_client(stream),
            Err(e) => eprintln!("Error: {}", e),
        }
    }
}
