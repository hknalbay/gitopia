package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gitopia/gitopia/x/gitopia/types"
	"github.com/gitopia/gitopia/x/gitopia/utils"
)

func (k msgServer) CreateIssue(goCtx context.Context, msg *types.MsgCreateIssue) (*types.MsgCreateIssueResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetUser(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("creator (%v) doesn't exist", msg.Creator))
	}

	repository, found := k.GetRepository(ctx, msg.RepositoryId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("repository id (%d) doesn't exist", msg.RepositoryId))
	}
	repository.IssuesCount += 1

	createdAt := ctx.BlockTime().Unix()
	closedAt := time.Time{}.Unix()

	var issue = types.Issue{
		Creator:       msg.Creator,
		Iid:           repository.IssuesCount,
		Title:         msg.Title,
		State:         types.Issue_OPEN,
		Description:   msg.Description,
		CommentsCount: 0,
		RepositoryId:  msg.RepositoryId,
		Weight:        0,
		CreatedAt:     createdAt,
		UpdatedAt:     createdAt,
		ClosedAt:      closedAt,
	}

	if len(msg.Assignees) > 0 || len(msg.LabelIds) > 0 {
		var organization types.Organization

		if repository.Owner.Type == types.RepositoryOwner_ORGANIZATION {
			organization, found = k.GetOrganization(ctx, repository.Owner.Id)
			if !found {
				return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("organization (%v) doesn't exist", repository.Owner.Id))
			}
		}

		if !utils.HaveIssuePermission(repository, msg.Creator, organization) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("user (%v) doesn't have permission to perform this operation", msg.Creator))
		}

		for _, a := range msg.Assignees {
			_, found := k.GetUser(ctx, a)
			if !found {
				return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("assignee (%v) doesn't exist", a))
			}
		}
		issue.Assignees = msg.Assignees
		for _, labelId := range msg.LabelIds {
			if i, exists := utils.RepositoryLabelIdExists(repository.Labels, labelId); exists {
				issue.Labels = append(issue.Labels, repository.Labels[i].Id)
			} else {
				return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("label id (%v) doesn't exists in repository", labelId))
			}
		}
	}

	id := k.AppendIssue(
		ctx,
		issue,
	)

	var repositoryIssue = types.RepositoryIssue{
		Iid: repository.IssuesCount,
		Id:  id,
	}

	repository.Issues = append(repository.Issues, &repositoryIssue)

	k.SetRepository(ctx, repository)

	return &types.MsgCreateIssueResponse{
		Id:  id,
		Iid: issue.Iid,
	}, nil
}

func (k msgServer) UpdateIssue(goCtx context.Context, msg *types.MsgUpdateIssue) (*types.MsgUpdateIssueResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetUser(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("creator (%v) doesn't exist", msg.Creator))
	}

	issue, found := k.GetIssue(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("issue id (%d) doesn't exist", msg.Id))
	}

	if msg.Creator != issue.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	issue.Title = msg.Title
	issue.Description = msg.Description
	issue.Weight = msg.Weight
	issue.UpdatedAt = ctx.BlockTime().Unix()
	issue.Assignees = msg.Assignees

	k.SetIssue(ctx, issue)

	return &types.MsgUpdateIssueResponse{}, nil
}

func (k msgServer) UpdateIssueTitle(goCtx context.Context, msg *types.MsgUpdateIssueTitle) (*types.MsgUpdateIssueTitleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetUser(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("creator (%v) doesn't exist", msg.Creator))
	}

	issue, found := k.GetIssue(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("issue id (%d) doesn't exist", msg.Id))
	}

	if issue.Title == msg.Title {
		return &types.MsgUpdateIssueTitleResponse{}, nil
	}

	if msg.Creator != issue.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	oldTitle := issue.Title

	issue.Title = msg.Title
	issue.UpdatedAt = ctx.BlockTime().Unix()
	issue.CommentsCount += 1

	var comment = types.Comment{
		Creator:     "GITOPIA",
		ParentId:    msg.Id,
		CommentIid:  issue.CommentsCount,
		Body:        utils.UpdateTitleCommentBody(msg.Creator, oldTitle, issue.Title),
		System:      true,
		CreatedAt:   issue.UpdatedAt,
		UpdatedAt:   issue.UpdatedAt,
		CommentType: types.Comment_ISSUE,
	}

	id := k.AppendComment(
		ctx,
		comment,
	)

	issue.Comments = append(issue.Comments, id)

	k.SetIssue(ctx, issue)

	return &types.MsgUpdateIssueTitleResponse{}, nil
}

