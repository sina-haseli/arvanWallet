# arvanWallet

This project is my code challenge for ArvanCloud company interview.<br>
It uses redis pub/sub as a queue for communicate with wallet to apply credit.<br>
there is 1 api for discount:<br>
1. `/api/balance/:user_id`: get the balance of a user<br>
For run project you can use the following command:<br>
`$ go run main.go serve` or `$ wallet serve`
