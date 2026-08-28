package storage

import "context"

func (s *Store) Healthy() bool { return s.Ping(context.Background()) == nil }
