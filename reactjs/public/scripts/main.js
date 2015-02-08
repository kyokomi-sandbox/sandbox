var React  = require("React");

var CommentBox = React.createFactory(require("./example.js"));

React.render(
  CommentBox({pollInterval: 2000, url: "comments.json"}),
  document.getElementById('content')
);