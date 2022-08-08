// Code generated by ent, DO NOT EDIT.

package duplicatenumbermessage

import (
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Hello applies equality check predicate on the "hello" field. It's identical to HelloEQ.
func Hello(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHello), v))
	})
}

// World applies equality check predicate on the "world" field. It's identical to WorldEQ.
func World(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWorld), v))
	})
}

// HelloEQ applies the EQ predicate on the "hello" field.
func HelloEQ(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHello), v))
	})
}

// HelloNEQ applies the NEQ predicate on the "hello" field.
func HelloNEQ(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHello), v))
	})
}

// HelloIn applies the In predicate on the "hello" field.
func HelloIn(vs ...string) predicate.DuplicateNumberMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldHello), v...))
	})
}

// HelloNotIn applies the NotIn predicate on the "hello" field.
func HelloNotIn(vs ...string) predicate.DuplicateNumberMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldHello), v...))
	})
}

// HelloGT applies the GT predicate on the "hello" field.
func HelloGT(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHello), v))
	})
}

// HelloGTE applies the GTE predicate on the "hello" field.
func HelloGTE(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHello), v))
	})
}

// HelloLT applies the LT predicate on the "hello" field.
func HelloLT(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHello), v))
	})
}

// HelloLTE applies the LTE predicate on the "hello" field.
func HelloLTE(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHello), v))
	})
}

// HelloContains applies the Contains predicate on the "hello" field.
func HelloContains(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHello), v))
	})
}

// HelloHasPrefix applies the HasPrefix predicate on the "hello" field.
func HelloHasPrefix(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHello), v))
	})
}

// HelloHasSuffix applies the HasSuffix predicate on the "hello" field.
func HelloHasSuffix(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHello), v))
	})
}

// HelloEqualFold applies the EqualFold predicate on the "hello" field.
func HelloEqualFold(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHello), v))
	})
}

// HelloContainsFold applies the ContainsFold predicate on the "hello" field.
func HelloContainsFold(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHello), v))
	})
}

// WorldEQ applies the EQ predicate on the "world" field.
func WorldEQ(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWorld), v))
	})
}

// WorldNEQ applies the NEQ predicate on the "world" field.
func WorldNEQ(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWorld), v))
	})
}

// WorldIn applies the In predicate on the "world" field.
func WorldIn(vs ...string) predicate.DuplicateNumberMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldWorld), v...))
	})
}

// WorldNotIn applies the NotIn predicate on the "world" field.
func WorldNotIn(vs ...string) predicate.DuplicateNumberMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldWorld), v...))
	})
}

// WorldGT applies the GT predicate on the "world" field.
func WorldGT(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWorld), v))
	})
}

// WorldGTE applies the GTE predicate on the "world" field.
func WorldGTE(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWorld), v))
	})
}

// WorldLT applies the LT predicate on the "world" field.
func WorldLT(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWorld), v))
	})
}

// WorldLTE applies the LTE predicate on the "world" field.
func WorldLTE(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWorld), v))
	})
}

// WorldContains applies the Contains predicate on the "world" field.
func WorldContains(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldWorld), v))
	})
}

// WorldHasPrefix applies the HasPrefix predicate on the "world" field.
func WorldHasPrefix(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldWorld), v))
	})
}

// WorldHasSuffix applies the HasSuffix predicate on the "world" field.
func WorldHasSuffix(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldWorld), v))
	})
}

// WorldEqualFold applies the EqualFold predicate on the "world" field.
func WorldEqualFold(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldWorld), v))
	})
}

// WorldContainsFold applies the ContainsFold predicate on the "world" field.
func WorldContainsFold(v string) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldWorld), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DuplicateNumberMessage) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DuplicateNumberMessage) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DuplicateNumberMessage) predicate.DuplicateNumberMessage {
	return predicate.DuplicateNumberMessage(func(s *sql.Selector) {
		p(s.Not())
	})
}
