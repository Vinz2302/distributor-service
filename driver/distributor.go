package driver

import (
	handler "distributor-service/modules/v1/utilities/distributor/handler"
)

var (
	DistributorHandler = handler.NewDistributorHandler()
)