func (k msgServer) UpdateIssueDescription(goCtx context.Context, msg *types.MsgUpdateIssueDescription) (*types.MsgUpdateIssueDescriptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetUser(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("creator (%v) doesn't exist", msg.Creator))
	}

	issue, found := k.GetIssue(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("issue id (%d) doesn't exist", msg.Id))
	}

	if issue.Description == msg.Description {
		return &types.MsgUpdateIssueDescriptionResponse{}, nil
	}

	if msg.Creator != issue.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	issue.Description = msg.Description
	issue.UpdatedAt = ctx.BlockTime().Unix()
	issue.CommentsCount += 1

	var comment = types.Comment{
		Creator:     "GITOPIA",
		ParentId:    msg.Id,
		CommentIid:  issue.CommentsCount,
		Body:        utils.UpdateDescriptionCommentBody(msg.Creator),
		System:      true,
		CreatedAt:   issue.UpdatedAt,
		UpdatedAt:   issue.UpdatedAt,
		CommentType: types.Comment_ISSUE,
	}

	id := k.AppendComment(
		ctx,
		comment,
	)

	issue.Comments = append(issue.Comments, id)

	k.SetIssue(ctx, issue)

	return &types.MsgUpdateIssueDescriptionResponse{}, nil
}

func (k msgServer) ToggleIssueState(goCtx context.Context, msg *types.MsgToggleIssueState) (*types.MsgToggleIssueStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetUser(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("creator (%v) doesn't exist", msg.Creator))
	}

	issue, found := k.GetIssue(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("issue id (%d) doesn't exist", msg.Id))
	}

	repository, found := k.GetRepository(ctx, issue.RepositoryId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("repository id (%d) doesn't exist", issue.RepositoryId))
	}

	if msg.Creator != issue.Creator {
		var organization types.Organization
		if repository.Owner.Type == types.RepositoryOwner_ORGANIZATION {
			organization, found = k.GetOrganization(ctx, repository.Owner.Id)
			if !found {
				return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("organization (%v) doesn't exist", repository.Owner.Id))
			}
		}

		if !utils.HaveIssuePermission(repository, msg.Creator, organization) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("user (%v) doesn't have permission to perform this operation", msg.Creator))
		}
	}

	if issue.State == types.Issue_OPEN {
		issue.State = types.Issue_CLOSED
		issue.ClosedBy = msg.Creator
		issue.ClosedAt = ctx.BlockTime().Unix()
	} else if issue.State == types.Issue_CLOSED {
		issue.State = types.Issue_OPEN
		issue.ClosedBy = string("")
		issue.ClosedAt = time.Time{}.Unix()
	} else {
		/* TODO: specify error */
		return nil, sdkerrors.Error{}
	}

	issue.CommentsCount += 1
	issue.UpdatedAt = ctx.BlockTime().Unix()

	var comment = types.Comment{
		Creator:     "GITOPIA",
		ParentId:    msg.Id,
		CommentIid:  issue.CommentsCount,
		Body:        utils.IssueToggleStateCommentBody(msg.Creator, issue.State),
		System:      true,
		CreatedAt:   issue.UpdatedAt,
		UpdatedAt:   issue.UpdatedAt,
		CommentType: types.Comment_ISSUE,
	}

	id := k.AppendComment(
		ctx,
		comment,
	)

	issue.Comments = append(issue.Comments, id)

	k.SetIssue(ctx, issue)

	return &types.MsgToggleIssueStateResponse{
		State: issue.State.String(),
	}, nil
}

