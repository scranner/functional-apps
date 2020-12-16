import io.ktor.application.Application
import io.ktor.application.call
import io.ktor.http.HttpStatusCode
import io.ktor.response.respondText
import io.ktor.routing.get
import io.ktor.routing.routing

fun Application.module(testing: Boolean = false) {

    routing {
        get("/") {
            val x = call.request.queryParameters["x"]?.toBigDecimalOrNull()
            val y = call.request.queryParameters["y"]?.toBigDecimalOrNull()

            try {
                val result = modulo(x, y)
                call.respondText("{\"result\":\"$result\"}")
            } catch (e: Error) {
                call.response.status(HttpStatusCode.BadRequest)
                call.respondText("Invalid Parameters")
            }
        }
        get("/live") {
            call.respondText("OK")
        }
        get("/ready") {
            call.respondText("OK")
        }
    }
}
