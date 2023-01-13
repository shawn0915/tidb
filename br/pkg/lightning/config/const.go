// Copyright 2019 PingCAP, Inc.
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

package config

import (
	"github.com/docker/go-units"
)

const (
	// mydumper
	ReadBlockSize ByteSize = 64 * units.KiB
	IOBufferSize  ByteSize = 4 * units.MiB
	MaxRegionSize ByteSize = 256 * units.MiB
	// See: https://github.com/tikv/tikv/blob/e030a0aae9622f3774df89c62f21b2171a72a69e/etc/config-template.toml#L360
	// lower the max-key-count to avoid tikv trigger region auto split
	SplitRegionSize         ByteSize = 96 * units.MiB
	SplitRegionKeys         int      = 1_280_000
	MaxSplitRegionSizeRatio int      = 10

	BufferSizeScale = 5

	defaultMaxAllowedPacket = 64 * units.MiB

	DefaultBatchSize ByteSize = 100 * units.GiB
)
