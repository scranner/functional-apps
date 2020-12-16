const axios = require('axios');
const _ = require('lodash');
const CircuitBreaker = require('opossum');
const {
    subtractUrl,
    factorialUrl,
    addUrl,
    squaredUrl,
    moduloUrl,
    multiplyUrl
} = require('./config');

const wrapCall2Param = funcToWrap => {
    const breaker = new CircuitBreaker((x, y) => funcToWrap(x, y), {
        timeout: 5000, // 1 seconds is a failure
        errorThresholdPercentage: 50, // 50% of reqs fail and we trip the circuit
        resetTimeout: 5000, // circuit untrips after 5 seconds to avoid spam clicking.
        allowWarmUp: true, // warmup in case of early deployment jitters
        errorFilter: (error) => error.status === 400, //only trip on 400s
    });
    breaker.fallback(() => 'Error: service unavailable');
    return breaker;
};

const createDataSource2params = (url) => async ({x, y}) => {
    let result;
    try {
        const request = await axios.get(`http://${url}`, { params: {x:x, y:y}});
        result = _.get(request, 'data.result', undefined)
    } catch (e) { console.log(`Request to subtract url: ${url} failed.\n${e.stack}`) }
    return result;
};

const paramWrapper = (url) => {
    const wrappedFunction = wrapCall2Param(createDataSource2params(url));
    return (x,y) => wrappedFunction.fire(x, y);
};

const testLiveness = (subtractUrl, factorialUrl, addUrl, squaredUrl, moduloUrl, multiplyUrl) => async () => {
    const urls = [subtractUrl, factorialUrl, addUrl, squaredUrl, moduloUrl, multiplyUrl];
        const results = await Promise.all(urls.map( async url => {
            let request = {};

            try {
                request = await axios.get(`http://${url}/ready`, {
                    timeout: 10 * 1000,
                });
                } catch (e) {
                    request.status = 400;
                }
            if (request.status === 200) return { url, status: 'available' };
            return { url, status: 'error' };
        }));
        return results;
};

const rootValue = {
    subtract: paramWrapper(subtractUrl),
    factorial: paramWrapper(factorialUrl),
    add: paramWrapper(addUrl),
    squared: paramWrapper(squaredUrl),
    modulo: paramWrapper(moduloUrl),
    multiply: paramWrapper(multiplyUrl),
    liveness: testLiveness(subtractUrl, factorialUrl, addUrl, squaredUrl, moduloUrl, multiplyUrl),
};

module.exports = { rootValue };