const express = require('express');

const app = express();
const port = 8080;

// Middleware untuk log IP publik
app.use((req, res, next) => {
  const xff = req.headers['x-forwarded-for'];
  const clientIP = xff ? xff.split(',')[0].trim() : req.socket.remoteAddress;
  console.log(`Received request: ${req.method} ${req.url} from ${clientIP}`);
  next();
});

app.get('/', (req, res) => {
  res.send('Hello from Node Service! 🚀');
});

app.listen(port, () => {
  console.log(`Node service running on :${port}`);
});