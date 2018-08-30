// Copyright 2016 The go-vgo Project Developers. See the COPYRIGHT
// file at the top-level directory of this distribution and at
// https://github.com/go-vgo/robotgo/blob/master/LICENSE
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// http://www.apache.org/licenses/LICENSE-2.0> or the MIT license
// <LICENSE-MIT or http://opensource.org/licenses/MIT>, at your
// option. This file may not be copied, modified, or distributed
// except according to those terms.

package main

import (
	"fmt"
	"time"
	"net/http"
	"github.com/go-vgo/robotgo"
	// "go-vgo/robotgo"
	"bytes"
	"encoding/json"
	"os/user"
)

type Wrapper struct {
	User string `json:"user"`
	Location string `json:"location"`
	Action string `json:"action"`
}

func send(user string, title string) {
	url := "http://localhost:3000/window"

	jsonStr, err := json.Marshal(Wrapper{User: user, Location: title, Action: "change"})
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("x-project-token", "TOKEN")
	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Err sending telemetry - ", err)
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()
}

func window() {
	////////////////////////////////////////////////////////////////////////////////
	// Window Handle
	////////////////////////////////////////////////////////////////////////////////

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	currentTitle := ""
	for {
		// get current Window title
		title := robotgo.GetTitle()

		s, e := robotgo.FindName(robotgo.GetPID())
		if e == nil {
			if title != currentTitle {
				currentTitle = title
				toSend := "[" + s + "] " + currentTitle
				fmt.Println(toSend)
				send(user.Username, toSend)
			}
		}

		time.Sleep(5000 * time.Millisecond)
	}
}

func main() {
	window()
}
