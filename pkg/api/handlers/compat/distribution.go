package compat

import (
	"github.com/containers/podman/v4/pkg/api/handlers/utils"
	"net/http"
	"strings"
)

func DistributionInspect(w http.ResponseWriter, r *http.Request) {
	image := utils.GetName(r)
	if strings.HasPrefix(image, "sha26:") {

	}
}
