# Design Notes

This page documents the tradeoffs and error semantics behind `toolsearch`.

## Design tradeoffs

- **BM25 via Bleve.** `toolsearch` uses Bleve's BM25 implementation for a strong lexical baseline without bringing in a full search stack.
- **Field weighting by duplication.** Instead of multi-field indexing, the searcher builds a single weighted document by duplicating name/namespace/tag tokens. This keeps the index schema simple and predictable.
- **Deterministic ordering.** Documents are sorted by tool ID before indexing. Ties are broken by tool ID to avoid nondeterministic results.
- **In-memory index.** Uses Bleve's in-memory index for speed and simplicity; this trades persistence for low overhead.
- **Fingerprint-based rebuild.** Index rebuilds only when the tool set changes (based on a fingerprint), reducing overhead for repeated searches.

## Error semantics

- BM25 search returns standard `error` values from Bleve.
- `Search` returns empty slices on empty queries or no docs; it does not treat these as errors.
- `Close` frees index resources; callers should treat errors from `Close` as operational warnings.

## Extension points

- **Custom weights:** configure `NameBoost`, `NamespaceBoost`, and `TagsBoost`.
- **Safety caps:** limit search surface with `MaxDocs` or `MaxDocTextLen`.
- **Alternative engines:** implement `toolindex.Searcher` to swap BM25 out for semantic search later.

## Operational guidance

- Start with BM25 if lexical search quality matters; switch to semantic only if needed.
- Keep `MaxDocTextLen` modest to avoid oversized indices from long descriptions.
- Use deterministic doc ordering to keep search results stable across deploys.
