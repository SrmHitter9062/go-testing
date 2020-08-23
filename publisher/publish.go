package publisher

import "fmt"

type PublishInterface interface {
	Publish(data []byte) error
}

type Publisher struct {}


func (p *Publisher) Publish(data []byte) error {
	fmt.Println("data published")
	return nil
}
