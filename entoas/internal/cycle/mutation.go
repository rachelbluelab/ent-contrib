// Code generated by ent, DO NOT EDIT.

package cycle

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/contrib/entoas/internal/cycle/predicate"
	"entgo.io/contrib/entoas/internal/cycle/user"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUser = "User"
)

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op               Op
	typ              string
	id               *int
	clearedFields    map[string]struct{}
	friends          map[int]struct{}
	removedfriends   map[int]struct{}
	clearedfriends   bool
	followers        map[int]struct{}
	removedfollowers map[int]struct{}
	clearedfollowers bool
	following        map[int]struct{}
	removedfollowing map[int]struct{}
	clearedfollowing bool
	parent           *int
	clearedparent    bool
	children         map[int]struct{}
	removedchildren  map[int]struct{}
	clearedchildren  bool
	done             bool
	oldValue         func(context.Context) (*User, error)
	predicates       []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("cycle: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// AddFriendIDs adds the "friends" edge to the User entity by ids.
func (m *UserMutation) AddFriendIDs(ids ...int) {
	if m.friends == nil {
		m.friends = make(map[int]struct{})
	}
	for i := range ids {
		m.friends[ids[i]] = struct{}{}
	}
}

// ClearFriends clears the "friends" edge to the User entity.
func (m *UserMutation) ClearFriends() {
	m.clearedfriends = true
}

// FriendsCleared reports if the "friends" edge to the User entity was cleared.
func (m *UserMutation) FriendsCleared() bool {
	return m.clearedfriends
}

// RemoveFriendIDs removes the "friends" edge to the User entity by IDs.
func (m *UserMutation) RemoveFriendIDs(ids ...int) {
	if m.removedfriends == nil {
		m.removedfriends = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.friends, ids[i])
		m.removedfriends[ids[i]] = struct{}{}
	}
}

// RemovedFriends returns the removed IDs of the "friends" edge to the User entity.
func (m *UserMutation) RemovedFriendsIDs() (ids []int) {
	for id := range m.removedfriends {
		ids = append(ids, id)
	}
	return
}

// FriendsIDs returns the "friends" edge IDs in the mutation.
func (m *UserMutation) FriendsIDs() (ids []int) {
	for id := range m.friends {
		ids = append(ids, id)
	}
	return
}

// ResetFriends resets all changes to the "friends" edge.
func (m *UserMutation) ResetFriends() {
	m.friends = nil
	m.clearedfriends = false
	m.removedfriends = nil
}

// AddFollowerIDs adds the "followers" edge to the User entity by ids.
func (m *UserMutation) AddFollowerIDs(ids ...int) {
	if m.followers == nil {
		m.followers = make(map[int]struct{})
	}
	for i := range ids {
		m.followers[ids[i]] = struct{}{}
	}
}

// ClearFollowers clears the "followers" edge to the User entity.
func (m *UserMutation) ClearFollowers() {
	m.clearedfollowers = true
}

// FollowersCleared reports if the "followers" edge to the User entity was cleared.
func (m *UserMutation) FollowersCleared() bool {
	return m.clearedfollowers
}

// RemoveFollowerIDs removes the "followers" edge to the User entity by IDs.
func (m *UserMutation) RemoveFollowerIDs(ids ...int) {
	if m.removedfollowers == nil {
		m.removedfollowers = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.followers, ids[i])
		m.removedfollowers[ids[i]] = struct{}{}
	}
}

// RemovedFollowers returns the removed IDs of the "followers" edge to the User entity.
func (m *UserMutation) RemovedFollowersIDs() (ids []int) {
	for id := range m.removedfollowers {
		ids = append(ids, id)
	}
	return
}

// FollowersIDs returns the "followers" edge IDs in the mutation.
func (m *UserMutation) FollowersIDs() (ids []int) {
	for id := range m.followers {
		ids = append(ids, id)
	}
	return
}

// ResetFollowers resets all changes to the "followers" edge.
func (m *UserMutation) ResetFollowers() {
	m.followers = nil
	m.clearedfollowers = false
	m.removedfollowers = nil
}

// AddFollowingIDs adds the "following" edge to the User entity by ids.
func (m *UserMutation) AddFollowingIDs(ids ...int) {
	if m.following == nil {
		m.following = make(map[int]struct{})
	}
	for i := range ids {
		m.following[ids[i]] = struct{}{}
	}
}

// ClearFollowing clears the "following" edge to the User entity.
func (m *UserMutation) ClearFollowing() {
	m.clearedfollowing = true
}

// FollowingCleared reports if the "following" edge to the User entity was cleared.
func (m *UserMutation) FollowingCleared() bool {
	return m.clearedfollowing
}

// RemoveFollowingIDs removes the "following" edge to the User entity by IDs.
func (m *UserMutation) RemoveFollowingIDs(ids ...int) {
	if m.removedfollowing == nil {
		m.removedfollowing = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.following, ids[i])
		m.removedfollowing[ids[i]] = struct{}{}
	}
}

// RemovedFollowing returns the removed IDs of the "following" edge to the User entity.
func (m *UserMutation) RemovedFollowingIDs() (ids []int) {
	for id := range m.removedfollowing {
		ids = append(ids, id)
	}
	return
}

// FollowingIDs returns the "following" edge IDs in the mutation.
func (m *UserMutation) FollowingIDs() (ids []int) {
	for id := range m.following {
		ids = append(ids, id)
	}
	return
}

// ResetFollowing resets all changes to the "following" edge.
func (m *UserMutation) ResetFollowing() {
	m.following = nil
	m.clearedfollowing = false
	m.removedfollowing = nil
}

