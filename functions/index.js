exports.handler = async event => {
  let data = {
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

  const autor = event.queryStringParameters.autor || '';
  if (autor != ''){
    //result = data["poetas"][autor.toString()]["poemas"]
    result = data["poetas"][autor.toString()]
  }
  else{
    //result = data["poetas"]["Rafael Alberti"]["poemas"]
    result = autor
  }

  return {
    statusCode: 200,
    body: result
  }
}