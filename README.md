# Minesweeper

This is a classic game of Minesweeper (yes, the one you know from way back in those win95 days). I'm using it as an opportunity to learn new languages & frameworks. In particular, the backend for this game will be done with Go, using any tools I find interesting along the way, and a frontend with Vue.
I'll be setting myself a deadline: to complete this over the weekend, with whatever time I find in between. This will be a good indicator to myself of the learning curve of the Go ecosystem, and it's ease of use.

## Short term plans

- ✔ ~~Template Go API project, with a featured framework that includes auto generating docs (Swagger, or Swagger like)~~
- ✔ ~~Template Vue project~~
- ✔ ~~Basic API endpoints to create game (store game matrix in db)~~
- ✔ ~~Basic API endpoints to play the game~~
- ✔ ~~Basic UI to visualize the game board and interact with the backend~~
- User auth workflow
- Deploy to AWS
- Setup pipelines (might do with github actions instead of aws pipelines, easier to setup)
- DTO Validation
- Testing (normally this'd be much higher up the list, but this is a prototype)
- Nice FE error handling

## Random comments

- The Go syntax is very weird, nothing like I've ever used before, BUT, its conventions are very strongly enforced (e.g. uppercase field access, unused imports)
- This seems like a project well suited for a NoSQL db, at least for a quick prototype
- Since there's no generics, decorators, or any sort of metaprogramming, ORMs and Swagger seem to rely on comments, which are very unsafe and many times left outdated
- This reminds me a lot of C, specifically the fact that since the syntax is not too complex yet programs can have a lot of freedom, it leaves that freedom to the developer to structure code as they see fit. That's very cool, but also it makes me unsure of what's the best approach to take
- Clicks (i.e. state changes) return the whole board. This could be much more optimized, but because clicking can change the game state, or more than one cell, this was the fast approach to an MVP