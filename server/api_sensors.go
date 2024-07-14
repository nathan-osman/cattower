package server

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type apiSensorsOverviewResponse struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func (s *Server) apiSensorsOverview(c *gin.Context) {
	r, err := s.influxdb.Query(
		`SELECT LAST("value") from "temperature" GROUP BY "location"`,
	)
	if err != nil {
		panic(err)
	}
	o := []apiSensorsOverviewResponse{}
	for _, r := range r.Results {
		for _, s := range r.Series {
			v, err := s.Values[0][1].(json.Number).Float64()
			if err != nil {
				panic(err)
			}
			o = append(o, apiSensorsOverviewResponse{
				Name:  s.Tags["location"],
				Value: v,
			})
		}
	}
	c.JSON(http.StatusOK, o)
}
