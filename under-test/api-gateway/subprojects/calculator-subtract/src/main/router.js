const { subtract } = require('./subtract');

const route = app => {
    app.get('/', (req,res) => {
        const resp = subtract(req.query.x, req.query.y);

        if (!resp) {
            res.status(400);
            res.send(JSON.stringify({ error: 'Invalid Parameter'}))
        } else {
            res.send(JSON.stringify({ result: `${resp}` }))
        }
    });

    app.get('/live', (req,res) => res.send(JSON.stringify({ live: "OK "})));
    app.get('/ready', (req,res) => res.send(JSON.stringify({ ready: "OK "})))
};

module.exports = { route };