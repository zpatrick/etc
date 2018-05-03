# Quality

The goal of this Tech Talk is to present some ideas and processes that I have found useful in the pursuit of writing high-quality code. 
My hope is that others will try some of these ideas, and then pick and choose those that they find work well for them. 
I also think it's worth mentioning what this Tech Talk is _not_ trying to be:
* A comprensive list of all ideas, processes, etc. that lend themselves to writing high-quality code
* A formula (e.g. _"If you follow these suggestions, then you will have high-quality code"_)
* For everybody. As stated above, these are some things _I_ have found useful. 
I believe most people will find at least some of the ideas presented here useful, but maybe they won't, and that's ok!


# Terms
There are a lot of great articles out there that go into detail about the characteristics of high-quality code. 
I wanted to take a more practical approach with this talk, but I think it's worth (briefly) going over some common terms that are worth knowing and thinking about: 

**KISS**: An acronym for "Keep it simple, stupid". 
The KISS principle states that most systems work best if they are kept simple rather than made complicated; 
therefore simplicity should be a key goal in design and unnecessary complexity should be avoided.

**DRY**: An acronym for "Don't repeat yourself". 
The DRY principle aims at reducing repetition and redundancy by using patterns and abstractions. 

**Coupling**: The degree to which modules depend on each other to function correctly.
It is typically desirable to have low levels of coupling. 

**Cohesion**: The degree to which the elements inside of a module belong together. 
It is typically desirable to have high levels of cohesion. 

**Extensibility**: The degree to which a system can be extended for additional functionality, and the amount of effort required to implement the extension. 

# Approach 
This section lists some ideas for approaching software design and implementation. 
I put a little bit of thought into the order of these ideas, 
but the order is really meant to be more a suggestion than a requirement.

### Care
I know, this is one of those "duh" suggestions, but hear me out: 
The best tips in the world aren't going to help if you don't care about the quality of your own work. 
We all have deadlines, projects we don't like, lack of sleep, bad codebases ...the list goes on.
There are a million excuses not to care, but I think it's almost impossible to write high-quality code without caring. 
On the bright side, if you're reading this, there's a good chance you at least care a little bit. 
Being in the proper mindset before writing code can go a long way. 

### Know the Language 
Become an expert (or at least very good) in the language you are working in. 
I know that's easier said than done, but it's crucial. 
Here are some things that I've found useful when trying to master a language:
* Read a book written for that language (bonus if there is an "official" book for the language)
* Read articles and tutorials online (Google and Reddit are your friends there)
* Pair program with someone who is already an expert on the language (obviously, not applicable for everyone)
* View the source code of popular projects written in that language
* Practice writing projects in that language ([CodeKata](http://codekata.com/) is one of many good sources for practice ideas)

### Think Abstractly

### Start at the End

### Iterate, Iterate, Iterate
- have a large toolbelt (know tricks, see other's implementations, etc.)
- think abstractly
- start at the end (consumers and testers, code IS ux)
- iterate, iterate, iterate

# Example
You want to make a suite of card games (poker, blackjack, etc.). 
You realize that you will have some common requirements when it comes to card, e.g.
	- how to express cards as objects
	- how to express a deck of cards as an object
	- shuffling a deck of cards
	- dealing from a deck of cards
	- how to express a hand of cards as an object             
