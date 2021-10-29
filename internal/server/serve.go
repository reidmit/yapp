package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/reidmit/yapp/internal/config"
	"github.com/reidmit/yapp/internal/ytt"
	"gopkg.in/yaml.v2"
)

func Serve(appConfig *config.AppConfig) {
	setUpHandlers(appConfig)

	log.Printf("Listening on port %v\n", appConfig.Port)

	http.ListenAndServe(":"+strconv.FormatInt(appConfig.Port, 10), nil)
}

func setUpHandlers(appConfig *config.AppConfig) {
	for _, route := range config.GetHandledRoutes(appConfig.Routes) {
		route := route // lol

		log.Printf("Setting up route %s %s\n", route.Method, route.Path)

		http.HandleFunc(route.Path, func(res http.ResponseWriter, req *http.Request) {
			if req.Method == route.Method {
				log.Printf("%s %s\n", route.Method, route.Path)

				dataValues, err := generateDataValuesFromRequest(req)
				if err != nil {
					log.Printf("Error generating data values from request: %v", err)
					http.Error(res, "uh oh", 500)
					return
				}

				newRouteConfig, err := ytt.Run(appConfig, route, dataValues)
				if err != nil {
					log.Printf("Error running ytt: %v", err)
					http.Error(res, "uh oh", 500)
					return
				}

				if newRouteConfig.Status != nil {
					res.WriteHeader(*newRouteConfig.Status)
				}

				fmt.Printf("rendered: %+v\n", newRouteConfig)

				if newRouteConfig.Body != nil {
					responseBody, err := yaml.Marshal(newRouteConfig.Body)
					if err != nil {
						log.Printf("Error marshalling response body: %v", err)
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
			"query":   req.URL.Query(),
		},
	}, nil
}
