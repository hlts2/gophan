package phantom

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

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

func loadjs() string {
	jsP := filepath.Join(getCachePath(), "load.js")

	if _, err := os.Stat(jsP); os.IsNotExist(err) {
		ioutil.WriteFile(jsP, []byte(js), os.ModePerm)
	}

	return jsP
}
