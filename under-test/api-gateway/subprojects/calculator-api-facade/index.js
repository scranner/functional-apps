const express = require('express');
const graphqlHTTP = require('express-graphql');
const cors = require('cors');
const http = require('http');
const { installHealthChecks } = require('./src/health-checks');
const { schema } = require('./src/schema');
const { rootValue } = require('./src/root-value');

var app = express();

app.use('/graphql', cors(), graphqlHTTP({
    schema,
    rootValue,
    graphiql: true
}));

const server = http.createServer(app);

installHealthChecks(server);

server.listen(80);

console.log('Running a GraphQL API server at http://localhost/graphql');