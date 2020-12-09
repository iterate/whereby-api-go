package wherebyiface

import "github.com/iterate/whereby-api-go"

// Interface guard. This will fail at compile time if the interface is no longer
// implemented by the standard client.
var _ Client = (*whereby.Client)(nil)
