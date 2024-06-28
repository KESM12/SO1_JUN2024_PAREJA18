#[macro_use] extern crate rocket;

use rocket::serde::json::Json;
use rocket::serde::{Deserialize, Serialize};
use redis::Commands;

#[derive(Deserialize, Serialize)]
struct Data {
    Texto: String,
    Pais: String,
}

#[post("/set", format = "json", data = "<data>")]
async fn set_data(data: Json<Data>) -> Result<&'static str, &'static str> {
    // Crear cliente de redis
    let client = redis::Client::open("redis://redis:6379/") //localhost
        .map_err(|_| "Error al crear el cliente de Redis")?;

    // Conexion a redis
    let mut con = client.get_connection()
        .map_err(|_| "Error al conectarse con redis.")?;

    //Insertar hash en redis 
    let _: () = con.hincr("paises", &data.Pais, 1)
    .map_err(|_| "Failed to set data in Redis")?;
    Ok("Data set")
}

#[launch]
fn rocket() -> _ {
    rocket::build().mount("/", routes![set_data])
}