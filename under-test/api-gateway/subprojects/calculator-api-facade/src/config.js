const _ = require('lodash');

const getUrl = (envVariable, serviceDiscoveryEnvVariable, fallbackUrl) => {
    const envVar1 = process.env[envVariable];
    if (envVar1) return envVar1; //if the user passed custom urls, use them
    const envVar2 = process.env[serviceDiscoveryEnvVariable];
    if (envVar2) return envVar2; // otherwise try known service discovery variables
    return fallbackUrl // otherwise use known ingress
};

const config = {
    PORT: 80,
    subtractUrl:
        getUrl('SUBTRACT_URL',
            'SUBTRACT_CALCULATOR_SUBTRACT_SERVICE_HOST',
            'subtract.calculator-project.qpc.hal.davecutting.uk'),
    factorialUrl:
        getUrl('FACTORIAL_URL',
            'FACTORIAL_CALCULATOR_FACTORIAL_SERVICE_HOST',
            'factorial.calculator-project.qpc.hal.davecutting.uk'),
    addUrl:
        getUrl('ADD_URL',
            'ADD_CALCULATOR_ADD_SERVICE_HOST',
            'add.calculator-project.qpc.hal.davecutting.uk'),
    squaredUrl:
        getUrl('SQUARED_URL',
            'SQUARED_CALCULATOR_SQUARED_SERVICE_HOST',
            'squared.calculator-project.qpc.hal.davecutting.uk'),
    moduloUrl:
        getUrl('MODULO_URL',
            'MODULO_CALCULATOR_MODULO_SERVICE_HOST',
            'modulo.calculator-project.qpc.hal.davecutting.uk'),
    multiplyUrl:
        getUrl('MULTIPLY_URL',
            'MULTIPLY_CALCULATOR_MULTIPLY_SERVICE_HOST',
            'multiply.calculator-project.qpc.hal.davecutting.uk'),
};

module.exports = config;