func (k msgServer) AddIssueAssignees(goCtx context.Context, msg *types.MsgAddIssueAssignees) (*types.MsgAddIssueAssigneesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetUser(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("creator (%v) doesn't exist", msg.Creator))
	}

	issue, found := k.GetIssue(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("issue id (%d) doesn't exist", msg.Id))
	}

	repository, found := k.GetRepository(ctx, issue.RepositoryId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("repository id (%d) doesn't exist", issue.RepositoryId))
	}
	var organization types.Organization

	if repository.Owner.Type == types.RepositoryOwner_ORGANIZATION {
		organization, found = k.GetOrganization(ctx, repository.Owner.Id)
		if !found {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("organization (%v) doesn't exist", repository.Owner.Id))
		}
	}

	if !utils.HaveIssuePermission(repository, msg.Creator, organization) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("user (%v) doesn't have permission to perform this operation", msg.Creator))
	}

	if len(issue.Assignees)+len(msg.Assignees) > 10 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "issue can't have more than 10 assignees")
	}

	for _, a := range msg.Assignees {
		if _, exists := utils.AssigneeExists(issue.Assignees, a); exists {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("assignee (%v) already assigned", a))
		}
		_, found := k.GetUser(ctx, a)
		if !found {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("assignee (%v) doesn't exist", a))
		}
		issue.Assignees = append(issue.Assignees, a)
	}

	issue.CommentsCount += 1
	issue.UpdatedAt = ctx.BlockTime().Unix()

	var comment = types.Comment{
		Creator:     "GITOPIA",
		ParentId:    msg.Id,
		CommentIid:  issue.CommentsCount,
		Body:        utils.AddAssigneesCommentBody(msg.Creator, msg.Assignees),
		System:      true,
		CreatedAt:   issue.UpdatedAt,
		UpdatedAt:   issue.UpdatedAt,
		CommentType: types.Comment_ISSUE,
	}

	id := k.AppendComment(
		ctx,
		comment,
	)

	issue.Comments = append(issue.Comments, id)

	k.SetIssue(ctx, issue)

	return &types.MsgAddIssueAssigneesResponse{}, nil
}

func (k msgServer) RemoveIssueAssignees(goCtx context.Context, msg *types.MsgRemoveIssueAssignees) (*types.MsgRemoveIssueAssigneesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetUser(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("creator (%v) doesn't exist", msg.Creator))
	}

	issue, found := k.GetIssue(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("issue id (%d) doesn't exist", msg.Id))
	}

	repository, found := k.GetRepository(ctx, issue.RepositoryId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("repository id (%d) doesn't exist", issue.RepositoryId))
	}
	var organization types.Organization

	if repository.Owner.Type == types.RepositoryOwner_ORGANIZATION {
		organization, found = k.GetOrganization(ctx, repository.Owner.Id)
		if !found {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("organization (%v) doesn't exist", repository.Owner.Id))
		}
	}

	if !utils.HaveIssuePermission(repository, msg.Creator, organization) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("user (%v) doesn't have permission to perform this operation", msg.Creator))
	}

	if len(issue.Assignees) < len(msg.Assignees) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "can't remove more than user assigned")
	}

	for _, a := range msg.Assignees {
		if i, exists := utils.AssigneeExists(issue.Assignees, a); exists {
			issue.Assignees = append(issue.Assignees[:i], issue.Assignees[i+1:]...)
		} else {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("assignee (%v) aren't assigned", a))
		}
	}

	issue.CommentsCount += 1
	issue.UpdatedAt = ctx.BlockTime().Unix()

	var comment = types.Comment{
		Creator:     "GITOPIA",
		ParentId:    msg.Id,
		CommentIid:  issue.CommentsCount,
		Body:        utils.RemoveAssigneesCommentBody(msg.Creator, msg.Assignees),
		System:      true,
		CreatedAt:   issue.UpdatedAt,
		UpdatedAt:   issue.UpdatedAt,
		CommentType: types.Comment_ISSUE,
	}

	id := k.AppendComment(
		ctx,
		comment,
	)

	issue.Comments = append(issue.Comments, id)

	k.SetIssue(ctx, issue)

	return &types.MsgRemoveIssueAssigneesResponse{}, nil
}

