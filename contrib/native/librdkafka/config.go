// Copyright 2015-2016 trivago GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package librdkafka

// #cgo CFLAGS: -I/usr/local/include -std=c99
// #cgo LDFLAGS: -L/usr/local/lib -L/usr/local/opt/librdkafka/lib -lrdkafka
// #include "wrapper.h"
import "C"

import (
	"strconv"
)

// Config is a wrapper for rd_kafka_conf_t
type Config struct {
	handle *C.rd_kafka_conf_t
}

// NewConfig creates a new main config wrapper.
// Make sure to call Close() when done.
func NewConfig() Config {
	return Config{
		handle: C.rd_kafka_conf_new(),
	}
}

// Close frees up the native handle
func (c *Config) Close() {
	C.rd_kafka_conf_destroy(c.handle)
}

// Set sets a string value in this config
func (c *Config) Set(key, value string) {
	nativeErr := new(ErrorHandle)
	if C.rd_kafka_conf_set(c.handle, C.CString(key), C.CString(value), nativeErr.buffer(), nativeErr.len()) != 0 {
		Log.Print(nativeErr)
	}
}

// SetI sets an integer value in this config
func (c *Config) SetI(key string, value int) {
	nativeErr := new(ErrorHandle)
	strValue := strconv.Itoa(value)
	if C.rd_kafka_conf_set(c.handle, C.CString(key), C.CString(strValue), nativeErr.buffer(), nativeErr.len()) != 0 {
		Log.Print(nativeErr)
	}
}

// SetB sets a boolean value in this config
func (c *Config) SetB(key string, value bool) {
	nativeErr := new(ErrorHandle)
	var boolValue string
	if value {
		boolValue = "true"
	} else {
		boolValue = "false"
	}
	if C.rd_kafka_conf_set(c.handle, C.CString(key), C.CString(boolValue), nativeErr.buffer(), nativeErr.len()) != 0 {
		Log.Print(nativeErr)
	}
}
