package jmccartney13.calculatorsquared;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.math.BigDecimal;
import java.math.BigInteger;
import java.util.logging.Level;
import java.util.logging.Logger;

import static jmccartney13.calculatorsquared.Squared.squared;

@RestController
class SquaredController {

    final Logger logger = Logger.getLogger(SquaredController.class.getName());

    @RequestMapping("/")
    public ResponseEntity<String> squareValue(
            @RequestParam(value="x", required=false) BigDecimal x) {
        logger.log(Level.INFO, "Processing request for x: " + x);
        return squared(x);
    }

    @RequestMapping("/live")
    public ResponseEntity<String> live() {
        return new ResponseEntity<>("OK", HttpStatus.OK);
    }

    @RequestMapping("/ready")
    public ResponseEntity<String> ready() {
        return new ResponseEntity<>("OK", HttpStatus.OK);
    }
}