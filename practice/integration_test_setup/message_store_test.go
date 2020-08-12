package some

import "testing"

type Buildable interface {
	Build(t *testing.T)
}

func (message Message) Build(t *testing.T) {
	message, err := NewMessage()
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { ds.Delete(message.MessageID) })
}

func TestMessageStore(t *testing.T) {
	obj, cleanup := factory.Build(t, factory.DataStore{
		db: NewS3DataStore(),
		Messages: []Message{
			fixutes.BasicMessage(),
		},
	})
	defer cleanup()

}
