#!/bin/bash
docker push "$1"/calculator-add &&
docker push "$1"/calculator-api-facade &&
docker push "$1"/calculator-factorial &&
docker push "$1"/calculator-modulo &&
docker push "$1"/calculator-multiply &&
docker push "$1"/calculator-squared &&
docker push "$1"/calculator-subtract