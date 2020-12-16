import java.math.BigDecimal
import kotlin.test.assertEquals
import kotlin.test.assertFailsWith
import kotlin.test.fail
import org.junit.Test

class moduloTests {

    @Test
    fun modulo_whenValid_returnsResultAsString() {
        try {
            val result: String = modulo(BigDecimal("4"), BigDecimal("2"))
            assertEquals(result, "${4 % 2}")
        } catch (e: Exception) {
            fail()
        }
    }

    @Test
    fun modulo_whenValid_returnsResultAsString1() {
        try {
            val result: String = modulo(BigDecimal("2e4"), BigDecimal("4"))
            assertEquals(result, "0")
        } catch (e: Exception) {
            fail("Test")
        }
    }

    @Test
    fun modulo_whenInvalid_ThrowsError() {
        assertFailsWith(Error::class) {
            modulo(null, BigDecimal("4"))
        }
    }

    @Test
    fun modulo_whenInvalid_ThrowsError1() {
        assertFailsWith(Error::class) {
            modulo(BigDecimal("4"), null)
        }
    }
}
