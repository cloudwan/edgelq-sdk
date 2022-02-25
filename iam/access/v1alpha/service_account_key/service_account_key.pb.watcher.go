// Code generated by protoc-gen-goten-access
// Resource: ServiceAccountKey
// DO NOT EDIT!!!

package service_account_key_access

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/view"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	service_account_key_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha/service_account_key"
	service_account_key "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/service_account_key"
)

// TODO: This watcher is for stateless watch type only de facto as of now.
// Ordering and multiple filters do not cooperate at all.
// We could:
// * Add Cursor as parameter to NewWatcher + page size to WatcherConfig
// * Add ResetCursor method to Watcher - only one page available at the time byt still dynamic
// * Merge many pages with sorting/pagination within watcher and forward aggregated changes.

// Watcher is higher level stateful watcher with dynamic + multi filter support.
type Watcher struct {
	client             service_account_key_client.ServiceAccountKeyServiceClient
	config             *WatcherConfig
	outputEvtChan      chan WatcherEvent
	iOutputEvtChan     chan gotenaccess.WatcherEvent
	queryWatcherStates map[int]*queryWatcherState
	watcherEvtsChan    chan *QueryWatcherEvent
	filters            []*WatcherFilterParams
	filtersResetChan   chan []*WatcherFilterParams
	inSyncFlag         int32
	nextIdentifier     int
	watcherCtx         context.Context
	watcherCtxCancel   context.CancelFunc
	useIChan           int32
}

type WatcherConfig struct {
	*gotenaccess.WatcherConfigBase

	// common params that must be shared across queries
	WatchType watch_type.WatchType
	View      view.View
	FieldMask *service_account_key.ServiceAccountKey_FieldMask
	OrderBy   *service_account_key.OrderBy
	ChunkSize int
}

type WatcherFilterParams struct {
	Parent *service_account_key.ParentReference
	Filter *service_account_key.Filter
}

func (p *WatcherFilterParams) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%s.%s", p.Parent, p.Filter)
}

func (p *WatcherFilterParams) GetIFilter() gotenresource.Filter {
	return p.Filter
}

func (p *WatcherFilterParams) GetIParentRef() gotenresource.Reference {
	return p.Parent
}

type queryWatcherState struct {
	cancel          context.CancelFunc
	watcher         *QueryWatcher
	inSync          bool
	cache           map[service_account_key.Name]*service_account_key.ServiceAccountKey
	filter          *WatcherFilterParams
	inSnapshot      bool
	pendingSnapshot []*service_account_key.ServiceAccountKeyChange
}

func NewWatcher(client service_account_key_client.ServiceAccountKeyServiceClient, config *WatcherConfig, filters ...*WatcherFilterParams) *Watcher {
	ctx, cancel := context.WithCancel(context.Background())
	return &Watcher{
		client:             client,
		config:             config,
		filters:            filters,
		outputEvtChan:      make(chan WatcherEvent, config.WatcherEventBufferSize),
		queryWatcherStates: map[int]*queryWatcherState{},
		watcherEvtsChan:    make(chan *QueryWatcherEvent, config.WatcherEventBufferSize),
		filtersResetChan:   make(chan []*WatcherFilterParams, 1),
		watcherCtx:         ctx,
		watcherCtxCancel:   cancel,
	}
}

func (pw *Watcher) Events() <-chan WatcherEvent {
	return pw.outputEvtChan
}

func (pw *Watcher) IEvents() <-chan gotenaccess.WatcherEvent {
	if atomic.CompareAndSwapInt32(&pw.useIChan, 0, 1) {
		pw.iOutputEvtChan = make(chan gotenaccess.WatcherEvent, pw.config.WatcherEventBufferSize)
		go func() {
			for {
				select {
				case <-pw.watcherCtx.Done():
				case evt := <-pw.outputEvtChan:
					select {
					case <-pw.watcherCtx.Done():
					case pw.iOutputEvtChan <- &evt:
					}
				}
			}
		}()
	}
	return pw.iOutputEvtChan
}

