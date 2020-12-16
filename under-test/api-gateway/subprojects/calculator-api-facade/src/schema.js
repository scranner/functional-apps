const { buildSchema } = require('graphql');

const schema = buildSchema(`

  type LivenessResponse {
    url: String!
    status: String!
  }

  type Query {
    add(x: String!, y: String!): String!,
    factorial(x: String!): String!,
    modulo(x: String!, y: String!): String!,
    multiply(x: String!, y: String!): String!,
    squared(x: String!): String!,
    subtract(x: String!, y: String!): String!,
    liveness: [LivenessResponse]!
  }
`);

module.exports = { schema };