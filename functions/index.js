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

  const autor = event.queryStringParameters.autor || '';
  if (autor == ''){
    const result = "Poemas de Federico García Lorca: " + data["poetas"]["Federico García Lorca"]["poemas"].toString()
  }
  else{
    const result = "Poemas de Rafael Alberti: " + data["poetas"]["Rafael Alberti"]["poemas"].toString()
  }

  return {
    statusCode: 200,
    body: result
  }
}