# Copyright 2021 Changkun Ou. All rights reserved.
# Use of this source code is governed by a MIT
# license that can be found in the LICENSE file.

# cat:cs.GR+OR+cat:cs.CG+OR
python fetch_papers.py --search-query cat:cs.HC+OR+cat:cs.AI+OR+cat:stat.ML
python download_pdfs.py
python parse_pdf_to_text.py
python thumb_pdf.py
python analyze.py
python buildsvm.py
python make_cache.py
docker-compose restart arxiv