func (pw *Watcher) InSync() bool {
	return atomic.LoadInt32(&pw.inSyncFlag) > 0
}

func (pw *Watcher) GetFilters() []*WatcherFilterParams {
	copied := make([]*WatcherFilterParams, 0, len(pw.filters))
	for _, query := range pw.filters {
		cQuery := &WatcherFilterParams{}
		if query.Parent != nil {
			cQuery.Parent = &service_account_key.ParentReference{}
			*cQuery.Parent = *query.Parent
		}
		if query.Filter != nil {
			strValue, _ := query.Filter.ProtoString()
			cQuery.Filter = &service_account_key.Filter{}
			_ = cQuery.Filter.ParseProtoString(strValue)
		}
		copied = append(copied, cQuery)
	}
	return copied
}

func (pw *Watcher) GetIFilters() []gotenaccess.WatcherFilterParams {
	filters := pw.GetFilters()
	result := make([]gotenaccess.WatcherFilterParams, 0, len(filters))
	for _, filter := range filters {
		result = append(result, filter)
	}
	return result
}

func (pw *Watcher) ResetFilters(ctx context.Context, filters ...*WatcherFilterParams) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case pw.filtersResetChan <- filters:
	}
	return nil
}

func (pw *Watcher) ResetIFilters(ctx context.Context, filters ...gotenaccess.WatcherFilterParams) error {
	typedFilters := make([]*WatcherFilterParams, 0, len(filters))
	for _, filter := range filters {
		typedFilters = append(typedFilters, filter.(*WatcherFilterParams))
	}
	return pw.ResetFilters(ctx, typedFilters...)
}

func (pw *Watcher) Run(ctx context.Context) error {
	log := ctxlogrus.Extract(ctx).
		WithField("watcher", "serviceAccountKey-watcher")
	ctx = ctxlogrus.ToContext(ctx, log)

	log.Debugf("Initializing initial queries")
	pw.resetFilters(ctx, pw.filters)

	log.Debugf("running")
	defer func() {
		close(pw.outputEvtChan)
		pw.watcherCtxCancel()
	}()

	for {
		select {
		case newFilters := <-pw.filtersResetChan:
			pw.resetFilters(ctx, newFilters)
		case evt := <-pw.watcherEvtsChan:
			if evt.LostSync {
				pw.onQueryLostSync(ctx, evt)
			} else if evt.CheckSize {
				pw.onQueryVerifySnapshotSize(ctx, evt)
			} else if evt.Reset {
				pw.onQueryResetStateIndicator(evt)
			} else {
				pw.onQueryChanges(ctx, evt)
			}
		case <-ctx.Done():
			log.Debugf("context done, reason: %s", ctx.Err())
			if errors.Is(ctx.Err(), context.Canceled) {
				return nil
			} else {
				return ctx.Err()
			}
		}
	}
}

func (pw *Watcher) resetFilters(ctx context.Context, filters []*WatcherFilterParams) {
	pw.filters = filters

	filtersMap := map[string]*WatcherFilterParams{}
	for _, filter := range filters {
		filtersMap[filter.String()] = filter
	}

	// Eliminate existing queries from queriesMap, eliminate queries that should be stopped
	for identifier, state := range pw.queryWatcherStates {
		key := state.filter.String()
		if _, exists := filtersMap[key]; !exists {
			state.cancel()
			delete(pw.queryWatcherStates, identifier)
		} else {
			delete(filtersMap, key)
		}
	}

	// start new queries with new filters
	for _, filter := range filtersMap {
		wctx, cancel := context.WithCancel(ctx)
		newWatcher := NewQueryWatcher(
			pw.nextIdentifier,
			pw.client,
			&QueryWatcherParams{
				Parent:           filter.Parent,
				Filter:           filter.Filter,
				View:             pw.config.View,
				FieldMask:        pw.config.FieldMask,
				OrderBy:          pw.config.OrderBy,
				ChunkSize:        pw.config.ChunkSize,
				WatchType:        pw.config.WatchType,
				RecoveryDeadline: pw.config.RecoveryDeadline,
				RetryTimeout:     pw.config.RetryTimeout,
			},
			pw.watcherEvtsChan)
		pw.queryWatcherStates[pw.nextIdentifier] = &queryWatcherState{
			cancel:     cancel,
			watcher:    newWatcher,
			cache:      map[service_account_key.Name]*service_account_key.ServiceAccountKey{},
			filter:     filter,
			inSnapshot: true,
		}
		pw.nextIdentifier++

		go func() {
			_ = newWatcher.Run(wctx)
		}()
	}

	// if all in sync (possible if queries were eliminated with resetQueries call), send snapshot, otherwise
	// communicate "lost sync"
	if pw.allInSync() {
		pw.processSnapshot(ctx)
	} else if pw.inSyncFlag > 0 {
		pw.processSyncLost(ctx)
	}
}

