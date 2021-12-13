import express from 'express';
import httpProxy from 'express-http-proxy';

const app = express();
const port = 8080;

const calculatorServiceProxy = httpProxy('http://calculator');
const userServiceProxy = httpProxy('http://user');

app.get('/', (req, res) => {
    res.send('Cool');
});

app.use('/calculator', calculatorServiceProxy);

app.use('/user', userServiceProxy);

app.get('/hc', (req, res) => {
    console.log('HC of gateway');
    res.sendStatus(200);
})

app.listen(port, () => {
    console.log('Gateway is UP');
})