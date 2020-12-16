package jmccartney13.calculatorsquared;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import java.math.BigDecimal;

class Squared {
    public static ResponseEntity<String> squared(BigDecimal toSquare) {
        if (toSquare == null) return new ResponseEntity<>("Invalid Parameters", HttpStatus.BAD_REQUEST);
        return new ResponseEntity<>("{\"result\":\"" + toSquare.pow(2).toBigInteger().toString() + "\"}", HttpStatus.OK);
    }
}