func (pw *Watcher) onQueryLostSync(ctx context.Context, evt *QueryWatcherEvent) {
	state := pw.queryWatcherStates[evt.Identifier]
	if state == nil {
		return
	}
	state.inSync = false
	state.cache = map[service_account_key.Name]*service_account_key.ServiceAccountKey{}
	state.inSnapshot = true
	if pw.inSyncFlag > 0 {
		pw.processSyncLost(ctx)
	}
}

func (pw *Watcher) onQueryVerifySnapshotSize(ctx context.Context, evt *QueryWatcherEvent) {
	state := pw.queryWatcherStates[evt.Identifier]
	if state == nil {
		return
	}
	currentSize := state.computeSize(evt.Changes)
	if currentSize != evt.SnapshotSize {
		ctxlogrus.Extract(ctx).
			Warnf("Detected mismatch in snapshot size for filter %s: expected %d, has %d",
				state.filter, evt.SnapshotSize, currentSize)
		state.cancel()
		delete(pw.queryWatcherStates, evt.Identifier)

		// refresh this particular watcher
		wctx, cancel := context.WithCancel(ctx)
		newWatcher := NewQueryWatcher(
			pw.nextIdentifier,
			pw.client,
			&QueryWatcherParams{
				Parent:           state.filter.Parent,
				Filter:           state.filter.Filter,
				View:             pw.config.View,
				FieldMask:        pw.config.FieldMask,
				OrderBy:          pw.config.OrderBy,
				ChunkSize:        pw.config.ChunkSize,
				WatchType:        pw.config.WatchType,
				RecoveryDeadline: pw.config.RecoveryDeadline,
				RetryTimeout:     pw.config.RetryTimeout,
			},
			pw.watcherEvtsChan)
		pw.queryWatcherStates[pw.nextIdentifier] = &queryWatcherState{
			cancel:     cancel,
			watcher:    newWatcher,
			cache:      state.cache,
			filter:     state.filter,
			inSnapshot: true,
		}
		pw.nextIdentifier++

		go func() {
			_ = newWatcher.Run(wctx)
		}()
	}
}

func (pw *Watcher) onQueryResetStateIndicator(evt *QueryWatcherEvent) {
	state := pw.queryWatcherStates[evt.Identifier]
	if state == nil {
		return
	}
	state.inSnapshot = true
	state.pendingSnapshot = nil
}

func (pw *Watcher) onQueryChanges(ctx context.Context, evt *QueryWatcherEvent) {
	var changes []*WatcherEventChange
	state := pw.queryWatcherStates[evt.Identifier]
	if state == nil {
		return
	}

	// evt.InSync is set when snapshot is complete or for every incremental update
	// Since Watcher is stateful, we will collect query snapshot data without processing
	// till its complete.
	if state.inSnapshot {
		state.pendingSnapshot = append(state.pendingSnapshot, evt.Changes...)
		if evt.InSync {
			// processes snapshot for this query only
			// changes however may be only a diff from previous known state
			changes = state.processSnapshot()
		}
	} else {
		changes = state.transform(evt.Changes)
	}
	if evt.InSync {
		if !state.inSync {
			state.inSync = true
			if pw.allInSync() {
				pw.processSnapshot(ctx)
				return
			}
		} else if pw.inSyncFlag > 0 {
			evt := &WatcherEvent{WatcherEventBase: gotenaccess.NewWatcherEventBase(false), Changes: changes}
			pw.sendEvent(ctx, evt)
		}
	}
}

