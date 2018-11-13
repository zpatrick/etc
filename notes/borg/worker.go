package main

import (
	"log"

	"github.com/zpatrick/borg"
)

type Worker struct{}

func NewWorker() *Worker {
	return &Worker{}
}

// this is interesting in that it uses the borg to recusively 
// spread out workers amongst all nodes, and only until that process
// is done do we actually do the work.

/*
would this actually work? 
3-node borg and 5 jobs:

borg structure:
 A---B 
  \ /   
   C 

A gets [1, 2, 3, 4, 5]
  B gets [1, 2]
    A gets [1] and completes job 1. A is blocking until this job is done.  
    C gets [2] and completes job 2. C is blocking until this job is done. 
  C gets [3, 4, 5]
    C tries to send A[3] but A is currently busy on [1]. 
    B[4, 5]
      B tries to send A[4] but A is currently busy on [1].
      B tries to send C[5] but C is currently busy on [2]. 
      This isn't great because now B isnt' doing anything until A or C is done.
      Maybe borgs need to always self-reference?     
   needs to send [3]
 gets [5] and completes job 5. C is blocking until this job is done.
   
If it was a single-node borg, this would cause a huge stack trace I think.
Wait no, because the function stack stops after a successful send.
this also may be a problem with 3 - e.g. maybe there's always 1 stuck node.    
A---B
|   |
C---D

A(1, 2, 3, 4, 5)
  B(1, 2)
    A(1)
    D(2)  
  C(3, 4, 5)
    A(3) -> blocked by A(1)
    D(4, 5) -> blocked by D(2) 

What if non-blocking?
C(3, 4, 5)
  A(3)
  D(4, 5)
    B(4)
    C(5)

A| 1, 3
B| 4
C| 5
D| 2

Wow, that seems to work, lets do 3 again
 A---B
  \ /
   C

A(1, 2, 3, 4, 5)
  B(1, 2)
    A(1)
    B(2)
  C(3, 4, 5)
    A(3)
    B(4, 5)
      A(4)
      C(5)


A| 1 3 4
B| 2
C| 5
*/ 


/*
Maybe we have a rule that nodes cannot forward to the node they received from?
That requires nodes always have at least 2 connections, but may have B many connections as specified by user.    

With that rule, here is how we split the abc triange:
 A---B
  \ /
   C

a(1, 2, 3, 4, 5)
 b(1, 2)
  c(1)
  c(2)
 c(3, 4, 5)  
   b(3)
   b(4, 5)
     a(4)
     a(5)

a| 4 5
b| 3
c| 1 2


Now with 4
a---b
| X |
c---d

a(1, 2, 3, 4, 5)
 b(1, 2)
  c(1)
  d(2)
 c(3, 4)
   b(3)
   d(4)
 d(5)

a|
b| 3
c| 1
d| 2 4 5
*/    

/*
Hmmm, this doesn't seem to work. 
but maybe I'm thinking too much into it.
Really, we can determine if we want to do a job or not, and pass it along
if we don't. 

The pass-along vs recurse is good since we 'do work' by splitting the recursive call, and
can just pass-along a single job to the next guy if we don't want to do it. 

I also think scaling up/down should be builtin.
Add a webhook to trigger when to scale.      
*/
func (w *Worker) WorkExponentially(n borg.Node, jobIDs []int) error {
	switch len(jobs) {
		case 0:
			return nil
		case 1:
			return w.Do(jobIDs[0])
		default:
			half := len(jobs) / 2
			return n.Recurse(jobs[:half], jobs[half:])
	}
}

func (w *Worker) WorkLinearly(n borg.Node, jobIDs []int) error {
	if len(jobs) == 0 {
		return nil
	}

	go w.Do(jobs[0])
	return n.Pass(jobs[1:])
}

func (w *Worker) Do(jobID int) {
	log.Println("Doing job", jobID)
}
