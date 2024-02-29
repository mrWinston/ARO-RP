// Code generated by github.com/jewzaam/go-cosmosdb, DO NOT EDIT.

package cosmosdb

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/ugorji/go/codec"

	pkg "github.com/Azure/ARO-RP/pkg/api"
)

type fakeMaintenanceManifestDocumentTriggerHandler func(context.Context, *pkg.MaintenanceManifestDocument) error
type fakeMaintenanceManifestDocumentQueryHandler func(MaintenanceManifestDocumentClient, *Query, *Options) MaintenanceManifestDocumentRawIterator

var _ MaintenanceManifestDocumentClient = &FakeMaintenanceManifestDocumentClient{}

// NewFakeMaintenanceManifestDocumentClient returns a FakeMaintenanceManifestDocumentClient
func NewFakeMaintenanceManifestDocumentClient(h *codec.JsonHandle) *FakeMaintenanceManifestDocumentClient {
	return &FakeMaintenanceManifestDocumentClient{
		jsonHandle:                   h,
		maintenanceManifestDocuments: make(map[string]*pkg.MaintenanceManifestDocument),
		triggerHandlers:              make(map[string]fakeMaintenanceManifestDocumentTriggerHandler),
		queryHandlers:                make(map[string]fakeMaintenanceManifestDocumentQueryHandler),
	}
}

// FakeMaintenanceManifestDocumentClient is a FakeMaintenanceManifestDocumentClient
type FakeMaintenanceManifestDocumentClient struct {
	lock                         sync.RWMutex
	jsonHandle                   *codec.JsonHandle
	maintenanceManifestDocuments map[string]*pkg.MaintenanceManifestDocument
	triggerHandlers              map[string]fakeMaintenanceManifestDocumentTriggerHandler
	queryHandlers                map[string]fakeMaintenanceManifestDocumentQueryHandler
	sorter                       func([]*pkg.MaintenanceManifestDocument)
	etag                         int

	// returns true if documents conflict
	conflictChecker func(*pkg.MaintenanceManifestDocument, *pkg.MaintenanceManifestDocument) bool

	// err, if not nil, is an error to return when attempting to communicate
	// with this Client
	err error
}

// SetError sets or unsets an error that will be returned on any
// FakeMaintenanceManifestDocumentClient method invocation
func (c *FakeMaintenanceManifestDocumentClient) SetError(err error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.err = err
}

// SetSorter sets or unsets a sorter function which will be used to sort values
// returned by List() for test stability
func (c *FakeMaintenanceManifestDocumentClient) SetSorter(sorter func([]*pkg.MaintenanceManifestDocument)) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.sorter = sorter
}

// SetConflictChecker sets or unsets a function which can be used to validate
// additional unique keys in a MaintenanceManifestDocument
func (c *FakeMaintenanceManifestDocumentClient) SetConflictChecker(conflictChecker func(*pkg.MaintenanceManifestDocument, *pkg.MaintenanceManifestDocument) bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.conflictChecker = conflictChecker
}

// SetTriggerHandler sets or unsets a trigger handler
func (c *FakeMaintenanceManifestDocumentClient) SetTriggerHandler(triggerName string, trigger fakeMaintenanceManifestDocumentTriggerHandler) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.triggerHandlers[triggerName] = trigger
}

// SetQueryHandler sets or unsets a query handler
func (c *FakeMaintenanceManifestDocumentClient) SetQueryHandler(queryName string, query fakeMaintenanceManifestDocumentQueryHandler) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.queryHandlers[queryName] = query
}

func (c *FakeMaintenanceManifestDocumentClient) deepCopy(maintenanceManifestDocument *pkg.MaintenanceManifestDocument) (*pkg.MaintenanceManifestDocument, error) {
	var b []byte
	err := codec.NewEncoderBytes(&b, c.jsonHandle).Encode(maintenanceManifestDocument)
	if err != nil {
		return nil, err
	}

	maintenanceManifestDocument = nil
	err = codec.NewDecoderBytes(b, c.jsonHandle).Decode(&maintenanceManifestDocument)
	if err != nil {
		return nil, err
	}

	return maintenanceManifestDocument, nil
}

