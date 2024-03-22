## go-dev-skills
Go version used: `1.22.1`

### Task 1
I have created a problem between cats and more wild felines like lion, tiger, cheetah... Wild felines do a Roar and cats do only Meow. 
The interfaces Cat has `Meow()` and the Feline has `Roar()`, we have a `RoarService` that we want to our cats and felines make a Roar. But our Cat interface doesn't have the `Roar()` method, so we have to create an adapter to solve this. I created the CatFeline adapter that returns a Feline based on a Cat.

#### Run it
`go run main.go -task=1`

### Task 2
I have decided to use the same structure implemented on task 1, but with additional complexity, added channels, goroutines, wait group and mutex.
We add to a slice every roared feline, so we can print the length of it on the end of the method to compare the felines we received. The mutex prevent a race condition problem that ocurred when I was appending the feline names.
The wait group is used to wait for all consumers to end properly.
Goroutines used to consume felines from the felines channel.

#### Run it
`go run main.go -task=2`


### Tests
`go test ./...`