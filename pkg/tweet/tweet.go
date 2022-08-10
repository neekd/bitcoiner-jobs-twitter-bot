package tweet

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets"
	"github.com/michimani/gotwi/tweets/types"
	"github.com/mmcdole/gofeed"
	log "github.com/sirupsen/logrus"
)

type Agent struct {
	Client *gotwi.GotwiClient
}

func NewAgent(oAuthToken, oAuthTokenSecret string) (*Agent, error) {
	in := &gotwi.NewGotwiClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           oAuthToken,
		OAuthTokenSecret:     oAuthTokenSecret,
	}

	c, err := gotwi.NewGotwiClient(in)
	if err != nil {
		return nil, err
	}

	return &Agent{
		Client: c,
	}, nil
}

func (ta *Agent) SendTweet(text string) error {
	p := &types.ManageTweetsPostParams{
		Text: gotwi.String(text),
	}

	res, err := tweets.ManageTweetsPost(context.Background(), ta.Client, p)
	if err != nil {
		return err
	}

	log.Info(
		fmt.Sprintf(
			"Tweet Sent: [%s] %s",
			gotwi.StringValue(res.Data.ID),
			gotwi.StringValue(res.Data.Text),
		),
	)
	return nil
}

func FormatFeedItem(item *gofeed.Item, hashtags []string) string {
	city, description := formatDescription(item.Description)
	if city != "" {
		return fmt.Sprintf(
			"New job posted in %s for %s!\n\n%s- Apply here: %s\n\n%s",
			city, item.Title, description, item.Link, strings.Join(hashtags, " "),
		)
	} else {
		return fmt.Sprintf(
			"New job posted for %s!\n\n%s- Apply here: %s\n\n%s",
			item.Title, description, item.Link, strings.Join(hashtags, " "),
		)
	}

}

func formatDescription(description string) (string, string) {
	var rgx = regexp.MustCompile(`\((.*?)\)`)
	city := rgx.FindStringSubmatch(description)
	des := rgx.Split(description, 2)
	if len(city) < 2 {
		return "", des[0]
	}
	return strings.TrimSpace(city[1]), des[0]
}
