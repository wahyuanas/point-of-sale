package repository

type accountRepository struct {
	//db *sql.DB
}

func ImplAccountRepository(url string) (AccountRepository, error) {
	// db, err := sql.Open("mysql", url)
	// if err != nil {
	// 	return nil, err
	// }
	// err = db.Ping()
	// if err != nil {
	// 	return nil, err
	// }
	return &accountRepository{}, nil
}

func (acc *accountRepository) Store() error {
	return nil
}
func (acc *accountRepository) Update()        {}
func (acc *accountRepository) Delete()        {}
func (acc *accountRepository) GetById()       {}
func (acc *accountRepository) GetByUserName() {}
