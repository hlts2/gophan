
//[0] Javascript Code [1]URL, [2]OutPut
var args = require('system').args
var page = require('webpage').create()

if (args.length < 3) {
    phantom.exit()
}

page.open(args[1], function(status) {
    if(status === "success") {
        page.render(args[2])
    }
    phantom.exit()
});
