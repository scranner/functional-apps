const { createTerminus } = require('@godaddy/terminus');

const options = {
    healthChecks: {
        '/live': () => { Promise.resolve(JSON.stringify({ live: 'OK' })) },
        '/ready': () => { Promise.resolve(JSON.stringify({ ready: 'OK' })) }
    }
};

module.exports = {
    installHealthChecks: server => createTerminus(server, options),
};