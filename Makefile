# Copyright 2021 Changkun Ou. All rights reserved.
# Use of this source code is governed by a MIT
# license that can be found in the LICENSE file.

up:
	docker-compose up -d
down:
	docker-compose down
update:
	sh update.sh
build:
	docker build -t arxiv-preserver:latest .