func (pw *Watcher) processSyncLost(ctx context.Context) {
	ctxlogrus.Extract(ctx).Infof("Notifiying watcher-wide sync lost")
	atomic.StoreInt32(&pw.inSyncFlag, 0)
	evt := &WatcherEvent{WatcherEventBase: gotenaccess.NewWatcherEventBaseLostSync()}
	pw.sendEvent(ctx, evt)
}

func (pw *Watcher) processSnapshot(ctx context.Context) {
	ctxlogrus.Extract(ctx).Infof("Watcher in sync")
	atomic.StoreInt32(&pw.inSyncFlag, 1)
	snapshot := make([]*WatcherEventChange, 0)
	for _, state := range pw.queryWatcherStates {
		for _, item := range state.cache {
			snapshot = append(snapshot, NewAddWatcherEventChange(item))
		}
	}
	evt := &WatcherEvent{WatcherEventBase: gotenaccess.NewWatcherEventBase(true), Changes: snapshot}
	pw.sendEvent(ctx, evt)
}

func (pw *Watcher) allInSync() bool {
	for _, state := range pw.queryWatcherStates {
		if !state.inSync {
			return false
		}
	}
	return true
}

func (pw *Watcher) sendEvent(ctx context.Context, evt *WatcherEvent) {
	select {
	case <-ctx.Done():
	case pw.outputEvtChan <- *evt:
	}
}

type WatcherEventChange struct {
	pre, post *service_account_key.ServiceAccountKey
	name      *service_account_key.Name
}

func NewAddWatcherEventChange(resource *service_account_key.ServiceAccountKey) *WatcherEventChange {
	return &WatcherEventChange{
		post: resource,
		name: resource.GetName(),
	}
}

func NewModifyWatcherEventChange(current, previous *service_account_key.ServiceAccountKey) *WatcherEventChange {
	return &WatcherEventChange{
		pre:  previous,
		post: current,
		name: current.GetName(),
	}
}

func NewDeleteWatcherEventChange(deleted *service_account_key.ServiceAccountKey) *WatcherEventChange {
	return &WatcherEventChange{
		pre:  deleted,
		name: deleted.GetName(),
	}
}

func (c *WatcherEventChange) GetName() *service_account_key.Name {
	return c.name
}

func (c *WatcherEventChange) GetRawName() gotenresource.Name {
	return c.name
}

func (c *WatcherEventChange) IsAdd() bool {
	return c.pre == nil && c.post != nil
}

func (c *WatcherEventChange) IsModify() bool {
	return c.pre != nil && c.post != nil
}

func (c *WatcherEventChange) IsDelete() bool {
	return c.pre != nil && c.post == nil
}

func (c *WatcherEventChange) GetAdded() *service_account_key.ServiceAccountKey {
	if c.IsAdd() {
		return c.post
	}
	return nil
}

func (c *WatcherEventChange) GetDeleted() *service_account_key.ServiceAccountKey {
	if c.IsDelete() {
		return c.pre
	}
	return nil
}

func (c *WatcherEventChange) GetPrevious() *service_account_key.ServiceAccountKey {
	if c.IsModify() {
		return c.pre
	}
	return nil
}

func (c *WatcherEventChange) GetCurrent() *service_account_key.ServiceAccountKey {
	if c.IsAdd() || c.IsModify() {
		return c.post
	}
	return nil
}

func (c *WatcherEventChange) GetRawAdded() gotenresource.Resource {
	return c.GetAdded()
}

