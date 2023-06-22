package pkg

import "fmt"

type Router struct{}

func callClosure(closure func()) {
	closure()
}

// Route is the obs_pipelines router
func (r *Router) Route() {
	callClosure(func() {
		fmt.Println("getConfigVersion:", r.getConfigVersion)
	})
}

func (r *Router) getConfigVersion() {}
