helm push -f qis-crawler/qis-crawler/ dev &&
helm push -f qis-crawler-queue-api/qis-crawler-queue-api/ dev &&
helm push -f qis-index-ad-api/qis-index-ad-api/ dev &&
helm push -f qis-index-ad-submission-api/qis-index-ad-submission-api/ dev &&
helm push -f qis-index-search-api/qis-index-search-api/ dev &&
helm push -f redis-queue/ dev &&
helm push -f redis-search/ dev