# Swan Twitter Bot

Automatically tweet from the Bitcoiner Jobs RSS feed.

## How to use

Use the `config.example.yaml` to create a new config file named `config.yaml` with the required secrets, RSS feed URL, and hashtags. 

Build and run the Docker container: `docker build -t bitcoiner-jobs-twitter-bot . && docker run bitcoiner-jobs-twitter-bot` 

That's it! The bot should be up and running and you should see the last post from the RSS feed in your twitter feed. The bot will check back in every 5 minutes to see if any new jobs have been posted. 