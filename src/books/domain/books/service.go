package books

type Service struct {
	bookRepo Repository
}

func NewService(bookRepo *Repository) *Service {
	return &Service{bookRepo: *bookRepo}
}

func (s *Service) GetBookById(id int) (*Book, error) {
	book, err := s.bookRepo.GetById(id)
	if err != nil {
		return &Book{}, err
	}
	return book, nil
}

func (s *Service) GetBooks() ([]*Book, error) {
	books, err := s.bookRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *Service) AddBook(book *Book) (*Book, error) {
	// TODO: Check if books already exists
	_, err := s.bookRepo.Create(book)
	if err != nil {
		return &Book{}, err
	}
	return book, nil
}
