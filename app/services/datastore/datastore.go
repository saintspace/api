package datastore

// Datastore is responsible for storing, retrieving, and modifying data
type Datastore struct {
	kvStore iKeyValueStore
}

func New(kvStore iKeyValueStore) *Datastore {
	return &Datastore{
		kvStore: kvStore,
	}
}

type iKeyValueStore interface {
	EmailSubscriptionItemExists(email string) (bool, error)
	CreateEmailSubscriptionItem(email string, subscriptionToken string, isVerified bool) error
	VerifyEmailSubscription(email string) error
}

func (s *Datastore) CreateEmailSubscription(
	email string,
	subscriptionToken string,
	isVerified bool,
) error {
	return s.kvStore.CreateEmailSubscriptionItem(email, subscriptionToken, isVerified)
}

func (s *Datastore) CheckEmailSubscriptionExists(email string) (bool, error) {
	return s.kvStore.EmailSubscriptionItemExists(email)
}

func (s *Datastore) VerifyEmailSubscription(email string) error {
	return s.kvStore.VerifyEmailSubscription(email)
}
