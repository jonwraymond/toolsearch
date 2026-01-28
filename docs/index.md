# toolsearch

Search strategy library for toolindex.

## What this repo provides

- BM25 ranking
- Pluggable search strategies

## Example

```go
engine := toolsearch.NewBM25Engine(cfg)
```
