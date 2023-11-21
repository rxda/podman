package entities

import v1 "github.com/opencontainers/image-spec/specs-go/v1"

// DistributionInspect describes the result obtained from contacting the
// registry to retrieve image metadata
type DistributionInspect struct {
	// Descriptor contains information about the manifest, including
	// the content addressable digest
	Descriptor v1.Descriptor
	// Platforms contains the list of platforms supported by the image,
	// obtained by parsing the manifest
	Platforms []v1.Platform
}
