package counters
import(
	user "LoginAndRegisterApp/models"
	"sync")
type Counter struct {
	V   map[string]user.User
	mux sync.RWMutex
}

func Start() Counter {
	return Counter{
		V: make(map[string]user.User),
	}
}
func (c *Counter) getCounter(uPass string) user.User {

	c.mux.RLock()

	defer c.mux.RUnlock()

	return c.V[uPass]

}

func (c *Counter) InitCounter(uPass string, user user.User) {
	c.mux.Lock()

	c.V[uPass] = user
	c.mux.Unlock()

}