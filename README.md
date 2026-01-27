# toolsearch

`toolsearch` provides optional, higher-quality search strategies for
`toolindex`. It keeps heavier ranking/indexing dependencies out of the core
registry while still plugging into the same `toolindex.Searcher` interface.

Today it ships a BM25 searcher backed by Bleve.

## Install

```bash
go get github.com/jonwraymond/toolsearch
```

## Quick start (inject into toolindex)

```go
import (
  "github.com/jonwraymond/toolindex"
  "github.com/jonwraymond/toolsearch"
)

searcher := toolsearch.NewBM25Searcher(toolsearch.BM25Config{})

idx := toolindex.NewInMemoryIndex(toolindex.IndexOptions{
  Searcher: searcher,
})
```

See runnable examples in:
- `toolsearch/example/basic/main.go`
- `toolsearch/example/custom_config/main.go`
- `toolsearch/example/toolindex_integration/main.go`

## BM25 behavior and safety

- Deterministic:
  - documents are sorted by ID before fingerprinting and `MaxDocs` limiting
  - tie-breaking is score DESC, then ID ASC
- Efficient:
  - the Bleve index is cached and rebuilt only when the doc fingerprint changes
- Bounded:
  - `MaxDocs` limits indexed documents
  - `MaxDocTextLen` truncates very long descriptions

## Version compatibility

See `VERSIONS.md` for the authoritative, auto-generated compatibility matrix.

