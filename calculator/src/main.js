import express from 'express';

const app = express();
const port = 80;

app.get('/random', (req, res) => {
    res.send((Math.floor(Math.random() * 100)).toString());
});

app.get('/hc', (req, res) => {
    console.log('HC of calculator');
    res.sendStatus(200);
});

app.listen(port, () => {
    console.log('Calculator is UP');
});