package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProjectResponse struct {
	Results []Project `json:"results"`
}

type Issues struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Assignees []string `json:"assignees"`
	Priority string   `json:"priority"`
}

type IssueResponse struct {
	Results []Issues `json:"results"`
}

func main() {
	apiKey := "plane_api_75401edcd20343fdb616563659adcea8"
    assigneeID := "424cada8-95ca-4710-8c93-66ce6e71c7b0"

    projectURL := "https://api.plane.so/api/v1/workspaces/ikhsan-workspace/projects"
    req, _ := http.NewRequest("GET", projectURL, nil)
    req.Header.Add("x-api-key", apiKey)


    res, err := http.DefaultClient.Do(req)
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()
    projectBody, _ := ioutil.ReadAll(res.Body)


    var projectResp ProjectResponse
    if err := json.Unmarshal(projectBody, &projectResp); err != nil {
        panic(err)
    }

    for _, project := range projectResp.Results {
        issueURL := fmt.Sprintf("https://api.plane.so/api/v1/workspaces/ikhsan-workspace/projects/%s/issues", project.ID)
        reqIssue, _ := http.NewRequest("GET", issueURL, nil)
        reqIssue.Header.Add("x-api-key", apiKey)


        resIssue, err := http.DefaultClient.Do(reqIssue)
        if err != nil {
            fmt.Printf("❌ Gagal ambil issues dari project %s\n", project.Name)
            continue
        }
        defer resIssue.Body.Close()
        issueBody, _ := ioutil.ReadAll(resIssue.Body)


        var issueResp IssueResponse
        if err := json.Unmarshal(issueBody, &issueResp); err != nil {
            fmt.Printf("❌ Error parsing JSON dari project %s\n", project.Name)
            continue
        }


        // Step 3: Filter by assignee
        for _, issue := range issueResp.Results {
            for _, assignee := range issue.Assignees {
                if assignee == assigneeID {
                    fmt.Printf("✅Project Name : [%s] Issue : %s (%s) - Priority: %s\n", project.Name, issue.Name, issue.ID, issue.Priority)
                }
            }
        }
    }
}
