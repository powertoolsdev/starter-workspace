const express = require('express');
const app = express();

app.get('/', (req, res) => {
  console.log('about page received a request.');

  res.send(`about page changed! -- hello monica`);
});

const port = process.env.PT_CONTAINER_PORT;
app.listen(port, () => {
  console.log('about listening on port', port);
});
