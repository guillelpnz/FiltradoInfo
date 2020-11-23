const request = require('request-promise');
exports.handler = async event => {
  const data = {
    "poetas": {
      "Federico García Lorca": {
        "poemas": [
            "Casida del Llanto",
            "Casada infiel",
            "Asesinato",
            "Soneto de la dulce queja",
            "Danza da lúa en Santiago"
        ]
      },
      "Rafael Alberti": {
        "poemas": [
            "Amaranta",
            "Nocturno",
            "La paloma",
            "A galopar"
        ]
      }
    }
  }

  
  const body = JSON.parse(event.body);
  let {conver, texto} = body.message;

  let result = ""

  switch(texto){
    case "alberti":
      result = "Poemas de Rafael Alberti: " + data["poetas"]["Rafael Alberti"]["poemas"].toString();
      break;
    
    case "lorca":
      result = "Poemas de Federico García Lorca: " + data["poetas"]["Federico García Lorca"]["poemas"].toString();
      break;
    default:
      result = "Use /alberti o /lorca para conocer sus poemas más famosos"
  }

  return {
    statusCode: 200,
    body: JSON.stringify({texto:result, method:'sendMessage', chat_id:conver.id}),
    headers:{
        'Content-Type': 'application/json'
    }
  }
}