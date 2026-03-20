package grpc

type targetIdRequest interface {
	GetOwnerId() string
	GetElectricityId() string
}

func extractTargetId(req targetIdRequest) (ownerId string, electricityId string) {
	return req.GetOwnerId(), req.GetElectricityId()
}
