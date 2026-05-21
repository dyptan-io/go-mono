package storage

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type testItem struct {
	id string
}

func (i testItem) ID() ID { return ID(i.id) }

func TestInMemory(t *testing.T) {
	t.Parallel()

	t.Run("Insert", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			giveItems []testItem
			wantItems []testItem
			wantErr   error
		}{
			"single record": {
				giveItems: []testItem{{id: "test"}},
				wantItems: []testItem{{id: "test"}},
			},
			"multiple records": {
				giveItems: []testItem{{id: "test-1"}, {id: "test-2"}},
				wantItems: []testItem{{id: "test-1"}, {id: "test-2"}},
			},
			"error on missing ID": {
				giveItems: []testItem{{id: ""}},
				wantErr:   ErrMissingID,
			},
		}

		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				store := NewInMemory[testItem]()

				var insertErr error
				for _, item := range test.giveItems {
					insertErr = store.Insert(item)
				}

				if test.wantErr != nil {
					require.ErrorIs(t, insertErr, test.wantErr)
					return
				}

				require.NoError(t, insertErr)

				items, err := store.Find(func(_ testItem) bool { return true })
				require.NoError(t, err)
				require.ElementsMatch(t, test.wantItems, items)
			})
		}
	})

	t.Run("Get", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			giveItems []testItem
			giveID    ID
			wantItem  testItem
			wantErr   error
		}{
			"single record": {
				giveItems: []testItem{{id: "test"}},
				giveID:    "test",
				wantItem:  testItem{id: "test"},
			},
			"correct record from many": {
				giveItems: []testItem{{id: "test-1"}, {id: "test-2"}},
				giveID:    "test-2",
				wantItem:  testItem{id: "test-2"},
			},
			"error on not found": {
				giveItems: []testItem{{id: "test-1"}, {id: "test-2"}},
				giveID:    "does-not-exist",
				wantErr:   ErrNotFound,
			},
		}

		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				store := NewInMemory[testItem]()

				for _, item := range test.giveItems {
					require.NoError(t, store.Insert(item))
				}

				item, err := store.Get(test.giveID)

				if test.wantErr != nil {
					require.ErrorIs(t, err, test.wantErr)
					return
				}

				require.NoError(t, err)
				require.Equal(t, test.wantItem, item)
			})
		}
	})

	t.Run("Find", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			giveItems   []testItem
			giveMatcher Matcher[testItem]
			wantItems   []testItem
		}{
			"all records": {
				giveItems:   []testItem{{id: "test-1"}, {id: "test-2"}},
				giveMatcher: func(_ testItem) bool { return true },
				wantItems:   []testItem{{id: "test-1"}, {id: "test-2"}},
			},
			"filtered by predicate": {
				giveItems:   []testItem{{id: "test-1"}, {id: "test-2"}, {id: "test-3"}},
				giveMatcher: func(v testItem) bool { return strings.HasSuffix(v.id, "2") },
				wantItems:   []testItem{{id: "test-2"}},
			},
			"empty result": {
				giveItems:   []testItem{{id: "test-1"}, {id: "test-2"}},
				giveMatcher: func(_ testItem) bool { return false },
				wantItems:   nil,
			},
		}

		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				store := NewInMemory[testItem]()

				for _, item := range test.giveItems {
					require.NoError(t, store.Insert(item))
				}

				items, err := store.Find(test.giveMatcher)
				require.NoError(t, err)
				require.ElementsMatch(t, test.wantItems, items)
			})
		}
	})

	t.Run("Delete", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			giveItems  []testItem
			giveID     ID
			wantRemain []testItem
			wantErr    error
		}{
			"removes existing record": {
				giveItems:  []testItem{{id: "test-1"}, {id: "test-2"}},
				giveID:     "test-1",
				wantRemain: []testItem{{id: "test-2"}},
			},
			"error on not found": {
				giveItems: []testItem{{id: "test-1"}},
				giveID:    "does-not-exist",
				wantErr:   ErrNotFound,
			},
			"error on missing ID": {
				giveItems: []testItem{{id: "test-1"}},
				giveID:    "",
				wantErr:   ErrMissingID,
			},
		}

		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				store := NewInMemory[testItem]()

				for _, item := range test.giveItems {
					require.NoError(t, store.Insert(item))
				}

				err := store.Delete(test.giveID)

				if test.wantErr != nil {
					require.ErrorIs(t, err, test.wantErr)
					return
				}

				require.NoError(t, err)

				items, err := store.Find(func(_ testItem) bool { return true })
				require.NoError(t, err)
				require.ElementsMatch(t, test.wantRemain, items)
			})
		}
	})
}