func (c *FakeMaintenanceManifestDocumentClient) apply(ctx context.Context, partitionkey string, maintenanceManifestDocument *pkg.MaintenanceManifestDocument, options *Options, isCreate bool) (*pkg.MaintenanceManifestDocument, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.err != nil {
		return nil, c.err
	}

	maintenanceManifestDocument, err := c.deepCopy(maintenanceManifestDocument) // copy now because pretriggers can mutate maintenanceManifestDocument
	if err != nil {
		return nil, err
	}

	if options != nil {
		err := c.processPreTriggers(ctx, maintenanceManifestDocument, options)
		if err != nil {
			return nil, err
		}
	}

	existingMaintenanceManifestDocument, exists := c.maintenanceManifestDocuments[maintenanceManifestDocument.ID]
	if isCreate && exists {
		return nil, &Error{
			StatusCode: http.StatusConflict,
			Message:    "Entity with the specified id already exists in the system",
		}
	}
	if !isCreate {
		if !exists {
			return nil, &Error{StatusCode: http.StatusNotFound}
		}

		if maintenanceManifestDocument.ETag != existingMaintenanceManifestDocument.ETag {
			return nil, &Error{StatusCode: http.StatusPreconditionFailed}
		}
	}

	if c.conflictChecker != nil {
		for _, maintenanceManifestDocumentToCheck := range c.maintenanceManifestDocuments {
			if c.conflictChecker(maintenanceManifestDocumentToCheck, maintenanceManifestDocument) {
				return nil, &Error{
					StatusCode: http.StatusConflict,
					Message:    "Entity with the specified id already exists in the system",
				}
			}
		}
	}

	maintenanceManifestDocument.ETag = fmt.Sprint(c.etag)
	c.etag++

	c.maintenanceManifestDocuments[maintenanceManifestDocument.ID] = maintenanceManifestDocument

	return c.deepCopy(maintenanceManifestDocument)
}

// Create creates a MaintenanceManifestDocument in the database
func (c *FakeMaintenanceManifestDocumentClient) Create(ctx context.Context, partitionkey string, maintenanceManifestDocument *pkg.MaintenanceManifestDocument, options *Options) (*pkg.MaintenanceManifestDocument, error) {
	return c.apply(ctx, partitionkey, maintenanceManifestDocument, options, true)
}

// Replace replaces a MaintenanceManifestDocument in the database
func (c *FakeMaintenanceManifestDocumentClient) Replace(ctx context.Context, partitionkey string, maintenanceManifestDocument *pkg.MaintenanceManifestDocument, options *Options) (*pkg.MaintenanceManifestDocument, error) {
	return c.apply(ctx, partitionkey, maintenanceManifestDocument, options, false)
}

// List returns a MaintenanceManifestDocumentIterator to list all MaintenanceManifestDocuments in the database
func (c *FakeMaintenanceManifestDocumentClient) List(*Options) MaintenanceManifestDocumentIterator {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if c.err != nil {
		return NewFakeMaintenanceManifestDocumentErroringRawIterator(c.err)
	}

	maintenanceManifestDocuments := make([]*pkg.MaintenanceManifestDocument, 0, len(c.maintenanceManifestDocuments))
	for _, maintenanceManifestDocument := range c.maintenanceManifestDocuments {
		maintenanceManifestDocument, err := c.deepCopy(maintenanceManifestDocument)
		if err != nil {
			return NewFakeMaintenanceManifestDocumentErroringRawIterator(err)
		}
		maintenanceManifestDocuments = append(maintenanceManifestDocuments, maintenanceManifestDocument)
	}

	if c.sorter != nil {
		c.sorter(maintenanceManifestDocuments)
	}

	return NewFakeMaintenanceManifestDocumentIterator(maintenanceManifestDocuments, 0)
}

// ListAll lists all MaintenanceManifestDocuments in the database
func (c *FakeMaintenanceManifestDocumentClient) ListAll(ctx context.Context, options *Options) (*pkg.MaintenanceManifestDocuments, error) {
	iter := c.List(options)
	return iter.Next(ctx, -1)
}

// Get gets a MaintenanceManifestDocument from the database
func (c *FakeMaintenanceManifestDocumentClient) Get(ctx context.Context, partitionkey string, id string, options *Options) (*pkg.MaintenanceManifestDocument, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if c.err != nil {
		return nil, c.err
	}

	maintenanceManifestDocument, exists := c.maintenanceManifestDocuments[id]
	if !exists {
		return nil, &Error{StatusCode: http.StatusNotFound}
	}

	return c.deepCopy(maintenanceManifestDocument)
}

// Delete deletes a MaintenanceManifestDocument from the database
func (c *FakeMaintenanceManifestDocumentClient) Delete(ctx context.Context, partitionKey string, maintenanceManifestDocument *pkg.MaintenanceManifestDocument, options *Options) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.err != nil {
		return c.err
	}

	_, exists := c.maintenanceManifestDocuments[maintenanceManifestDocument.ID]
	if !exists {
		return &Error{StatusCode: http.StatusNotFound}
	}

	delete(c.maintenanceManifestDocuments, maintenanceManifestDocument.ID)
	return nil
}

