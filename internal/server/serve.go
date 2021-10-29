package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/reidmit/yapp/internal/config"
	"github.com/reidmit/yapp/internal/ytt"
	"gopkg.in/yaml.v2"
)

type handledRoute struct {
	method string
	path   string
	config config.RouteConfig
}

func Serve(appConfig *config.AppConfig, port int64, yttPath string) {
	setUpHandlers(appConfig, yttPath)

	fmt.Printf("Listening on port %v...\n", port)

	http.ListenAndServe(":"+strconv.FormatInt(port, 10), nil)
}

func setUpHandlers(appConfig *config.AppConfig, yttPath string) {
	for _, route := range getHandledRoutes(appConfig.Routes) {
		route := route // lol

		fmt.Printf("Setting up route %s %s...\n", route.method, route.path)

		http.HandleFunc(route.path, func(res http.ResponseWriter, req *http.Request) {
			if req.Method == route.method {
				res.Header().Set("Content-Type", "text/x-yaml")

				fmt.Printf("%s %s\n", req.Method, route.path)

				dataValues, err := generateDataValuesFromRequest(req)
				if err != nil {
					fmt.Printf("error generating data values: %v", err)
					http.Error(res, "uh oh", 500)
					return
				}

				err = ytt.Run(appConfig, dataValues)
				if err != nil {
					fmt.Printf("error running ytt: %v", err)
					http.Error(res, "uh oh", 500)
					return
				}

				// replace route config with new config generated by ytt
				route.config = appConfig.Routes[fmt.Sprintf("%s %s", route.method, route.path)]

				if route.config.Status != nil {
					res.WriteHeader(*route.config.Status)
				}

				if route.config.Body != nil {
					responseBody, err := yaml.Marshal(route.config.Body)
					if err != nil {
						fmt.Printf("error marshalling response body: %v", err)
						http.Error(res, "uh oh", 500)
						return
					}

					res.Write(responseBody)
				}

				return
			}

			res.WriteHeader(404)
		})
	}
}

func generateDataValuesFromRequest(req *http.Request) (map[string]interface{}, error) {
	reqBodyBytes, _ := ioutil.ReadAll(req.Body)
	reqBody := make(map[string]interface{})
	err := yaml.Unmarshal(reqBodyBytes, reqBody)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"request": map[string]interface{}{
			"body":    reqBody,
			"headers": req.Header,
		},
	}, nil
}

func getHandledRoutes(routes map[string]config.RouteConfig) []handledRoute {
	var handledRoutes []handledRoute

	for routeWithMethod, routeConfig := range routes {
		parts := strings.Split(routeWithMethod, " ")

		handledRoutes = append(handledRoutes, handledRoute{
			method: parts[0],
			path:   parts[1],
			config: routeConfig,
		})
	}

	return handledRoutes
}
