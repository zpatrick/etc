package main

import "time"

type Mailbox interface {
	Sync() error
}

type GMAILMailbox struct {
	Syncer Syncer
	api    gmail.API
	email  string
}

func NewGMailMailbox(api gmail.API, email string) *GMAILMailbox {
	// one way to fold multiple syncers into a single syncer object
	// so we can account for mailboxes with or without folders
	syncers := make([]Syncer, len(folders))
	for i, folder := range folders {
		syncers[i] = NewSyncer(NewGMAILMessageFetcher(api, folder))
	}

	return &GMAILMailbox{
		Syncer: Fold(syncers...),
		api:    api,
		email:  email,
	}
}

func NewGMAILMessageFetcher(api gmail.API, folder string) MessageFetcher {
	return func(since time.Time) ([]Message, error) {
		mimes, err := api.getSince(since)
		if err != nil {
			return nil, err
		}

		messages := make([]Message, len(mimes))
		for i, mime := range mimes {
			messages[i] = BuildMessage(mime)
			messages[i].Metadata = MessageMetadata{
				"source": "gmail",
				"foo":    "bar",
			}
		}

		return messages
	}
}

func NewEWSMailbox(redis redis.API, api ews.API, email string) *EWSMailbox {
	return &GMAILMailbox{
		// one way to wrap syncers w/ redis bounds
		Syncer: NewRedisSyncerWrapper(redis, NewSyncer(api)),
		api:    api,
		email:  email,
	}
}

func NewRedisSyncerWrapper(redis redis.API, syncer Syncer) Syncer {
	return func(since time.Time) error {
		bounds, err := redis.Bounds()
		if err != nil {
			return err
		}

		if bounds.before < since {
			return nil
		}

		return syncer.Sync(since)
	}
}
