# Borg Design Pattern

## Borg
  * An entity with two connections. 
The reason there are two connections is because it is the minimum amount required to create an expanding entity.
  * (Possible) A borg is always a closed system.
It will not function until it has reached a **closed state**.
A borg is in a **closed state** when all nodes have two connections.


A single-node borg will be connected to itself.
A double-node borg will be connected in a circle:
```
|-N-|
-   -
|-N-|
```

## Work
Work in a Borg works like recursive functions.
If you can figure it out with a partner, anything shoudl work.

Example: Something comes in saying we need to do jobs 1-50.
Tell each node in the borg to take one job.  
```
type WorkerNode struct {
	node *borg.Node
}

func (w *WorkerNode) RunJobs(jobs []int) error {
	switch len(jobs) {
		case 0:
			return nil
		case 1:
			log.Println("Got job", jobs[0])
			// push work off to goroutine 
			return nil
		default:
			log.Println("Got job", jobs[0])
			// push work off to goroutine
			jobs = jobs[:len(jobs)-1]
			half := len(jobs) / 2
			return w.borg.Pass(jobs[:half], jobs[half:])
	}
}



```  
