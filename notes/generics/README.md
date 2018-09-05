# Generics in Go

## High Level Thoughts
* A lot of problems generics try to solve can already be solved by interfaces.
Developers would just rather push the tediousness of implementing those interfaces to the compiler.
This is what the 'contract' objects try to solve in the current design.  
* Generics sometimes branch from interfaces in that sometimes the underlying type is important. 
Take a generic cache for example - what should the return type be for the "cache.Get" function? 


## Validation Scenarios
* Create a single `sum` function for all types. 
* Create a single `sort` function for any slice.
* Create a single `cache` struct that stores any type. 
  
## Ideas
* Create a new builtin type, `generic`, that represents all types (similar to `interface{}`).
* Allow users to define functions on the `generic` type to implement the desired interface. 
We will call these function definitions `contracts`.
The compiler will need special rules on the `contract` to determine which types can and cannot use the contract.    
* Create a new keyword, `infer`, which means "the same type as XXX". 
This will probably need many rules around it as there is a lot of questions to answer:
  - What happens when the function returns different types?
  - What happens when the function requires different inferred types?
  - How would one declare a new inferred type within a function body? 
