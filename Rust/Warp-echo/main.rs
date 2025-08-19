use serde::Deserialize;
use serde_json::json;
use warp::Filter;

#[derive(Deserialize)]
struct In { msg: String }

#[tokio::main]
async fn main() {
    // POST /echo  {"msg":"hello"} -> {"echo":"hello"}
    let echo = warp::path!("echo")
        .and(warp::post())
        .and(warp::body::json())
        .map(|input: In| warp::reply::json(&json!({ "echo": input.msg })));

    println!(r#"POST JSON -> http://localhost:3001/echo  {"msg":"hello"}"#);
    warp::serve(echo).run(([0, 0, 0, 0], 3001)).await;
}