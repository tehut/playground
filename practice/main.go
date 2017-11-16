package main

import (
	"strings"
)

type showRequest struct {
	AppName      string                 `json:"appName"`
	Components   map[string]interface{} `json:"components"`
	Environments map[string]interface{} `json:"environments"`
	Vendor       map[string]interface{} `json:"vendor"`
}

func responseCrawler(v map[string]interface{}, a map[string]string) map[string]string {
	e := 0
	for i := range v {
		switch v[i].(type) {
		case map[string]interface{}:
			e++
			responseCrawler(v[i].(map[string]interface{}), a)

		case string:
			// fmt.Println("this is a string")
			if strings.Index(i, "onnet") != -1 {
				a[i] = v[i].(string)
			}
		}
	}

	if e == 0 {
		return a
	}

}
func main() {
	resp := showRequest{
		AppName: "greatapp",
		Components: map[string]interface{}{
			"params.libsonnet": "{\n  global: {\n    // User-defined global parameters; accessible to all component and environments, Ex:\n    // replicas: 4,\n  },\n  components: {\n    // Component-level parameters, defined initially from 'ks prototype use ...'\n    // Each object below should correspond to a component in the components/ directory\n  },\n}\n",
		},
		Environments: map[string]interface{}{
			"base.libsonnet": "local components = std.extVar(\"__ksonnet/components\");\ncomponents + {\n  // Insert user-specified overrides here.\n}\n",
			"default": map[string]interface{}{
				"default.jsonnet":  "local base = import \"/Users/tehut/test/environments/base.libsonnet\";\nlocal k = import \"k.libsonnet\";\n\nbase + {\n  // Insert user-specified overrides here. For example if a component is named \"nginx-deployment\", you might have something like:\n  //   \"nginx-deployment\"+: k.deployment.mixin.metadata.labels({foo: \"bar\"})\n}\n",
				"params.libsonnet": "local params = import \"/Users/tehut/test/components/params.libsonnet\";\nparams + {\n  components +: {\n    // Insert component parameter overrides here. Ex:\n    // guestbook +: {\n    //   name: \"guestbook-dev\",\n    //   replicas: params.global.replicas,\n    // },\n  },\n}\n",
				"spec.json":        "{\n  \"server\": \"https://kubecfg-t-apiloadb-1k1vqoxu7q212-666493100.us-west-2.elb.amazonaws.com\",\n  \"namespace\": \"dev-tehut\"\n}",
			},
		},
		Vendor: map[string]interface{}{
			"base.libsonnet": "local components = std.extVar(\"__ksonnet/components\");\ncomponents + {\n  // Insert user-specified overrides here.\n}\n",
			"redis": map[string]interface{}{
				"default.jsonnet":  "local base = import \"/Users/tehut/test/environments/base.libsonnet\";\nlocal k = import \"k.libsonnet\";\n\nbase + {\n  // Insert user-specified overrides here. For example if a component is named \"nginx-deployment\", you might have something like:\n  //   \"nginx-deployment\"+: k.deployment.mixin.metadata.labels({foo: \"bar\"})\n}\n",
				"params.libsonnet": "local params = import \"/Users/tehut/test/components/params.libsonnet\";\nparams + {\n  components +: {\n    // Insert component parameter overrides here. Ex:\n    // guestbook +: {\n    //   name: \"guestbook-dev\",\n    //   replicas: params.global.replicas,\n    // },\n  },\n}\n",
				"spec.json":        "{\n  \"server\": \"https://kubecfg-t-apiloadb-1k1vqoxu7q212-666493100.us-west-2.elb.amazonaws.com\",\n  \"namespace\": \"dev-tehut\"\n}",
			},
		}}
	a := make(map[string]string)
	responseCrawler(resp.Environments, a)
}
