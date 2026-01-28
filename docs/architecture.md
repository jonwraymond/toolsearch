# Architecture

`toolsearch` is a pluggable search strategy for `toolindex`.

```mermaid
flowchart LR
  A[toolindex.Search] --> B[toolsearch.BM25Searcher]
  B --> C[Bleve index]
  C --> D[Ranked summaries]
```

## Determinism

- Docs are sorted by ID before indexing
- Tie-breaking uses score DESC, then ID ASC
- Fingerprints ensure stable caching
