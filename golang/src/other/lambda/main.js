var child_process = require('child_process');

exports.handler = function (event, context) {
    console.log("Hello World Node.js");
    var proc = child_process.spawn('./example', [JSON.stringify(event)], {stdio: 'inherit'});

    proc.on('close', function (code) {
        if (code !== 0) {
            return context.done(new Error("Process exited with non-zero status code"));
        }

        context.done(null, "SUCCESS");
    });
}
