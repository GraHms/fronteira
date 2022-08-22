package main

func main() {

	authPolicy := NewAuthPolicy("data/example.yaml")
	serviceDomain := authPolicy.Spec.Target
	ops := authPolicy.Spec.Operation
	r := NewRequest(serviceDomain)
	for i := range ops {
		r.Handler(ops[i])
	}

	r.Run(":8280") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
