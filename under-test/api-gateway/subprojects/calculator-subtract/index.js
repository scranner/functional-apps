const express = require('express');
const { route } = require('./src/main/router');
const { PORT } = require('./src/main/config');

const app = express();

route(app);

app.listen(PORT);

console.log(`Server Listening on port: localhost:${PORT}`);
