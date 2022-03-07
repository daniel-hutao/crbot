package ghbot

import (
	"context"
	"time"

	"github.com/daniel-hutao/crbot/internal/pkg/ferry"

	"github.com/google/go-github/v42/github"

	gh "github.com/daniel-hutao/crbot/internal/pkg/github"
	"github.com/daniel-hutao/crbot/internal/pkg/log"
)

const (
	owner = "merico-dev"
	//owner = "daniel-hutao"
	repo = "stream"
)

func GetMsg(messageChan ferry.MessageChan) error {
	ghOption := gh.Option{
		Owner:    owner,
		Repo:     repo,
		NeedAuth: true,
		WorkPath: "",
	}
	ghClient, err := gh.NewClient(&ghOption)
	if err != nil {
		return err
	}

	prList, _, err := ghClient.PullRequests.List(context.TODO(), ghClient.Owner, ghClient.Repo, &github.PullRequestListOptions{})
	if err != nil {
		return err
	}

	if len(prList) == 0 {
		log.Infof("No pr now.")
		return nil
	}

	datas := make([]ferry.Data, 0)
	for _, pr := range prList {
		var rr = make([]ferry.Reviewer, 0)
		for _, u := range pr.RequestedReviewers {
			rr = append(rr, ferry.Reviewer(u.GetLogin()))
		}
		datas = append(datas, ferry.Data{
			Title:              pr.GetTitle(),
			User:               pr.GetUser().GetLogin(),
			URL:                pr.GetHTMLURL(),
			Number:             pr.GetNumber(),
			CreatedAt:          pr.GetCreatedAt().Format("Jan 2 15:04"),
			RequestedReviewers: rr,
		})

		//prWithDetail, _, err := ghClient.PullRequests.Get(context.TODO(), ghClient.Owner, ghClient.Repo, pr.GetNumber())
		//if err != nil {
		//	return err
		//}
		//
		//_ = prWithDetail

	}

	messageChan <- ferry.Message{
		Now:   time.Now().Format("Jan 2 15:04"),
		Datas: datas,
	}

	return nil
}
