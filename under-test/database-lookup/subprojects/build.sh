#!/bin/bash
docker build qis-crawler -t "$1"/qis-crawler &&
docker build qis-crawler-queue-api -t "$1"/qis-crawler-queue-api &&
docker build qis-index-ad-api -t "$1"/qis-index-ad-api &&
docker build qis-index-ad-submission-api -t "$1"/qis-index-ad-submission-api &&
docker build qis-index-search-api -t "$1"/qis-index-search-api
