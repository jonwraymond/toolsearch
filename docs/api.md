# API Reference

## BM25Config

```go
type BM25Config struct {
  NameBoost      int
  NamespaceBoost int
  TagsBoost      int
  MaxDocs        int
  MaxDocTextLen  int
}
```

## BM25Searcher

```go
type BM25Searcher struct {}

func NewBM25Searcher(cfg BM25Config) *BM25Searcher

// implements toolindex.Searcher
func (s *BM25Searcher) Search(query string, limit int, docs []toolindex.SearchDoc) ([]toolindex.Summary, error)
```
