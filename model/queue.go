package pq

type PQueue interface {
	Push(element QueueElement) error
	Pop(element QueueElement) error
	Ack(session string) error
}
