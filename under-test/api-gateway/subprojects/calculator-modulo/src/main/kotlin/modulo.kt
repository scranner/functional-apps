import java.lang.Error
import java.math.BigDecimal

fun modulo(x: BigDecimal?, y: BigDecimal?): String {
    if (x == null || y == null) {
        throw Error("Invalid Parameters")
    }
    return (x % y).toBigInteger().toString()
}
