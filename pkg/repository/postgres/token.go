package psql

type TokenRepository struct {
	// db *bolt.Db
	db string
}

func NewTokenRepository(db string) *TokenRepository {
	return &TokenRepository{db: db}
}

// func (r *TokenRepository) Save(chatID int64, token string, bucket repository.Bucket) error {

// }

// func (r *TokenRepository) Get(chatID int64, bucket repository.Bucket) (string, error) {

// }
