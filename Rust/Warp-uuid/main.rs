use serde_json::json;
use uuid::Uuid;
use warp::Filter;

#[tokio::main]
async fn main() {
    // GET /uuid -> {"uuid": "..."}
    let route = warp::path!("uuid").map(|| warp::reply::json(&json!({ "uuid": Uuid::new_v4() })));

    println!("GET -> http://localhost:3000/uuid");
    warp::serve(route).run(([0, 0, 0, 0], 3000)).await;
}