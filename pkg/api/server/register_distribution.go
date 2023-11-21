package server

import (
	"github.com/containers/podman/v4/pkg/api/handlers/compat"
	"github.com/gorilla/mux"
)

func (s *APIServer) registerDistributionHandlers(r *mux.Router) error {
	// swagger:operation GET /distribution/{name}/json compat DistributionInspect
	// ---
	// tags:
	//  - distribution (compat)
	// summary: Get image information from the registry
	// description: Return image digest and platform information by contacting the registry.
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or id of the image
	// produces:
	// - application/json
	// responses:
	//   200:
	//     $ref: "#/responses/distributionInspectResponse"
	//   401:
	//     $ref: "#/responses/imageNotFound"
	//   500:
	//     $ref: "#/responses/internalError"
	r.HandleFunc(VersionedPath("/distribution/{name}/json"), compat.DistributionInspect)
	// Added non version path to URI to support docker non versioned paths
	r.HandleFunc("/distribution/{name}/json", compat.DistributionInspect)
	return nil
}
