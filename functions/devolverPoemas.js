const data = require('./poemas.json');

function devolverPoemas(autor){
    let poemas = data["poemas"]["poetas"][autor]["poemas"]

    if (poemas != ""){
        return poemas
    }
}

module.exports = { devolverPoemas }