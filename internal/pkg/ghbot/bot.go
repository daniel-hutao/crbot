package ghbot

import (
	"context"
	"fmt"
	"strings"
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
		datas = append(datas, ferry.Data{
			Title:     pr.GetTitle(),
			URL:       pr.GetHTMLURL(),
			Number:    pr.GetNumber(),
			CreatedAt: pr.GetCreatedAt().Format("Jan 2 15:04"),
		})
	}

	messageChan <- ferry.Message{
		Now:   time.Now().Format("Jan 2 15:04"),
		Datas: datas,
	}

	return nil

	//if len(prList) == 0 {
	//	log.Info("no pr now.")
	//	return nil
	//}

	//var retMsg = make([]string, 0)
	//
	//for _, pr := range prList {
	//	freshPr, _, err := ghClient.PullRequests.Get(context.TODO(), ghClient.Owner, ghClient.Repo, *pr.Number)
	//	if err != nil {
	//		panic(err)
	//	}
	//	msg := dealWithPr(freshPr)
	//	if len(msg) > 0 {
	//		retMsg = append(retMsg, msg)
	//	}
	//}
	//if len(retMsg) > 0 {
	//	return retMsg
	//}
	//return nil
}

func dealWithPr(pr *github.PullRequest) string {
	title := *pr.Title
	url := *pr.URL
	reviewers := pr.RequestedReviewers

	if len(reviewers) == 0 {
		log.Infof("no reviewers")
		return fmt.Sprintf("PR: %s\nTitle: %s\n这个 pr 没有指定 reviewers。", url, title)
	}

	var reviewerList = make([]string, 0)
	for _, r := range reviewers {
		reviewerList = append(reviewerList, r.GetLogin())
	}

	return fmt.Sprintf("PR: %s\nTitle: %s\n还在等待 %s 的 review.", url, title, strings.Join(reviewerList, ", "))
}
