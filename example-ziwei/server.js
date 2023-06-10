const express = require('express');
const app = express();

app.use(express.static(__dirname + '/')); // 设置静态文件目录

app.get('/', function(req, res) {
    res.sendFile(__dirname + '/Users/dradon/Downloads/zwds-amrta/example-ziwei/home.html'); // 设置默认页面
});

app.listen(3000, function() {
    console.log('Server is running on port 3000');
});