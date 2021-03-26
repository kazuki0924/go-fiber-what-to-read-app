package infrastructure

type RDB interface {
	// GET(uri string, f func(w http.ResponseWriter, req *http.Request))
	// POST(uri string, f func(w http.ResponseWriter, req *http.Request))
	// SERVE(port string)
	InitRDB()
	CloseRDB()
}
