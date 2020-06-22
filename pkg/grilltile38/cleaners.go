package grilltile38

import "bitbucket.org/swigy/grill"

func (gt *Tile38) FlushDB() grill.Cleaner {
	return grill.CleanerFunc(func() error {
		_, err := gt.Client().Do("flushdb")
		return err
	})
}