func (k msgServer) AddIssueLabels(goCtx context.Context, msg *types.MsgAddIssueLabels) (*types.MsgAddIssueLabelsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetUser(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("creator (%v) doesn't exist", msg.Creator))
	}

	issue, found := k.GetIssue(ctx, msg.IssueId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("issue id (%d) doesn't exist", msg.IssueId))
	}

	repository, found := k.GetRepository(ctx, issue.RepositoryId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("repository id (%d) doesn't exist", issue.RepositoryId))
	}
	var organization types.Organization

	if repository.Owner.Type == types.RepositoryOwner_ORGANIZATION {
		organization, found = k.GetOrganization(ctx, repository.Owner.Id)
		if !found {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("organization (%v) doesn't exist", repository.Owner.Id))
		}
	}

	if !utils.HaveIssuePermission(repository, msg.Creator, organization) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("user (%v) doesn't have permission to perform this operation", msg.Creator))
	}

	if len(issue.Labels)+len(msg.LabelIds) > 50 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "issue can't have more than 50 labels")
	}

	var labelNames []string

	for _, l := range msg.LabelIds {
		if i, exists := utils.RepositoryLabelIdExists(repository.Labels, l); exists {
			if _, exists := utils.LabelIdExists(issue.Labels, repository.Labels[i].Id); exists {
				return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("label id (%v) already exists in issue", l))
			}
			labelNames = append(labelNames, repository.Labels[i].Name)

			issue.Labels = append(issue.Labels, repository.Labels[i].Id)
		} else {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("label id (%v) doesn't exists in repository", l))
		}
	}

	issue.CommentsCount += 1
	issue.UpdatedAt = ctx.BlockTime().Unix()

	var comment = types.Comment{
		Creator:     "GITOPIA",
		ParentId:    msg.IssueId,
		CommentIid:  issue.CommentsCount,
		Body:        utils.AddLabelsCommentBody(msg.Creator, labelNames),
		System:      true,
		CreatedAt:   issue.UpdatedAt,
		UpdatedAt:   issue.UpdatedAt,
		CommentType: types.Comment_ISSUE,
	}

	id := k.AppendComment(
		ctx,
		comment,
	)

	issue.Comments = append(issue.Comments, id)

	k.SetIssue(ctx, issue)

	return &types.MsgAddIssueLabelsResponse{}, nil
}

func (k msgServer) RemoveIssueLabels(goCtx context.Context, msg *types.MsgRemoveIssueLabels) (*types.MsgRemoveIssueLabelsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetUser(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("creator (%v) doesn't exist", msg.Creator))
	}

	issue, found := k.GetIssue(ctx, msg.IssueId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("issue id (%d) doesn't exist", msg.IssueId))
	}

	repository, found := k.GetRepository(ctx, issue.RepositoryId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("repository id (%d) doesn't exist", issue.RepositoryId))
	}
	var organization types.Organization

	if repository.Owner.Type == types.RepositoryOwner_ORGANIZATION {
		organization, found = k.GetOrganization(ctx, repository.Owner.Id)
		if !found {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("organization (%v) doesn't exist", repository.Owner.Id))
		}
	}

	if !utils.HaveIssuePermission(repository, msg.Creator, organization) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("user (%v) doesn't have permission to perform this operation", msg.Creator))
	}

	if len(issue.Labels) < len(msg.LabelIds) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "can't remove more than existing labels")
	}

	var labelNames []string

	for _, l := range msg.LabelIds {
		if i, exists := utils.RepositoryLabelIdExists(repository.Labels, l); exists {
			if j, exists := utils.LabelIdExists(issue.Labels, repository.Labels[i].Id); exists {
				labelNames = append(labelNames, repository.Labels[i].Name)

				issue.Labels = append(issue.Labels[:j], issue.Labels[j+1:]...)
			} else {
				return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("label id (%v) doesn't exists in issue", l))
			}
		} else {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("label id (%v) doesn't exists in repository", l))
		}
	}

	issue.CommentsCount += 1
	issue.UpdatedAt = ctx.BlockTime().Unix()

	var comment = types.Comment{
		Creator:     "GITOPIA",
		ParentId:    msg.IssueId,
		CommentIid:  issue.CommentsCount,
		Body:        utils.RemoveLabelsCommentBody(msg.Creator, labelNames),
		System:      true,
		CreatedAt:   issue.UpdatedAt,
		UpdatedAt:   issue.UpdatedAt,
		CommentType: types.Comment_ISSUE,
	}

	id := k.AppendComment(
		ctx,
		comment,
	)

	issue.Comments = append(issue.Comments, id)

	k.SetIssue(ctx, issue)

	return &types.MsgRemoveIssueLabelsResponse{}, nil
}

func (k msgServer) DeleteIssue(goCtx context.Context, msg *types.MsgDeleteIssue) (*types.MsgDeleteIssueResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	issue, found := k.GetIssue(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("issue id (%d) doesn't exist", msg.Id))
	}

	if msg.Creator != issue.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveIssue(ctx, msg.Id)

	return &types.MsgDeleteIssueResponse{}, nil
}