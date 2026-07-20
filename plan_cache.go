package graphql

import (
	"container/list"
	"sync"
	"sync/atomic"

	"github.com/graphql-go/graphql/gqlerrors"
)

type PlanCache struct {
	opts PlanCacheOptions

	mu      sync.Mutex
	entries map[string]*list.Element
	order   *list.List

	hits   atomic.Uint64
	misses atomic.Uint64
}

type PlanCacheOptions struct {
	MaxEntries    int
	MaxQueryBytes int
	Normalize     bool
}

const (
	defaultPlanCacheMaxEntries    = 1024
	defaultPlanCacheMaxQueryBytes = 64 * 1024
)

type PlanResult struct {
	Plan      *Plan
	SynthArgs map[string]interface{}
	Errors    []gqlerrors.FormattedError
}

type planCacheEntry struct {
	schema *Schema
	result PlanResult
}

type planCacheItem struct {
	key string
	e   *planCacheEntry
}

func NewPlanCache(opts PlanCacheOptions) *PlanCache { _ = "STUB: not implemented"; return nil }

func (c *PlanCache) Get(schema *Schema, query, operationName string) PlanResult {
	_ = "STUB: not implemented"
	return *new(PlanResult)
}

func (c *PlanCache) HitsMisses() (hits, misses uint64) { _ = "STUB: not implemented"; return 0, 0 }

func (c *PlanCache) Reset() { _ = "STUB: not implemented"; return }

func (c *PlanCache) shouldCache(querySize int) bool { _ = "STUB: not implemented"; return false }

func (c *PlanCache) lookup(schema *Schema, key string) (PlanResult, bool) {
	_ = "STUB: not implemented"
	return *new(PlanResult), false
}

func (c *PlanCache) store(schema *Schema, key string, pr PlanResult) {
	_ = "STUB: not implemented"
	return
}

func planAndValidate(schema *Schema, query, operationName string) PlanResult {
	_ = "STUB: not implemented"
	return *new(PlanResult)
}
