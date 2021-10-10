
# changkun's arXiv preserver

This project is derived from the great [arxiv-sanity](www.arxiv-sanity.com), a web interface that attempts to tame the overwhelming flood of papers on Arxiv. It allows researchers to keep track of recent papers, search for papers, sort papers by similarity to any paper, see recent popular papers, to add papers to a personal library, and to get personalized recommendations of (new or old) Arxiv papers. This code is currently running live at [arxiv.changkun.de/](https://arxiv.changkun.de/), where it's serving Arxiv papers from **Human-Computer Interaction (cs.HC), Computer Graphics (cs.GR), and Computational Geometry (cs.CG)**. With this code base you could replicate the website to any of your favorite subsets of Arxiv by simply changing the categories in `fetch_papers.py`.


## Dependencies

- [ImageMagick](http://www.imagemagick.org/script/index.php) and [pdftotext](https://poppler.freedesktop.org/)
  + `sudo apt-get install imagemagick poppler-utils`
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## First Setup (Processing pipeline)

The processing pipeline requires you to run a series of scripts, and at this stage I really encourage you to manually inspect each script, as they may contain various inline settings you might want to change. In order, the processing pipeline is:

1. Run `fetch_papers.py` to query arxiv API and create a file `data/db/db.p` that contains all information for each paper. This script is where you would modify the **query**, indicating which parts of arxiv you'd like to use. Note that if you're trying to pull too many papers arxiv will start to rate limit you. You may have to run the script multiple times, and I recommend using the arg `--start-index` to restart where you left off when you were last interrupted by arxiv.
2. Run `download_pdfs.py`, which iterates over all papers in parsed pickle and downloads the papers into folder `pdf`
3. Run `parse_pdf_to_text.py` to export all text from pdfs to files in `txt`
4. Run `thumb_pdf.py` to export thumbnails of all pdfs to `thumb`
5. Run `analyze.py` to compute tfidf vectors for all documents based on bigrams. Saves a `data/db/tfidf.p`, `data/db/tfidf_meta.p` and `data/db/sim_dict.p` pickle files.
6. Run `buildsvm.py` to train SVMs for all users (if any), exports a pickle `data/db/user_sim.p`
7. Run `make_cache.py` for various preprocessing so that server starts faster (and make sure to run `sqlite3 data/db/as.db < schema.sql` if this is the very first time ever you're starting arxiv-sanity, which initializes an empty database).
8. Start the mongodb daemon in the background. Mongodb can be installed by following the instructions here - https://docs.mongodb.com/tutorials/install-mongodb-on-ubuntu/.
  * Start the mongodb server with - `sudo service mongod start`.
  * Verify if the server is running in the background : The last line of /var/log/mongodb/mongod.log file must be -
`[initandlisten] waiting for connections on port <port> `
9. Run the flask server with `serve.py`. Visit localhost:5000 and enjoy sane viewing of papers!

**protip: numpy/BLAS**: The script `analyze.py` does quite a lot of heavy lifting with numpy. I recommend that you carefully set up your numpy to use BLAS (e.g. OpenBLAS), otherwise the computations will take a long time. With ~25,000 papers and ~5000 users the script runs in several hours on my current machine with a BLAS-linked numpy.

Once the local setup is working. Then we can build a docker image
so that the server functionality is working:

```bash
make build
make up
```

### Daily Update

Run the following command will update the website:

```bash
make update
```

Setting up a cron task should be ideal to execute the update command:

```
0 2 * * * cd /media/changkun/ExtensionField1/arxiv-hci-preserver && sh update.sh
```