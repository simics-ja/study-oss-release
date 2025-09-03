const http = require('http');
const os = require('os');

const port = process.env.PORT || 3000;

const server = http.createServer((req, res) => {
  if (req.url === '/') {
    res.writeHead(200, { 'Content-Type': 'text/html' });
    res.end(`
      <!DOCTYPE html>
      <html>
        <head>
          <title>OSS Release Demo App</title>
          <style>
            body {
              font-family: Arial, sans-serif;
              max-width: 800px;
              margin: 50px auto;
              padding: 20px;
              background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
              color: white;
            }
            .container {
              background: rgba(255, 255, 255, 0.1);
              border-radius: 10px;
              padding: 30px;
              backdrop-filter: blur(10px);
            }
            h1 { margin-bottom: 20px; }
            .info { margin: 10px 0; }
            .label { font-weight: bold; }
          </style>
        </head>
        <body>
          <div class="container">
            <h1>🚀 OSS Release Demo Application</h1>
            <div class="info">
              <span class="label">Container ID:</span> ${os.hostname()}
            </div>
            <div class="info">
              <span class="label">Node Version:</span> ${process.version}
            </div>
            <div class="info">
              <span class="label">Platform:</span> ${os.platform()}
            </div>
            <div class="info">
              <span class="label">Architecture:</span> ${os.arch()}
            </div>
            <div class="info">
              <span class="label">Uptime:</span> ${Math.floor(process.uptime())} seconds
            </div>
            <div class="info">
              <span class="label">Memory Usage:</span> ${Math.round(process.memoryUsage().heapUsed / 1024 / 1024)} MB
            </div>
          </div>
        </body>
      </html>
    `);
  } else if (req.url === '/health') {
    res.writeHead(200, { 'Content-Type': 'application/json' });
    res.end(JSON.stringify({ status: 'healthy', timestamp: new Date().toISOString() }));
  } else {
    res.writeHead(404, { 'Content-Type': 'text/plain' });
    res.end('Not Found');
  }
});

server.listen(port, () => {
  console.log(`Server running at http://localhost:${port}/`);
  console.log('Health check endpoint: /health');
});

process.on('SIGTERM', () => {
  console.log('SIGTERM signal received: closing HTTP server');
  server.close(() => {
    console.log('HTTP server closed');
    process.exit(0);
  });
});