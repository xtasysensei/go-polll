# github.com/xtasysensei/go-poll

Welcome to the go-poll generated docs.

## Routes

<details>
<summary>`/`</summary>

- [LoggingMiddleware]()
- [Recoverer]()
- [ChangeMethod]()
- **/**
	- _GET_
		- [Index]()

</details>
<details>
<summary>`/ping`</summary>

- [LoggingMiddleware]()
- [Recoverer]()
- [ChangeMethod]()
- **/ping**
	- _GET_
		- [Health]()

</details>
<details>
<summary>`/v1/auth/login`</summary>

- [LoggingMiddleware]()
- [Recoverer]()
- [ChangeMethod]()
- **/v1**
	- **/auth**
		- **/login**
			- _POST_
				- [HandleLogin]()

</details>
<details>
<summary>`/v1/auth/register`</summary>

- [LoggingMiddleware]()
- [Recoverer]()
- [ChangeMethod]()
- **/v1**
	- **/auth**
		- **/register**
			- _POST_
				- [HandleRegister]()

</details>
<details>
<summary>`/v1/polls`</summary>

- [LoggingMiddleware]()
- [Recoverer]()
- [ChangeMethod]()
- **/v1**
	- **/polls**
		- **/**
			- _POST_
				- [HandleCreatePoll]()
			- _GET_
				- [RetrieveAllPolls]()

</details>
<details>
<summary>`/v1/polls/{pollId}`</summary>

- [LoggingMiddleware]()
- [Recoverer]()
- [ChangeMethod]()
- **/v1**
	- **/polls**
		- **/{pollId}**
			- _GET_
				- [RetrievePollByID]()

</details>
<details>
<summary>`/v1/polls/{pollId}/vote`</summary>

- [LoggingMiddleware]()
- [Recoverer]()
- [ChangeMethod]()
- **/v1**
	- **/polls**
		- **/{pollId}/vote**
			- _POST_
				- [HandleCastVote]()

</details>

Total # of routes: 7

