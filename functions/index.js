const request = require('request-promise');
const data = require('./poemas.json')
exports.handler = async function (event, context){

  let body = JSON.parse(event.body);
  let {chat, text} = body.message;

  let result = ""

  switch(texto){
    case "alberti":
      result = "Poemas de Rafael Alberti: " + data["poetas"]["Rafael Alberti"]["poemas"].toString();
      break;
    
    case "lorca":
      result = "Poemas de Federico García Lorca: " + data["poetas"]["Federico García Lorca"]["poemas"].toString();
      break;
    default:
      result = "Use /alberti o /lorca para conocer sus poemas más famosos";
      break;
  }

  return {
    statusCode: 200,
    body: JSON.stringify({text: result, method: 'sendMessage', chat_id:chat.id}),
    headers:{
      'Content-Type': 'application/json'
    }
  }
}