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
  if (autor == ''){
    result = "Poemas de Federico García Lorca: "+data["poetas"]["Federico García Lorca"]["poemas"].toString()
  }
  else{
    result = "Poemas de Rafael Alberti: "+data["poetas"]["Rafael Alberti"]["poemas"].toString()
  }

  return {
    statusCode: 200,
    body: result.toString()
  }
}