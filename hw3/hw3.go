package hw3

func SetForm(args ...string) map[string]struct{} {
	resp := make(map[string]struct{})
	for _, arg := range args {
		resp[arg] = struct{}{}
	}
	return resp
}
