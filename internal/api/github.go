package api

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/github"
	"github.com/joselitofilho/golang-echo-apigithub/internal/models"
	"github.com/joselitofilho/golang-echo-apigithub/internal/resources"
)

func GHInfos(ghLoopCh chan struct{}, ghClient *github.Client, org string, rankingResource resources.RankingResource) {
	for {
		go func() {
			fmt.Println("Retrieve github informations...")

			members, _, err := ghClient.Organizations.ListMembers(context.Background(), org, &github.ListMembersOptions{ListOptions: github.ListOptions{PerPage: 10}})
			if err != nil {
				// TODO: Implementar
				fmt.Println("ghClient.Organizations.ListMembers err:", err)
			}

			repos, _, err := ghClient.Repositories.ListByOrg(context.Background(), org, &github.RepositoryListByOrgOptions{})
			if err != nil {
				// TODO: Implementar
				fmt.Println("ghClient.Repositories.ListByOrg err:", err)
			}

			contributions := map[string]int{}
			for _, member := range members {
				contributions[member.GetLogin()] = 0
			}

			languages := []models.Language{}
			stargazersCount := 0
			for _, repo := range repos {
				if len(repo.GetLanguage()) > 0 {
					languages = append(languages, models.Language{Name: repo.GetLanguage()})
				}
				stargazersCount += repo.GetStargazersCount()

				contributors, _, err := ghClient.Repositories.ListContributors(context.Background(), repo.GetOwner().GetLogin(), repo.GetName(), &github.ListContributorsOptions{})
				if err != nil {
					// TODO: Implementar
					fmt.Println("ghClient.Repositories.ListContributors err:", err)
				}

				for _, contributor := range contributors {
					contributions[contributor.GetLogin()] += contributor.GetContributions()
				}
			}

			for _, member := range members {
				ranking := models.Ranking{
					Name:              member.GetLogin(),
					AvatarURL:         member.GetAvatarURL(),
					StargazersCount:   stargazersCount,
					FollowersCount:    member.GetFollowers(),
					ProjectsCount:     member.GetPublicRepos(),
					ContributtedCount: contributions[member.GetLogin()],
					Languages:         languages,
				}
				if err := rankingResource.Create(&ranking); err != nil {
					// TODO: Implementar
					fmt.Println("rankingResource.Create err:", err)
				}
			}

			time.Sleep(24 * time.Hour)
			ghLoopCh <- struct{}{}
		}()
		<-ghLoopCh
	}
}
