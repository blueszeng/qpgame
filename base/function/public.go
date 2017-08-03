package function
import (
	"time"
	"math/rand"
)

func RandInt64(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	if min >= max || min==0 || max==0{
		return max
	}
	return r.Intn(max-min)+min
}