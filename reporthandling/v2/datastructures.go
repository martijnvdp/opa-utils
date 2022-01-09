package v2

import (
	"time"

	"github.com/armosec/opa-utils/reporthandling"
	"github.com/armosec/opa-utils/reporthandling/results/v1/reportsummary"
	"github.com/armosec/opa-utils/reporthandling/results/v1/resourcesresults"

	"k8s.io/apimachinery/pkg/version"
)

// PostureReport posture scanning report structure
type PostureReport struct {
	Attributes           []reportsummary.PostureAttributes `json:"attributes"` //allow flexible properties for posture reports
	CustomerGUID         string                            `json:"customerGUID"`
	ClusterName          string                            `json:"clusterName"`
	ClusterCloudProvider string                            `json:"clusterCloudProvider"`
	ReportID             string                            `json:"reportGUID"`
	JobID                string                            `json:"jobID"`
	ClusterAPIServerInfo *version.Info                     `json:"clusterAPIServerInfo"`
	ReportGenerationTime time.Time                         `json:"generationTime"`
	SummaryDetails       reportsummary.SummaryDetails      `json:"summaryDetails,omitempty"` // Developing
	Results              []resourcesresults.Result         `json:"results,omitempty"`        // Developing
	Resources            []reporthandling.Resource         `json:"resources,omitempty"`
}