func (c *WatcherEventChange) GetRawDeleted() gotenresource.Resource {
	return c.GetDeleted()
}

func (c *WatcherEventChange) GetRawPrevious() gotenresource.Resource {
	return c.GetPrevious()
}

func (c *WatcherEventChange) GetRawCurrent() gotenresource.Resource {
	return c.GetCurrent()
}

type WatcherEvent struct {
	gotenaccess.WatcherEventBase
	Changes []*WatcherEventChange
}

func (e *WatcherEvent) GetAt(index int) *WatcherEventChange {
	if e == nil || len(e.Changes) <= index {
		return nil
	}
	return e.Changes[index]
}

func (e *WatcherEvent) GetRawAt(index int) gotenaccess.WatcherEventChange {
	return e.GetAt(index)
}

func (e *WatcherEvent) Length() int {
	if e == nil {
		return 0
	}
	return len(e.Changes)
}

func (e *WatcherEvent) AppendChange(change *WatcherEventChange) {
	e.Changes = append(e.Changes, change)
}

func (e *WatcherEvent) AppendRawChange(change gotenaccess.WatcherEventChange) {
	e.Changes = append(e.Changes, change.(*WatcherEventChange))
}

// Merge makes a shallow merge of two events
func (e *WatcherEvent) Merge(src *WatcherEvent) {
	if src == nil {
		return
	}
	if src.LostSync() || src.Resync() {
		e.WatcherEventBase = src.WatcherEventBase
		e.Changes = src.Changes
		return
	}
	if !e.WatcherEventBase.Resync() {
		e.WatcherEventBase = src.WatcherEventBase
		e.Changes = append(e.Changes, src.Changes...)
		return
	}

	// merge non-resync into resync
	byName := make(map[service_account_key.Name]*service_account_key.ServiceAccountKey, len(src.Changes)+len(e.Changes))
	for _, currentChange := range e.Changes {
		serviceAccountKey := currentChange.GetCurrent()
		byName[*serviceAccountKey.GetName()] = serviceAccountKey
	}
	for _, change := range src.Changes {
		if change.IsDelete() {
			delete(byName, *change.GetName())
		} else {
			serviceAccountKey := change.GetCurrent()
			byName[*serviceAccountKey.GetName()] = serviceAccountKey
		}
	}
	e.Changes = make([]*WatcherEventChange, 0, len(byName))
	for _, serviceAccountKey := range byName {
		e.Changes = append(e.Changes, NewAddWatcherEventChange(serviceAccountKey))
	}
}

func (qws *queryWatcherState) processSnapshot() []*WatcherEventChange {
	watcherEventChanges := make([]*WatcherEventChange, 0, len(qws.pendingSnapshot))
	prevCache := qws.cache
	qws.cache = map[service_account_key.Name]*service_account_key.ServiceAccountKey{}

	for _, change := range qws.pendingSnapshot {
		var currentItem *service_account_key.ServiceAccountKey
		switch item := change.ChangeType.(type) {
		case *service_account_key.ServiceAccountKeyChange_Added_:
			currentItem = item.Added.GetServiceAccountKey()
		case *service_account_key.ServiceAccountKeyChange_Current_:
			currentItem = item.Current.GetServiceAccountKey()
		}
		name := *currentItem.GetName()
		qws.cache[name] = currentItem
		previous, exists := prevCache[name]
		if !exists {
			watcherEventChanges = append(watcherEventChanges, NewAddWatcherEventChange(currentItem))
		} else {
			if !proto.Equal(currentItem, previous) {
				watcherEventChanges = append(watcherEventChanges, NewModifyWatcherEventChange(currentItem, previous))
			}
			delete(prevCache, name)
		}
	}
	// remaining items
	for _, resource := range prevCache {
		watcherEventChanges = append(watcherEventChanges, NewDeleteWatcherEventChange(resource))
	}
	qws.pendingSnapshot = nil
	qws.inSnapshot = false
	return watcherEventChanges
}

