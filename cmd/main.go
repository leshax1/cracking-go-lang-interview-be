package main

import "log"

//var client *mongo.Client

func main() {
	log.Println("run")
	/*if !utils.VariablesCheck() {
		panic("Define all required OS Variables")
	}

	client, err := connectToMongoDB()

	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
		panic("Update mongo connection config")
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/", fileServer)
	mux.Handle("POST /buy/", middleware.SecretKeyMiddleware(&handlers.BuyHandler{DB: client}))
	mux.Handle("GET /balance/", middleware.SecretKeyMiddleware(&handlers.BalanceHandler{DB: client}))
	mux.Handle("GET /pnl/", &handlers.PnlHandler{DB: client})
	mux.Handle("GET /total/", &handlers.TotalHandler{DB: client})

	log.Println("Starting server on :3001...")
	err = http.ListenAndServe(":3001", mux)
	log.Println(err)*/

}

/*func connectToMongoDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("okxMongoConnectionString"))

	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB")
	return client, nil
}*/
