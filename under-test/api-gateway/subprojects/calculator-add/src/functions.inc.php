<?php

function add($x, $y) {
    if(is_numeric($x) && is_numeric($y)) {
        $answer=$x+$y;
        return json_encode(array("result" => "".$answer.""));
    }
    http_response_code(400);
    return json_encode("Invalid Parameters");
}
