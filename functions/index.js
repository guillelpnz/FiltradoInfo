const devolverPoemas = require('./devolverPoemas');
const lib = require('./devolverPoemas')

console.log(lib.devolverPoemas("Rafael Alberti"))

exports.handler = async event => {
  const autor = event.queryStringParameters.autor || '';

  if (autor != ''){
    result = devolverPoemas(autor)
  }

  return {
    statusCode: 200,
    body: result.toString(),
  }
}