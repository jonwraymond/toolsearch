package toolsearch

import (
	"testing"

	"github.com/jonwraymond/toolindex"
)

func TestFingerprint_SameDocsProduceSameFingerprint(t *testing.T) {
	docs := []toolindex.SearchDoc{
		{
			ID:      "tool-1",
			DocText: "description one",
			Summary: toolindex.Summary{ID: "tool-1", Name: "Tool1", Namespace: "ns1"},
		},
		{
			ID:      "tool-2",
			DocText: "description two",
			Summary: toolindex.Summary{ID: "tool-2", Name: "Tool2", Namespace: "ns2"},
		},
	}

	fp1 := computeFingerprint(docs)
	fp2 := computeFingerprint(docs)

	if fp1 != fp2 {
		t.Errorf("same docs produced different fingerprints: %s vs %s", fp1, fp2)
	}
	if fp1 == "" {
		t.Error("fingerprint is empty")
	}
}

func TestFingerprint_DifferentDocsProduceDifferentFingerprint(t *testing.T) {
	docs1 := []toolindex.SearchDoc{
		{ID: "tool-1", DocText: "description one"},
	}
	docs2 := []toolindex.SearchDoc{
		{ID: "tool-2", DocText: "description two"},
	}

	fp1 := computeFingerprint(docs1)
	fp2 := computeFingerprint(docs2)

	if fp1 == fp2 {
		t.Error("different docs produced same fingerprint")
	}
}

func TestFingerprint_OrderMatters(t *testing.T) {
	doc1 := toolindex.SearchDoc{ID: "tool-1", DocText: "one"}
	doc2 := toolindex.SearchDoc{ID: "tool-2", DocText: "two"}

	fp1 := computeFingerprint([]toolindex.SearchDoc{doc1, doc2})
	fp2 := computeFingerprint([]toolindex.SearchDoc{doc2, doc1})

	if fp1 == fp2 {
		t.Error("different order should produce different fingerprints")
	}
}

func TestFingerprint_IncludesAllFields(t *testing.T) {
	base := toolindex.SearchDoc{
		ID:      "tool-1",
		DocText: "description",
		Summary: toolindex.Summary{
			ID:               "tool-1",
			Name:             "Tool1",
			Namespace:        "ns1",
			ShortDescription: "short desc",
			Tags:             []string{"tag1", "tag2"},
		},
	}

	// Each variation should produce a different fingerprint
	variations := []toolindex.SearchDoc{
		{ID: "tool-1-changed", DocText: base.DocText, Summary: base.Summary},
		{ID: base.ID, DocText: "changed", Summary: base.Summary},
		{
			ID:      base.ID,
			DocText: base.DocText,
			Summary: toolindex.Summary{
				ID:               base.Summary.ID,
				Name:             "ChangedName",
				Namespace:        base.Summary.Namespace,
				ShortDescription: base.Summary.ShortDescription,
				Tags:             base.Summary.Tags,
			},
		},
		{
			ID:      base.ID,
			DocText: base.DocText,
			Summary: toolindex.Summary{
				ID:               base.Summary.ID,
				Name:             base.Summary.Name,
				Namespace:        "changed-ns",
				ShortDescription: base.Summary.ShortDescription,
				Tags:             base.Summary.Tags,
			},
		},
		{
			ID:      base.ID,
			DocText: base.DocText,
			Summary: toolindex.Summary{
				ID:               base.Summary.ID,
				Name:             base.Summary.Name,
				Namespace:        base.Summary.Namespace,
				ShortDescription: "changed short desc",
				Tags:             base.Summary.Tags,
			},
		},
		{
			ID:      base.ID,
			DocText: base.DocText,
			Summary: toolindex.Summary{
				ID:               base.Summary.ID,
				Name:             base.Summary.Name,
				Namespace:        base.Summary.Namespace,
				ShortDescription: base.Summary.ShortDescription,
				Tags:             []string{"different-tag"},
			},
		},
	}

	baseFP := computeFingerprint([]toolindex.SearchDoc{base})

	for i, v := range variations {
		vFP := computeFingerprint([]toolindex.SearchDoc{v})
		if vFP == baseFP {
			t.Errorf("variation %d should produce different fingerprint from base", i)
		}
	}
}

func TestFingerprint_TagOrderIndependent(t *testing.T) {
	// Same tags in different orders should produce same fingerprint
	doc1 := toolindex.SearchDoc{
		ID:      "tool-1",
		DocText: "description",
		Summary: toolindex.Summary{
			ID:   "tool-1",
			Name: "Tool1",
			Tags: []string{"alpha", "bravo", "charlie"},
		},
	}
	doc2 := toolindex.SearchDoc{
		ID:      "tool-1",
		DocText: "description",
		Summary: toolindex.Summary{
			ID:   "tool-1",
			Name: "Tool1",
			Tags: []string{"charlie", "alpha", "bravo"},
		},
	}

	fp1 := computeFingerprint([]toolindex.SearchDoc{doc1})
	fp2 := computeFingerprint([]toolindex.SearchDoc{doc2})

	if fp1 != fp2 {
		t.Errorf("same tags in different order should produce same fingerprint: %s vs %s", fp1, fp2)
	}
}

func TestFingerprint_EmptyDocs(t *testing.T) {
	var docs []toolindex.SearchDoc
	fp := computeFingerprint(docs)

	// Should return a consistent fingerprint for empty docs
	fp2 := computeFingerprint(nil)
	if fp != fp2 {
		t.Error("empty slice and nil should produce same fingerprint")
	}
	if fp == "" {
		t.Error("fingerprint should not be empty for empty docs")
	}
}
