/*
 * Copyright (c) 2022, NVIDIA CORPORATION.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package nvlib

import (
	"gitlab.com/nvidia/cloud-native/go-nvlib/pkg/nvmdev"
	"gitlab.com/nvidia/cloud-native/go-nvlib/pkg/nvpci"
)

// Interface provides the API to the nvlib package
type Interface struct {
	Nvpci  nvpci.Interface
	Nvmdev nvmdev.Interface
}

// New creates a new instance of the 'nvlib' interface
func New() Interface {
	return Interface{
		Nvpci:  nvpci.New(),
		Nvmdev: nvmdev.New(),
	}
}