// ChangeFeed is unimplemented
func (c *FakeMaintenanceManifestDocumentClient) ChangeFeed(*Options) MaintenanceManifestDocumentIterator {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if c.err != nil {
		return NewFakeMaintenanceManifestDocumentErroringRawIterator(c.err)
	}

	return NewFakeMaintenanceManifestDocumentErroringRawIterator(ErrNotImplemented)
}

func (c *FakeMaintenanceManifestDocumentClient) processPreTriggers(ctx context.Context, maintenanceManifestDocument *pkg.MaintenanceManifestDocument, options *Options) error {
	for _, triggerName := range options.PreTriggers {
		if triggerHandler := c.triggerHandlers[triggerName]; triggerHandler != nil {
			c.lock.Unlock()
			err := triggerHandler(ctx, maintenanceManifestDocument)
			c.lock.Lock()
			if err != nil {
				return err
			}
		} else {
			return ErrNotImplemented
		}
	}

	return nil
}

// Query calls a query handler to implement database querying
func (c *FakeMaintenanceManifestDocumentClient) Query(name string, query *Query, options *Options) MaintenanceManifestDocumentRawIterator {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if c.err != nil {
		return NewFakeMaintenanceManifestDocumentErroringRawIterator(c.err)
	}

	if queryHandler := c.queryHandlers[query.Query]; queryHandler != nil {
		c.lock.RUnlock()
		i := queryHandler(c, query, options)
		c.lock.RLock()
		return i
	}

	return NewFakeMaintenanceManifestDocumentErroringRawIterator(ErrNotImplemented)
}

// QueryAll calls a query handler to implement database querying
func (c *FakeMaintenanceManifestDocumentClient) QueryAll(ctx context.Context, partitionkey string, query *Query, options *Options) (*pkg.MaintenanceManifestDocuments, error) {
	iter := c.Query("", query, options)
	return iter.Next(ctx, -1)
}

func NewFakeMaintenanceManifestDocumentIterator(maintenanceManifestDocuments []*pkg.MaintenanceManifestDocument, continuation int) MaintenanceManifestDocumentRawIterator {
	return &fakeMaintenanceManifestDocumentIterator{maintenanceManifestDocuments: maintenanceManifestDocuments, continuation: continuation}
}

type fakeMaintenanceManifestDocumentIterator struct {
	maintenanceManifestDocuments []*pkg.MaintenanceManifestDocument
	continuation                 int
	done                         bool
}

func (i *fakeMaintenanceManifestDocumentIterator) NextRaw(ctx context.Context, maxItemCount int, out interface{}) error {
	return ErrNotImplemented
}

func (i *fakeMaintenanceManifestDocumentIterator) Next(ctx context.Context, maxItemCount int) (*pkg.MaintenanceManifestDocuments, error) {
	if i.done {
		return nil, nil
	}

	var maintenanceManifestDocuments []*pkg.MaintenanceManifestDocument
	if maxItemCount == -1 {
		maintenanceManifestDocuments = i.maintenanceManifestDocuments[i.continuation:]
		i.continuation = len(i.maintenanceManifestDocuments)
		i.done = true
	} else {
		max := i.continuation + maxItemCount
		if max > len(i.maintenanceManifestDocuments) {
			max = len(i.maintenanceManifestDocuments)
		}
		maintenanceManifestDocuments = i.maintenanceManifestDocuments[i.continuation:max]
		i.continuation += max
		i.done = i.Continuation() == ""
	}

	return &pkg.MaintenanceManifestDocuments{
		MaintenanceManifestDocuments: maintenanceManifestDocuments,
		Count:                        len(maintenanceManifestDocuments),
	}, nil
}

func (i *fakeMaintenanceManifestDocumentIterator) Continuation() string {
	if i.continuation >= len(i.maintenanceManifestDocuments) {
		return ""
	}
	return fmt.Sprintf("%d", i.continuation)
}

// NewFakeMaintenanceManifestDocumentErroringRawIterator returns a MaintenanceManifestDocumentRawIterator which
// whose methods return the given error
func NewFakeMaintenanceManifestDocumentErroringRawIterator(err error) MaintenanceManifestDocumentRawIterator {
	return &fakeMaintenanceManifestDocumentErroringRawIterator{err: err}
}

type fakeMaintenanceManifestDocumentErroringRawIterator struct {
	err error
}

func (i *fakeMaintenanceManifestDocumentErroringRawIterator) Next(ctx context.Context, maxItemCount int) (*pkg.MaintenanceManifestDocuments, error) {
	return nil, i.err
}

func (i *fakeMaintenanceManifestDocumentErroringRawIterator) NextRaw(context.Context, int, interface{}) error {
	return i.err
}

func (i *fakeMaintenanceManifestDocumentErroringRawIterator) Continuation() string {
	return ""
}
