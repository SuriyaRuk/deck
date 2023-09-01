package state

// LimitKeyQuotasCollection stores and indexes limit-key-quota credentials.
type LimitKeyQuotasCollection struct {
	credentialsCollection
}

func newLimitKeyQuotasCollection(common collection) *LimitKeyQuotasCollection {
	return &LimitKeyQuotasCollection{
		credentialsCollection: credentialsCollection{
			collection: common,
			CredType:   "limit-key-quota",
		},
	}
}

// Add adds a limit-key-quota credential to LimitKeyQuotasCollection
func (k *LimitKeyQuotasCollection) Add(limitKeyQuota LimitKeyQuota) error {
	cred := (entity)(&limitKeyQuota)
	return k.credentialsCollection.Add(cred)
}

// Get gets a limit-key-quota credential by key or ID.
func (k *LimitKeyQuotasCollection) Get(keyOrID string) (*LimitKeyQuota, error) {
	cred, err := k.credentialsCollection.Get(keyOrID)
	if err != nil {
		return nil, err
	}

	limitKeyQuota, ok := cred.(*LimitKeyQuota)
	if !ok {
		panic(unexpectedType)
	}
	return &LimitKeyQuota{LimitKeyQuota: *limitKeyQuota.DeepCopy()}, nil
}

// GetAllByConsumerID returns all limit-key-quota credentials
// belong to a Consumer with id.
func (k *LimitKeyQuotasCollection) GetAllByConsumerID(id string) ([]*LimitKeyQuota,
	error,
) {
	creds, err := k.credentialsCollection.GetAllByConsumerID(id)
	if err != nil {
		return nil, err
	}

	var res []*LimitKeyQuota
	for _, cred := range creds {
		r, ok := cred.(*LimitKeyQuota)
		if !ok {
			panic(unexpectedType)
		}
		res = append(res, &LimitKeyQuota{LimitKeyQuota: *r.DeepCopy()})
	}
	return res, nil
}

// Update updates an existing limit-key-quota credential.
func (k *LimitKeyQuotasCollection) Update(limitKeyQuota LimitKeyQuota) error {
	cred := (entity)(&limitKeyQuota)
	return k.credentialsCollection.Update(cred)
}

// Delete deletes a limit-key-quota credential by key or ID.
func (k *LimitKeyQuotasCollection) Delete(keyOrID string) error {
	return k.credentialsCollection.Delete(keyOrID)
}

// GetAll gets all limit-key-quota credentials.
func (k *LimitKeyQuotasCollection) GetAll() ([]*LimitKeyQuota, error) {
	creds, err := k.credentialsCollection.GetAll()
	if err != nil {
		return nil, err
	}

	var res []*LimitKeyQuota
	for _, cred := range creds {
		r, ok := cred.(*LimitKeyQuota)
		if !ok {
			panic(unexpectedType)
		}
		res = append(res, &LimitKeyQuota{LimitKeyQuota: *r.DeepCopy()})
	}
	return res, nil
}
