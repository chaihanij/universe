### Universe
Universe is utility library 

#### Usage
```go
*package environ*ment

var (
	Url       string
	Port string
)

func Init() {
	_ = env.Require("URL", "simple url")
	_ = env.Require("PORT", "listening port")
	
	// Call this to ensure every required env is available, if not, panic
	env.Assert()
}
```