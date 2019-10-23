// Copyright (C) 2019 Vanessa Sochat.
// This Source Code Form is subject to the terms of the
// Mozilla Public License, v. 2.0. If a copy of the MPL was not distributed
// with this file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main
 
import (
	"fmt"	
	nu "github.com/vsoch/nu-plugin/pkg/plugin"
)

// sink is required for you to implement to parse parameters and print
// whatever you like to stdout!  It should take the plugin object passed, 
// along with params from nushell. You are responsible 
// for parsing the params, and can use both plugin.Func.getStringPrimitive 
// and plugin.Func.getIntPrimitive
func sink(plugin *nu.SinkPlugin, params interface{}) {

	// can also be getIntPrimitive
	// still testing here
	namedParams := plugin.Func.GetNamedParams(params)	
	fmt.Println(namedParams)
	
	// You must also return the tag with your response
	//tag := plugin.Func.GetTag(params)

	// This can also be printStringResponse
	//plugin.Func.PrintIntResponse(intLength, tag)

}

func main() {

	name := "coffee"
	usage := "Generate ascii coffee on demand"
	plugin := nu.NewSinkPlugin(name, usage)

	// Add named arguments: key, config.SyntaxShape, config.ArgType
	// Switch should be provided with empty string
	plugin.Config.AddNamedParam("sugar", "Switch", "")
	plugin.Config.AddNamedParam("black", "Switch", "")
	plugin.Config.AddNamedParam("strength", "Optional", "String")

	// Run the filter function
	plugin.Run(sink)
}
