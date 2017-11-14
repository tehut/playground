package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/heptio/playground/api"
)

// var limiter *rate.Limiter

func main() {
	corsObj := handlers.AllowedOrigins([]string{"*"})
	// limiter = rate.NewLimiter(config.RateLimit, config.RateLimitBurst)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/init", ksInit)
	router.HandleFunc("/show", ksShow)
	router.HandleFunc("/generate", ksGenerate)
	router.HandleFunc("/depget", index)
	router.HandleFunc("/env", index)
	log.Printf("Serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(corsObj)(router)))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func ksInit(w http.ResponseWriter, r *http.Request) {
	// TODO: make sure to put the return value back api.InitResponse
	resp := api.InitResponse{
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
	}

	init, _ := json.Marshal(resp)
	s := string(init)
	fmt.Fprintf(w, s)
	// return resp
}

func ksShow(w http.ResponseWriter, r *http.Request) {
	resp := api.ShowResponse{
		Components: map[string]interface{}{
			"params.libsonnet": "{\n  global: {\n    // User-defined global parameters; accessible to all component and environments, Ex:\n    // replicas: 4,\n  },\n  components: {\n    // Component-level parameters, defined initially from 'ks prototype use ...'\n    // Each object below should correspond to a component in the components/ directory\n  },\n}\n",
			"ui.libsonnet":     "{\n  global: {\n    // User-defined global parameters; accessible to all component and environments, Ex:\n    // replicas: 4,\n  },\n  components: {\n    // Component-level parameters, defined initially from 'ks prototype use ...'\n    // Each object below should correspond to a component in the components/ directory\n  },\n}\n",
			"redis.libsonnet":  "{\n  global: {\n    // User-defined global parameters; accessible to all component and environments, Ex:\n    // replicas: 4,\n  },\n  components: {\n    // Component-level parameters, defined initially from 'ks prototype use ...'\n    // Each object below should correspond to a component in the components/ directory\n  },\n}\n",
		}}

	show, _ := json.Marshal(resp)
	s := string(show)
	fmt.Fprintf(w, s)
}

func ksGenerate(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "generate.json")
	resp := api.GenerateResponse{
		Components: map[string]interface{}{
			"params.libsonnet": "{\n  global: {\n    // User-defined global parameters; accessible to all component and environments, Ex:\n    // replicas: 4,\n  },\n  components: {\n    // Component-level parameters, defined initially from 'ks prototype use ...'\n    // Each object below should correspond to a component in the components/ directory\n  },\n}\n",
			"ui.libsonnet":     "{\n  global: {\n    // User-defined global parameters; accessible to all component and environments, Ex:\n    // replicas: 4,\n  },\n  components: {\n    // Component-level parameters, defined initially from 'ks prototype use ...'\n    // Each object below should correspond to a component in the components/ directory\n  },\n}\n",
			"redis.libsonnet":  "{\n  global: {\n    // User-defined global parameters; accessible to all component and environments, Ex:\n    // replicas: 4,\n  },\n  components: {\n    // Component-level parameters, defined initially from 'ks prototype use ...'\n    // Each object below should correspond to a component in the components/ directory\n  },\n}\n",
		}}

	gen, _ := json.Marshal(resp)
	s := string(gen)
	fmt.Fprintf(w, s)
}
