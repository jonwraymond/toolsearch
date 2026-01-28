# Architecture

`toolsearch` is a pluggable search strategy for `toolindex`.

## Indexing flow

```mermaid
flowchart LR
  A[Search docs] --> B[Fingerprint]
  B --> C[BM25 index]
  C --> D[Ranked summaries]
```

## Search sequence

```mermaid
sequenceDiagram
  participant Index
  participant Searcher
  participant Bleve

  Index->>Searcher: Search(query, docs)
  Searcher->>Searcher: Build index if fingerprint changed
  Searcher->>Bleve: Query
  Bleve-->>Searcher: Hits
  Searcher-->>Index: Summaries
```

## Determinism

- Docs are sorted by ID before indexing
- Tie-breaking uses score DESC, then ID ASC
- Fingerprints ensure stable caching
