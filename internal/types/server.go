package types

type Server interface {
	Run()
}

type BannerError struct {
	Message string
}
