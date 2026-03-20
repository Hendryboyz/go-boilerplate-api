package helpers

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	wire.Bind(new(Clock), new(RealClock)),
	NewRealClock,
)
