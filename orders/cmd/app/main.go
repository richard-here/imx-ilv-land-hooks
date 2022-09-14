package main

import "github.com/richard-here/imx-ilv-land-hooks/orders/internal/imx"

func main() {
	imx.FindById("1")
	imx.FindById("2")
	// cfg := mysql.Config{
	// 	User:                 os.Getenv("DBUSER"),
	// 	Passwd:               os.Getenv("DBPASS"),
	// 	Net:                  "tcp",
	// 	Addr:                 "127.0.0.1:3307",
	// 	DBName:               "orders",
	// 	AllowNativePasswords: true,
	// }

	// db, err := sql.Open("mysql", cfg.FormatDSN())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// pingErr := db.Ping()
	// if pingErr != nil {
	// 	log.Fatal(pingErr)
	// }

	// fmt.Println("Connected to `orders` db")
}