// SetParentID sets the "parent" edge to the User entity by id.
func (m *UserMutation) SetParentID(id int) {
	m.parent = &id
}

// ClearParent clears the "parent" edge to the User entity.
func (m *UserMutation) ClearParent() {
	m.clearedparent = true
}

// ParentCleared reports if the "parent" edge to the User entity was cleared.
func (m *UserMutation) ParentCleared() bool {
	return m.clearedparent
}

// ParentID returns the "parent" edge ID in the mutation.
func (m *UserMutation) ParentID() (id int, exists bool) {
	if m.parent != nil {
		return *m.parent, true
	}
	return
}

// ParentIDs returns the "parent" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// ParentID instead. It exists only for internal usage by the builders.
func (m *UserMutation) ParentIDs() (ids []int) {
	if id := m.parent; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetParent resets all changes to the "parent" edge.
func (m *UserMutation) ResetParent() {
	m.parent = nil
	m.clearedparent = false
}

// AddChildIDs adds the "children" edge to the User entity by ids.
func (m *UserMutation) AddChildIDs(ids ...int) {
	if m.children == nil {
		m.children = make(map[int]struct{})
	}
	for i := range ids {
		m.children[ids[i]] = struct{}{}
	}
}

// ClearChildren clears the "children" edge to the User entity.
func (m *UserMutation) ClearChildren() {
	m.clearedchildren = true
}

// ChildrenCleared reports if the "children" edge to the User entity was cleared.
func (m *UserMutation) ChildrenCleared() bool {
	return m.clearedchildren
}

// RemoveChildIDs removes the "children" edge to the User entity by IDs.
func (m *UserMutation) RemoveChildIDs(ids ...int) {
	if m.removedchildren == nil {
		m.removedchildren = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.children, ids[i])
		m.removedchildren[ids[i]] = struct{}{}
	}
}

// RemovedChildren returns the removed IDs of the "children" edge to the User entity.
func (m *UserMutation) RemovedChildrenIDs() (ids []int) {
	for id := range m.removedchildren {
		ids = append(ids, id)
	}
	return
}

// ChildrenIDs returns the "children" edge IDs in the mutation.
func (m *UserMutation) ChildrenIDs() (ids []int) {
	for id := range m.children {
		ids = append(ids, id)
	}
	return
}

// ResetChildren resets all changes to the "children" edge.
func (m *UserMutation) ResetChildren() {
	m.children = nil
	m.clearedchildren = false
	m.removedchildren = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 0)
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 5)
	if m.friends != nil {
		edges = append(edges, user.EdgeFriends)
	}
	if m.followers != nil {
		edges = append(edges, user.EdgeFollowers)
	}
	if m.following != nil {
		edges = append(edges, user.EdgeFollowing)
	}
	if m.parent != nil {
		edges = append(edges, user.EdgeParent)
	}
	if m.children != nil {
		edges = append(edges, user.EdgeChildren)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeFriends:
		ids := make([]ent.Value, 0, len(m.friends))
		for id := range m.friends {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeFollowers:
		ids := make([]ent.Value, 0, len(m.followers))
		for id := range m.followers {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeFollowing:
		ids := make([]ent.Value, 0, len(m.following))
		for id := range m.following {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeParent:
		if id := m.parent; id != nil {
			return []ent.Value{*id}
		}
	case user.EdgeChildren:
		ids := make([]ent.Value, 0, len(m.children))
		for id := range m.children {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 5)
	if m.removedfriends != nil {
		edges = append(edges, user.EdgeFriends)
	}
	if m.removedfollowers != nil {
		edges = append(edges, user.EdgeFollowers)
	}
	if m.removedfollowing != nil {
		edges = append(edges, user.EdgeFollowing)
	}
	if m.removedchildren != nil {
		edges = append(edges, user.EdgeChildren)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeFriends:
		ids := make([]ent.Value, 0, len(m.removedfriends))
		for id := range m.removedfriends {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeFollowers:
		ids := make([]ent.Value, 0, len(m.removedfollowers))
		for id := range m.removedfollowers {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeFollowing:
		ids := make([]ent.Value, 0, len(m.removedfollowing))
		for id := range m.removedfollowing {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeChildren:
		ids := make([]ent.Value, 0, len(m.removedchildren))
		for id := range m.removedchildren {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 5)
	if m.clearedfriends {
		edges = append(edges, user.EdgeFriends)
	}
	if m.clearedfollowers {
		edges = append(edges, user.EdgeFollowers)
	}
	if m.clearedfollowing {
		edges = append(edges, user.EdgeFollowing)
	}
	if m.clearedparent {
		edges = append(edges, user.EdgeParent)
	}
	if m.clearedchildren {
		edges = append(edges, user.EdgeChildren)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeFriends:
		return m.clearedfriends
	case user.EdgeFollowers:
		return m.clearedfollowers
	case user.EdgeFollowing:
		return m.clearedfollowing
	case user.EdgeParent:
		return m.clearedparent
	case user.EdgeChildren:
		return m.clearedchildren
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	case user.EdgeParent:
		m.ClearParent()
		return nil
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeFriends:
		m.ResetFriends()
		return nil
	case user.EdgeFollowers:
		m.ResetFollowers()
		return nil
	case user.EdgeFollowing:
		m.ResetFollowing()
		return nil
	case user.EdgeParent:
		m.ResetParent()
		return nil
	case user.EdgeChildren:
		m.ResetChildren()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
