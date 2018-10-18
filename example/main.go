/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"github.com/golang/glog"
	"github.com/kubernetes-csi/csi-lib-fc/fibrechannel"
)

func main() {
	c := fibrechannel.Connector{}
	//Host5 and host6 respectively
	c.TargetWWNs = []string{"10000000c9a02834", "10000000c9a02835"}
	c.Lun = "1"
	dp, err := fibrechannel.Attach(c, &fibrechannel.OSioHandler{})
	glog.Infof("Path is: %s\n", dp)
	if err != nil {
		glog.Errorf("Error from Connect: %s\n", err)
	}

	fibrechannel.Detach(dp, &fibrechannel.OSioHandler{})
}
