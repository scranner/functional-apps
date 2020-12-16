docker build calculator-add -t "$1"/calculator-add &&
docker build calculator-api-facade -t "$1"/calculator-api-facade &&
docker build calculator-factorial -t "$1"/calculator-factorial &&
docker build calculator-modulo -t "$1"/calculator-modulo &&
docker build calculator-multiply -t "$1"/calculator-multiply &&
docker build calculator-squared -t "$1"/calculator-squared &&
docker build calculator-subtract -t "$1"/calculator-subtract