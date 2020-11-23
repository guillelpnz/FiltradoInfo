const Telegraf = require('telegraf');
const startAction = require('./tamanio/start')
const inlineAction = require('./tamanio/inline')
const bot = new Telegraf(process.env.TAMANIO_BOT_TOKEN);

bot.start(ctx => {
return startAction(ctx)
})

bot.on('inline_query', (ctx) => {
return inlineAction(ctx)
})

exports.handler = async event => {
await bot.handleUpdate(JSON.parse(event.body));
return { statusCode: 200, body: '' };
}