package phantom

//According to PhantomJS dev's comment on GitHub, the full support of ES6 will come with PhantomJS 2.5.
var js = `
var page = require('webpage').create()
var args = require('system').args

//([0]javascript [1] URL, [2] Output File)
page.open(args[1], function(state) {
	if (state == "success") {
		page.render(args[2])
	}
	phantom.exit()
})
`
