package usecase

type Usecase struct {
	p Provider
}

func NewUsecase(p Provider) *Usecase {
	return &Usecase{
		p: p,
	}
}
