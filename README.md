# toolsearch

> **DEPRECATED**: This package has been merged into `tooldiscovery/search`.
> Please migrate to `github.com/jonwraymond/tooldiscovery/search`.
>
> See [MIGRATION.md](./MIGRATION.md) for migration instructions.

---

[![Docs](https://img.shields.io/badge/docs-ai--tools--stack-blue)](https://jonwraymond.github.io/ai-tools-stack/)

`toolsearch` provides optional, higher-quality search strategies for
`toolindex`. It keeps heavier ranking/indexing dependencies out of the core
registry while still plugging into the same `toolindex.Searcher` interface.

Today it ships a BM25 searcher backed by Bleve.

## Install

```bash
# DEPRECATED - use tooldiscovery/search instead
go get github.com/jonwraymond/tooldiscovery/search
```

## Migration

This package is deprecated. Please update your imports:

```go
// Old import (deprecated)
import "github.com/jonwraymond/toolsearch"

// New import
import "github.com/jonwraymond/tooldiscovery/search"
```

See [MIGRATION.md](./MIGRATION.md) for complete migration instructions.

## Documentation

- [MIGRATION.md](./MIGRATION.md) — migration guide to tooldiscovery/search
- [tooldiscovery docs](https://jonwraymond.github.io/ai-tools-stack/) — consolidated documentation

## Version compatibility

See `VERSIONS.md` for the authoritative, auto-generated compatibility matrix.
