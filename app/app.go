package app

type App interface {
	ListenAndServe() error
}

func New() (App, error) { return nil, nil }
