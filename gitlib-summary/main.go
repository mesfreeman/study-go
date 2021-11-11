package main

import (
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
	"github.com/tidwall/gjson"
)

var (
	gitLibHost   = ""                       // 仓库域名
	commiter     = ""                       // 提交者名称
	privateToken = ""                       // 访问令牌
	summaryInfo  = make(map[string]summary) // 统计信息
)

// 项目信息
type project struct {
	ID   int    `json:"id"`
	Name string `json:"path_with_namespace"`
}

// 提交信息
type commit struct {
	ID           string `json:"id"`
	CommiterName string `json:"committer_name"`
}

// 统计信息
type summary struct {
	Additions int64 `json:"additions"`
	Deletions int64 `json:"deletions"`
	Total     int64 `json:"total"`
}

// 统计GitLab仓库的master分支下某人提交的代码量
// 注：只统计了master分支，维度为行数，修改操作其实是先删除后新增的操作
func main() {
	fmt.Println("请输入GitLab域名、提交者及访问令牌，以空格分隔，如：https://www.gitlab.com 张三 TVEX-waMuPHhRryu9_kN")
	fmt.Scanf("%s %s %s", &gitLibHost, &commiter, &privateToken)
	if gitLibHost == "" || commiter == "" || privateToken == "" {
		fmt.Println("格式错误")
		return
	}

	for page := 1; page < 20; page++ {
		api := fmt.Sprintf(`%s/api/v4/projects/?private_token=%s&page=%d&per_page=200`, gitLibHost, privateToken, page)
		_, body, errors := gorequest.New().Get(api).End()
		if errors != nil {
			fmt.Printf("接口请求出错，Api：%s，错误信息：%s\n\n", api, errors[0].Error())
			continue
		}

		if body == "[]" {
			fmt.Printf("执行结束，共 %d 页\n\n", page-1)
			return
		}

		projects := make([]project, 0)
		json.Unmarshal([]byte(body), &projects)
		for _, project := range projects {
			fmt.Printf("项目ID：%d，项目名称：%s，开始获取master分支代码量...\n", project.ID, project.Name)
			additions, deletions, total := commitIds(project.ID, "master")
			fmt.Printf("项目ID：%d，项目名称：%s，提交代码量获取成功，新增 %d 行，删除 %d 行, 总变更 %d 行\n\n", project.ID, project.Name, additions, deletions, total)

			// 放入统计Map中
			summaryInfo[project.Name] = summary{
				Additions: additions,
				Deletions: deletions,
				Total:     total,
			}
		}
	}

	// 打印统计数据
	fmt.Println("=========== 统计结束如下 ===========")
	fmt.Printf("项目名称\t新增行数\t删除行数\t，总变更行数\n")
	for projectName, item := range summaryInfo {
		fmt.Printf("%s\t%d\t%d\t%d\n", projectName, item.Additions, item.Deletions, item.Total)
	}
}

// 获取分支提交记录集合
func commitIds(projectId int, branchName string) (additions, deletions, total int64) {
	// 遍历分页
	for page := 1; page < 10000000; page++ {
		api := fmt.Sprintf(`%s/api/v4/projects/%d/repository/commits?ref_name=%s&private_token=%s&page=%d&per_page=200`, gitLibHost, projectId, branchName, privateToken, page)
		_, body, errors := gorequest.New().Get(api).End()
		if errors != nil {
			fmt.Printf("接口请求出错，Api：%s，错误信息：%s\n\n", api, errors[0].Error())
			continue
		}

		if body == "[]" {
			break
		}

		commits := make([]commit, 0)
		json.Unmarshal([]byte(body), &commits)
		for _, commit := range commits {
			if commit.CommiterName != commiter {
				continue
			}

			singleAdditions, singleDeletions, singleTotal := codeSummary(projectId, commit.ID)
			additions += singleAdditions
			deletions += singleDeletions
			total += singleTotal
		}
	}

	return additions, deletions, total
}

// 获取提交的代码量统计
func codeSummary(projectId int, commitId string) (additions, deletions, total int64) {
	api := fmt.Sprintf(`%s/api/v4/projects/%d/repository/commits/%s?private_token=%s`, gitLibHost, projectId, commitId, privateToken)
	_, body, errors := gorequest.New().Get(api).End()
	if errors != nil {
		fmt.Printf("接口请求出错，Api：%s，错误信息：%s\n\n", api, errors[0].Error())
		return
	}

	additions = gjson.Get(body, "stats.additions").Int()
	deletions = gjson.Get(body, "stats.deletions").Int()
	total = gjson.Get(body, "stats.total").Int()
	return
}
