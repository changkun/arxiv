# Copyright 2021 Changkun Ou. All rights reserved.
# Use of this source code is governed by a MIT
# license that can be found in the LICENSE file.

FROM python:3.9.0
WORKDIR /arxiv
COPY . .
RUN pip install -r requirements.txt
CMD ["python", "./serve.py", "--prod", "--database", "mongo:27017"]