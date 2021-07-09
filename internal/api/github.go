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
			members, _, err := ghClient.Organizations.ListMembers(context.Background(), org, &github.ListMembersOptions{ListOptions: github.ListOptions{PerPage: 10}})
			if err != nil {
				// TODO: Implementar
				fmt.Println("ghClient.Organizations.ListMembers err:", err)
			}
			fmt.Println(len(members))

			for _, member := range members {
				repos, _, err := ghClient.Repositories.List(context.Background(), *member.Login, &github.RepositoryListOptions{})
				if err != nil {
					// TODO: Implementar
					fmt.Println("ghClient.Repositories.List err:", err)
				}

				languages := []models.Language{}
				stargazersCount := 0
				contributtedCount := 0
				for _, repo := range repos {
					if len(repo.GetLanguage()) > 0 {
						languages = append(languages, models.Language{Name: repo.GetLanguage()})
					}
					stargazersCount += repo.GetStargazersCount()

					contributorsStats, _, err := ghClient.Repositories.ListContributorsStats(context.Background(), repo.GetOwner(), repo.GetName())
					if err != nil {
						// TODO: Implementar
						fmt.Println("ghClient.Repositories.ListContributorsStats err:", err)
					}

					for _, stats := range contributorsStats {
						if stats.GetAuthor().GetID() == member.GetID() {
							contributtedCount += stats.GetTotal()
						}
					}
				}

				ranking := models.Ranking{
					Name:              member.GetName(),
					AvatarURL:         member.GetAvatarURL(),
					StargazersCount:   stargazersCount,
					FollowersCount:    member.GetFollowers(),
					ProjectsCount:     member.GetPublicRepos(),
					ContributtedCount: len(repos),
					Languages:         languages,
				}
				rankingResource.Create(ranking)
			}

			time.Sleep(24 * time.Hour)
			ghLoopCh <- struct{}{}
		}()
		<-ghLoopCh
	}
}
