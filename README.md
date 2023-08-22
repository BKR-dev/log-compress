# log-compress

### Overview

Idea is to let filebeat (ELK) aggregate the logs and use an additional output from filebeat to write to a file, read that and sent it off to a server for safe keeping. Meaning this project contains two distinct parts: 

### The Keeper

The Keeper listens for incoming requests, opens up a connection to stream data from the source. Parts will be written in a configurable size and keep track of times / parts in a sqlite that will be rotated with the log files on a weekly basis to ensure nothing will be lost due to missing db files or other connection related issues

### The Sender

Resides on the Filebeat server and streams files to the Keeper


# Progress / Currently working on

What is currenly being worked on is reflected in the feat(ure) branches which are merged into develop to be tested and then finally merged into master
