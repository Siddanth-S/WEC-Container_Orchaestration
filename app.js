// app.js
const express = require('express');
const app = express();
const port = process.env.PORT || 3000;
const appName = process.env.APP_NAME || "Unknown App";

app.get('/', (req, res) => {
    res.send(`Welcome to ${appName}!`);
});

app.listen(port, () => {
    console.log(`App is running on port ${port}`);
});