# Interflow Wallet API

[Interflow Api](https://github.com/interflowrepo/interflow-api) communicates directly with this API to create wallets, execute transactions and query transaction status.

This API communicates directly with the [Interflow Api](https://github.com/interflowrepo/interflow-api) via webhooks to notify when transactions have completed.

To run this application locally on your machine without having to use docker and connecting directly to our live database, run the command:

env $(cat .env | grep -v "#" | xargs) go run main.go

We leave the .env files open to make testing easier!

*This app is live

[**To test all Interflow features quickly and easily, follow the instructions in the Interflow repository**](https://github.com/interflowrepo/interflow/blob/main/README.md)



**NOTE: We had some issues with gits repositories and conflicts. Over time we have made some changes and created new repositories to ease and avoid problems, [the history and explanation of this you can find in this link](https://github.com/interflowrepo/interflow-api/blob/main/HACKATON-HISTORY.md)**
