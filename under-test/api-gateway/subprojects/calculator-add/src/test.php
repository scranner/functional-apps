<?php
echo "Test Script Starting\n";
require('functions.inc.php');

function testFunction($x, $y, $expect) {
    $answer=add($x, $y);

    $result = ($answer==$expect) ? "PASSED" : "FAILURE";
    echo "Test Result: ".$x."+".$y."=".$answer." (expected: ".$expect.")\t".$result."\n";
    if ($answer==$expect) {
        return false;
    }
    return true;
}

function en($num) {
    return json_encode(array("result" => "".$num));
}

function es($num) {
    return json_encode($num);
}

$errorString = "Invalid Parameters";

$testResult1 = testFunction(10, 5, en(15));
$testResult2 = testFunction("10", "5", en(15));
$testResult3 = testFunction("test", "5", es($errorString));
$testResult4 = testFunction("10", "wee", es($errorString));
$testResult5 = testFunction("10e2", "12", en(1012));
$testResult6 = testFunction(null, "beep", es($errorString));

if($testResult1 || $testResult2 ||
   $testResult3 || $testResult4 ||
   $testResult5 || $testResult6) {
    exit(1);
}






