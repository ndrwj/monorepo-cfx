const express = require('express');
const app = express();
const port = 8080;

app.get('/', (req, res) => {
  res.send('Hello from Node.js Service!');
});

app.listen(port, () => {
  console.log(`Node.js service running on port ${port}`);
});
