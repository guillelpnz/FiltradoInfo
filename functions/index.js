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
function devolverPoemas(autor){
  let poemas = data["poetas"][autor]["poemas"]

  if (poemas != ""){
    return poemas
  }
}

// console.log(devolverPoemas("Federico García Lorca"))

exports.handler = async event => {
  const autor = event.queryStringParameters.autor || '';

  let result = ''
  if (autor != ''){
    result = devolverPoemas(autor.toString())
  }

  return {
    statusCode: 200,
    body: result.toString()
  }
}