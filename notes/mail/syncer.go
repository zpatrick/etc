package main

type Mailbox interface {
	Sync() error
}

func (s SyncConsumer) Perform(m Mailbox) error {
	return m.Syncer.Sync()
}
