# gobandit
A golang library for solving multi armed bandit problem which can optimize your business choice on the fly without A/B testing.

## Thompson Sampling

```go
package main

import (
	"time"
	
	"github.com/rueian/gobandit/thompson"
)

func init() {
	thompson.Seed(time.Now().Unix())
}

func main() {
    // load your candidate statistics, for example:
    candidates := [][2]float64{
        {1000, 500}, // candidate 0 => likeCount: 1000, dislikeCount: 500
        {500, 1000}, // candidate 1 => likeCount: 500, dislikeCount: 1000
    }
    
    for {
        // let Thompson Sampling choose a candidate based on the statistics
        elected, _ := thompson.Choose(candidates)
        
        // use the elected one, and collect feedback to update the statistic for next choose
        if peopleLike(elected) {
            candidates[elected][0]++
        } else {
            candidates[elected][1]++
        }
        
        if done() {
            break
        }
    }
}
```