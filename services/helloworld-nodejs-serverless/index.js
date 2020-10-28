const express = require('express');
const app = express();

app.get('/', (req, res) => {
  console.log('Hello world');

  res.send(`Hello world`);
});

const port = process.env.PT_CONTAINER_PORT;
app.listen(port, () => {
  console.log('Hello world listening on port', port);
});
