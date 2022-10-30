// Copyright (c) 2022 Yandex LLC. All rights reserved.
// Author: Andrey Khaliullin <avhaliullin@yandex-team.ru>

package main

import (
	"fmt"
	"net/http"
)

func doHandle(req *http.Request) (interface{}, error) {
	path := req.URL.Path
	if len(path) == 0 || path[0] != '/' {
		path = "/" + path
	}
	switch path {
	case "/":
		return map[string]string{"message": "Hello world!"}, nil
	case "/user-error":
		return nil, BadRequest("something is wrong about your request")
	case "/internal-error":
		return nil, fmt.Errorf("server crashed")
	default:
		return nil, NotFound(path)
	}
}
