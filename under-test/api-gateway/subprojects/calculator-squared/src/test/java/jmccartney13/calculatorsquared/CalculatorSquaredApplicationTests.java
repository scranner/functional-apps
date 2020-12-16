package jmccartney13.calculatorsquared;

import org.junit.jupiter.api.Test;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import java.math.BigDecimal;

import static jmccartney13.calculatorsquared.Squared.squared;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.fail;

class CalculatorSquaredApplicationTests {

	@Test
	void squared_whenNull_Returns400() {
		assertEquals(
				new ResponseEntity<>("Invalid Parameters", HttpStatus.BAD_REQUEST),
				squared(null));
	}

	@Test
	void squared_whenString_Returns400() {
		try {
			new BigDecimal("weewoo");
			fail();
		} catch(Exception e) {}
	}

	@Test()
	void squared_whenValid_Returns200() {
		BigDecimal valueToSquare = new BigDecimal(String.valueOf(4));
		assertEquals(
				new ResponseEntity<>("{\"result\":\"" + valueToSquare.pow(2).toString() + "\"}", HttpStatus.OK),
				squared(valueToSquare));

	}

	@Test
	void squared_whenValidEx_Returns200() {
		BigDecimal valueToSquare = new BigDecimal(String.valueOf(2e2));
		assertEquals(
				new ResponseEntity<>("{\"result\":\"" + valueToSquare.pow(2).toBigInteger().toString() + "\"}", HttpStatus.OK),
				squared(valueToSquare));

	}

}
