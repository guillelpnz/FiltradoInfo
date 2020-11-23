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

exports.handler = async event => {
  const autor = event.queryStringParameters.autor || '';
  if (autor != ''){
    //result = data["poetas"][autor.toString()]["poemas"]
    result = autor.toString()
  }

  return {
    statusCode: 200,
    body: result.toString()
  }
}