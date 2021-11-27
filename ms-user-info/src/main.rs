#![feature(proc_macro_hygiene, decl_macro)]

#[macro_use] extern crate rocket;

use rocket::http::Status;


#[get("/info/<id>")]
fn info(id: i32) -> String {
    return format!("Info of user with id: {}", id);
}

#[get("/health")]
fn health() -> Status {
    return Status::Ok;
}

fn main() {
    rocket::ignite().mount("/", routes![info, health]).launch();
}