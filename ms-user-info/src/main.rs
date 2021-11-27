use actix_web::{get, web, App, HttpResponse, HttpServer, Responder};


#[get("/info/{id}")]
async fn info(id: web::Path<i32>) -> impl Responder {
    HttpResponse::Ok().body(format!("Info of user with id: {}", id))
}

#[get("/health")]
async fn health() -> impl Responder {
    println!("Health check OK!");
    return HttpResponse::Ok();
}


#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .service(info)
            .service(health)
    })
    .bind("127.0.0.1:8000")?
    .run()
    .await
}