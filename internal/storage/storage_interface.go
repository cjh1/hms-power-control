// MIT License
//
// (C) Copyright [2022] Hewlett Packard Enterprise Development LP
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
// THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR
// OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package storage

import (
	"github.com/Cray-HPE/hms-power-control/internal/model"
	"github.com/sirupsen/logrus"
	"time"
)

type StorageProvider interface {
	Init(Logger *logrus.Logger) error
	Ping() error

	StorePowerStatus(p model.PowerStatusComponent) error
	DeletePowerStatus(xname string) error
	GetPowerStatus(xname string) (model.PowerStatusComponent, error)
	GetAllPowerStatus() (model.PowerStatus, error)
}

type DistributedLockProvider interface {
	Init(Logger *logrus.Logger) error
	InitFromStorage(si interface{}, Logger *logrus.Logger)
	Ping() error
	GetDuration() time.Duration
	DistributedTimedLock(maxLockTime time.Duration) error
	Unlock() error
}
