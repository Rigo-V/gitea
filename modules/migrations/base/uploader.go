// Copyright 2019 The Gitea Authors. All rights reserved.
// Copyright 2018 Jonas Franz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package base

// Uploader uploads all the informations of one repository
type Uploader interface {
	CreateRepo(repo *Repository, includeWiki bool) error
	CreateMilestones(milestones ...*Milestone) error
	CreateReleases(releases ...*Release) error
	CreateLabels(labels ...*Label) error
	CreateIssues(issues ...*Issue) error
	CreateComments(comments ...*Comment) error
	CreatePullRequests(prs ...*PullRequest) error
	Rollback() error
}