func (qws *queryWatcherState) transform(rawChanges []*service_account_key.ServiceAccountKeyChange) []*WatcherEventChange {
	watcherEventChanges := make([]*WatcherEventChange, 0, len(rawChanges))

	for _, change := range rawChanges {
		switch item := change.ChangeType.(type) {
		case *service_account_key.ServiceAccountKeyChange_Added_:
			newItem := item.Added.GetServiceAccountKey()
			qws.cache[*newItem.GetName()] = newItem
			watcherEventChanges = append(watcherEventChanges, NewAddWatcherEventChange(newItem))
		case *service_account_key.ServiceAccountKeyChange_Modified_:
			updatedItem := item.Modified.GetServiceAccountKey()
			previousItem := qws.cache[*updatedItem.GetName()]
			qws.cache[*updatedItem.GetName()] = updatedItem
			watcherEventChanges = append(watcherEventChanges, NewModifyWatcherEventChange(updatedItem, previousItem))
		case *service_account_key.ServiceAccountKeyChange_Current_:
			currentItem := item.Current.GetServiceAccountKey()
			previousItem := qws.cache[*currentItem.GetName()]
			qws.cache[*currentItem.GetName()] = currentItem
			if previousItem != nil {
				watcherEventChanges = append(watcherEventChanges, NewModifyWatcherEventChange(currentItem, previousItem))
			} else {
				watcherEventChanges = append(watcherEventChanges, NewAddWatcherEventChange(currentItem))
			}
		case *service_account_key.ServiceAccountKeyChange_Removed_:
			deletedItem := qws.cache[*item.Removed.Name]
			delete(qws.cache, *item.Removed.Name)
			watcherEventChanges = append(watcherEventChanges, NewDeleteWatcherEventChange(deletedItem))
		}
	}
	return watcherEventChanges
}

func (qws *queryWatcherState) computeSize(pendingChanges []*service_account_key.ServiceAccountKeyChange) int64 {
	size := int64(len(qws.cache))
	for _, pendingChange := range pendingChanges {
		name := *pendingChange.GetServiceAccountKeyName()
		if pendingChange.IsDelete() {
			if _, exists := qws.cache[name]; exists {
				size--
			}
		} else {
			if _, exists := qws.cache[name]; !exists {
				size++
			}
		}
	}
	return size
}

func init() {
	gotenaccess.GetRegistry().RegisterWatcherConstructor(service_account_key.GetDescriptor(), func(cc grpc.ClientConnInterface, params *gotenaccess.WatcherConfigParams, filters ...gotenaccess.WatcherFilterParams) gotenaccess.Watcher {
		cfg := &WatcherConfig{
			WatcherConfigBase: params.CfgBase,
			WatchType:         params.WatchType,
			View:              params.View,
			ChunkSize:         params.ChunkSize,
		}
		if params.FieldMask != nil {
			cfg.FieldMask = params.FieldMask.(*service_account_key.ServiceAccountKey_FieldMask)
		}
		if params.OrderBy != nil {
			cfg.OrderBy = params.OrderBy.(*service_account_key.OrderBy)
		}
		typedFilters := make([]*WatcherFilterParams, 0, len(filters))
		for _, filter := range filters {
			typedFilters = append(typedFilters, filter.(*WatcherFilterParams))
		}
		return NewWatcher(service_account_key_client.NewServiceAccountKeyServiceClient(cc), cfg, typedFilters...)
	})
	gotenaccess.GetRegistry().RegisterWatcherFilterConstructor(service_account_key.GetDescriptor(), func(filter gotenresource.Filter, parent gotenresource.Reference) gotenaccess.WatcherFilterParams {
		params := &WatcherFilterParams{}
		if filter != nil {
			params.Filter = filter.(*service_account_key.Filter)
		}
		if parent != nil {
			params.Parent = parent.(*service_account_key.ParentReference)
		}
		return params
	})
}
