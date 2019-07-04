# Copyright 2019 Authors of project-gadgets
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

################################################################################
# Include definitions
################################################################################
include ../../Makefile.defs

################################################################################
# Target definition
################################################################################
TARGET := go-check-license

################################################################################
# Build target
################################################################################
.PHONY: build
build:
	$(GO_BUILD) -o $(BUILD_DIR)/$(TARGET) .

################################################################################
# Build Static target
################################################################################
.PHONY: build-static
build-static:
	$(GO_BUILD_STATIC) -o $(BUILD_DIR)/$(TARGET) .

################################################################################
# Install target
################################################################################
.PHONY: install
install:
	$(GO_INSTALL) .

################################################################################
# Clean target
################################################################################
.PHONY: clean
clean:
	$(GO_CLEAN) .