const d = require('./devolverPoemas');

exports.handler = async event => {
  const autor = event.queryStringParameters.autor || '';
  let result = ''
  if (autor != ''){
    result = d.devolverPoemas(autor)
  }

  return {
    statusCode: 200,
    body: result.toString(),
  }
}