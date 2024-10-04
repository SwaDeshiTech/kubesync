package job

type Job interface {
	Run() error
	Stop() error
}
