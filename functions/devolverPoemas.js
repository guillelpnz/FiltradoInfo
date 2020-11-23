const data = require('./poemas.json');

function devolverPoemas(autor){
    let poemas = data["poetas"][autor]["poemas"]

    if (poemas != ""){
        return poemas
    }
}

module.exports = { devolverPoemas }