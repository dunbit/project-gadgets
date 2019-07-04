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
include Makefile.defs

################################################################################
# All target
################################################################################
.PHONY: all
all: build

################################################################################
# Download target
################################################################################
.PHONY: download
download:
	@ $(GO_MOD) download
	@ $(GO_GET) -u github.com/onsi/ginkgo/ginkgo

################################################################################
# Build target
################################################################################
.PHONY: build
build: $(SUBDIR_ALL)

.PHONY: $(SUBDIR_ALL)
$(SUBDIR_ALL):
	@ $(MAKE) -C $(CMD_DIR)/$@ build

################################################################################
# Test target
################################################################################
.PHONY: test
test:
	$(GINKGO_TEST)

################################################################################
# Test Watch target
################################################################################
.PHONY: test-watch
test-watch:
	$(GINKGO_TEST_WATCH)

################################################################################
# Test Cover target
################################################################################
.PHONY: test-cover
test-cover:
	mkdir -p $(COVER_DIR)
	$(GINKGO_TEST_COVER)

################################################################################
# Install target
################################################################################
.PHONY: install
install:
	@ for i in $(SUBDIR_ALL); do $(MAKE) -C $(CMD_DIR)/$$i install; done

################################################################################
# Clean target
################################################################################
.PHONY: clean
clean:
	rm -Rf $(BUILD_DIR)
	rm -Rf $(COVER_DIR)
	@ find $(ROOT_DIR) -name *.coverprofile -delete
	@ for i in $(SUBDIR_ALL); do $(MAKE) -C $(CMD_DIR)/$$i clean; done