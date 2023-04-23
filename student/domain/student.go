package domain

import "time"

type Student struct {
	ID        int // auto increment
	Name      string
	Age       int
	Gender    int
	Hobbies   []int
	CreatedAt time.Time
}
