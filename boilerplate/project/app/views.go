package app

import "gongo/core"

func HomePage(req *core.Request) core.Response {
	return core.HttpResponse("Welcome to Gongo!")
}
