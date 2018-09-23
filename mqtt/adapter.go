//
// Copyright (c) 2018
// Mainflux
//
// SPDX-License-Identifier: Apache-2.0
//

// Package mqtt contains the domain concept definitions needed to support
// Mainflux mqtt adapter service functionality.
package mqtt

import "github.com/mainflux/mainflux"

var _ mainflux.MessagePublisher = (*adapterService)(nil)

type adapterService struct {
	pub mainflux.MessagePublisher
}

// New instantiates the MQTT adapter implementation.
func New(pub mainflux.MessagePublisher) mainflux.MessagePublisher {
	return &adapterService{pub}
}

func (as *adapterService) Publish(msg mainflux.RawMessage) error {
	return as.pub.Publish(msg)
}
