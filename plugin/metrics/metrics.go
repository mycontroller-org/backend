package metrics

import (
	"time"

	fml "github.com/mycontroller-org/backend/v2/pkg/model/field"
)

// Client interface
type Client interface {
	Close() error
	Ping() error
	Write(field *fml.Field) error
	WriteBlocking(field *fml.Field) error
	Query(queryConfig *QueryConfig) (map[string][]Data, error)
}

// Metrics database types
const (
	TypeInfluxdbV2 = "influxdb_v2"
	TypeVoidDB     = "void_db"
)

// Metric types
const (
	MetricTypeNone       = "none"
	MetricTypeString     = "string"
	MetricTypeCounter    = "counter"
	MetricTypeGauge      = "gauge"
	MetricTypeGaugeFloat = "gauge_float"
	MetricTypeBinary     = "binary"
	MetricTypeGEO        = "geo" // Geo Coordinates or GPS
)

// Metric query input parameters
const (
	QueryKeyName       = "name"
	QueryKeyMetricType = "metric_type"
	QueryKeyStart      = "start"
	QueryKeyStop       = "stop"
	QueryKeyWindow     = "window"
	QueryKeyTags       = "tags"
	QueryKeyFunctions  = "functions"
)

// QueryConfig parameters
type QueryConfig struct {
	Global     Query   `json:"global"`
	Individual []Query `json:"individual"`
}

// Query paramaters
type Query struct {
	Name       string            `json:"name"`
	MetricType string            `json:"metricType"`
	Start      string            `json:"start"`
	Stop       string            `json:"stop"`
	Window     string            `json:"window"`
	Tags       map[string]string `json:"tags"`
	Functions  []string          `json:"functions"`
}

// Data struct
type Data struct {
	Time       time.Time              `json:"timestamp"`
	MetricType string                 `json:"metricType"`
	Metric     map[string]interface{} `json:"metric"`
}

// Clone a query
func (q *Query) Clone() Query {
	tags := make(map[string]string)
	if q.Tags != nil {
		for k, v := range q.Tags {
			tags[k] = v
		}
	}
	functions := []string{}
	if q.Functions != nil {
		functions = q.Functions
	}
	return Query{
		Name:       q.Name,
		MetricType: q.MetricType,
		Start:      q.Start,
		Stop:       q.Stop,
		Window:     q.Window,
		Tags:       tags,
		Functions:  functions,
	}
}

// Merge data from another query
func (q *Query) Merge(new *Query) {
	if new != nil {
		// update default values
		if q.Tags == nil {
			q.Tags = make(map[string]string)
		}
		if q.Functions == nil {
			q.Functions = []string{}
		}
		// update vales
		if new.Name != "" {
			q.Name = new.Name
		}
		if new.MetricType != "" {
			q.MetricType = new.MetricType
		}
		if new.Start != "" {
			q.Start = new.Start
		}
		if new.Stop != "" {
			q.Stop = new.Stop
		}
		if new.Window != "" {
			q.Window = new.Window
		}
		if len(new.Tags) > 0 {
			for k, v := range new.Tags {
				q.Tags[k] = v
			}
		}
		if len(new.Functions) > 0 {
			for _, newFn := range new.Functions {
				found := false
				for _, orgFn := range q.Functions {
					if orgFn == newFn {
						found = true
						break
					}
				}
				if !found {
					q.Functions = append(q.Functions, newFn)
				}
			}
		}
	}
}
