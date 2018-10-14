# SYNOPSYS

A simple example of listener for jira. Sends a telegram message when webhook from jira arrived. 
Supports only eventTypes:
 - jira:issue_created
 - comment_created


# INSTALL

	$ make submodule_check
	$ make

A binary will be in bin directory

# Usage

Please make a configuration file in YAML. There is example:

```
listenport: "80"
telegramtoken: "TOKEN"
telegramchatid: "CHATID"
```

Run listener

	$ ./bin/jira-telegram-listener -config ./jira-telegram-listener.yml


# Licence

MIT
