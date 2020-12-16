#!/bin/bash
docker push "$1"/qis-crawler &&
docker push "$1"/qis-crawler-queue-api &&
docker push "$1"/qis-index-ad-api &&
docker push "$1"/qis-index-ad-submission-api &&
docker push "$1"/qis-index-search-api
