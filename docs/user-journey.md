# User Journey

This journey shows how `toolsearch` slots into the end-to-end workflow as a pluggable search engine.

## End-to-end flow (stack view)

![Diagram](assets/diagrams/user-journey.svg)

## Step-by-step

1. **Enable BM25** in `metatools-mcp` using the `toolsearch` build tag.
2. **Set env vars** (e.g., `METATOOLS_SEARCH_STRATEGY=bm25`).
3. **Search requests** now flow through the BM25 searcher.

## Example: configure BM25 via env

```bash
# build with toolsearch support
GOFLAGS="-tags=toolsearch" go build ./cmd/metatools

# choose BM25 strategy at runtime
export METATOOLS_SEARCH_STRATEGY=bm25
export METATOOLS_SEARCH_BM25_NAME_BOOST=4
export METATOOLS_SEARCH_BM25_TAGS_BOOST=2
```

## Expected outcomes

- Higher-quality lexical ranking for tool discovery.
- Deterministic ordering and tie-breaking.
- No API changes required in `toolindex` or `metatools-mcp`.

## Common failure modes

- Build without the `toolsearch` tag and request `bm25` strategy (fails fast).
- Oversized tool descriptions if `MaxDocTextLen` is not capped.
