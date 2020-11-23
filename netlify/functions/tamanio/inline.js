const aqui = "Mis pasos en esta calle resuenan en otra calle donde oigo mis pasos pasar en esta calle donde sólo es real la niebla."
module.exports = async (ctx) => {
    const search = (ctx.inlineQuery.query || "")
    if ( search==="" || isNaN(search)) {
        return
    } else {
        if ( search === "tamanioaqui" ){
            const answer = []
            let tam = aqui.split(" ").length
            answer.push({
                id: aqui,
                title: "Tamaño del poema aquí de Octavio Paz:",
                type: 'article',
                input_message_content: {
                    message_text: tam,
                    parse_mode: 'HTML'
                }
            })
        }
        if (search === "aqui"){
            const answer = []
            let tam = aqui.split(" ").length
            answer.push({
                id: "Aquí",
                title: "Poema aquí de Octavio Paz:",
                type: 'article',
                input_message_content: {
                    message_text: aqui,
                    parse_mode: 'HTML'
                }
            })
        }

        return ctx.answerInlineQuery(answer)
    }
}