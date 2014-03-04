var connect = require('connect'),
    http    = require('http');

var parseRoute = function (url) {
  var attributeParts = url.split('/');
  var attributes = [];

  console.log(attributeParts);
  return url;
}

var app = connect()
  .use(connect.logger('dev'))
  .use(function (request, response) {
    if(request.url === '/lists'){
      response.end("pizza list");
    } else if(parseRoute('/lists/:id')) {
      response.end("Here's your pizza id: id");
    } else {
      response.end("Hello. You shouldn't be seeing this. Tell your web master he's an idiot.");
    }
   });

var server = http.createServer(app);
server.listen(3000);
