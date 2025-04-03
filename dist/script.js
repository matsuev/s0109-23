// Use WebSocket transport endpoint.
const client = new Centrifuge('ws://127.0.0.1:8000/connection/websocket');

// Allocate Subscription to a channel.
const newsSub = client.newSubscription('news');

// React on `news` channel real-time publications.
newsSub.on('publication', function (ctx)
{
   console.log(ctx);
});

// Trigger subscribe process.
newsSub.subscribe();

// Trigger actual connection establishement.
client.connect();