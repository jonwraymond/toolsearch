# Usage

## Configure BM25

```go
searcher := toolsearch.NewBM25Searcher(toolsearch.BM25Config{
  NameBoost:      3,
  NamespaceBoost: 2,
  TagsBoost:      2,
  MaxDocs:        5000,
  MaxDocTextLen:  2000,
})
```

## Inject into toolindex

```go
idx := toolindex.NewInMemoryIndex(toolindex.IndexOptions{Searcher: searcher})
```

## Safety controls

- `MaxDocs`: cap indexed documents
- `MaxDocTextLen`: truncate long descriptions
