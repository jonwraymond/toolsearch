# Architecture

`toolsearch` is a pluggable search strategy for `toolindex`.

## Indexing flow


![Diagram](assets/diagrams/indexing-flow.svg)


## Search sequence


![Diagram](assets/diagrams/indexing-flow.svg)


## Determinism

- Docs are sorted by ID before indexing
- Tie-breaking uses score DESC, then ID ASC
- Fingerprints ensure stable caching
