# Quiet Hacker News Clone

A clone of a simple website using the Hacker News API. It presents two main features:

- Effective pullinf of API using parallel coroutines.
- Two caching strategy: refresh on demand and periodic refresh of cache in the backgorund

# Usage:

```
$ ./hackernews -h
Usage: hackernews [--cache CACHE] [--entries ENTRIES] [--port PORT] [--period PERIOD]

Options:
  --cache CACHE          Kind of cache to use. Options: 'none', 'refresh', 'background' [default: none]
  --entries ENTRIES, -e ENTRIES
                         Number of entries shown on the main page [default: 30]
  --port PORT            Port to launch the server on [default: 3000]
  --period PERIOD, -p PERIOD
                         Sets period for cache refresh [default: 30s]
  --help, -h             display this help and exit
```

A typical usage is:

```
./hackernews
```

And then navigate to the [http://localhost:3000](http://localhost:3000) address to see the page and check page loading time (at the bottom).