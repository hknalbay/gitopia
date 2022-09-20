// Code generated by mockery v2.14.0. DO NOT EDIT.

package types

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockQueryClient is an autogenerated mock type for the QueryClient type
type MockQueryClient struct {
	mock.Mock
}

// AnyRepository provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) AnyRepository(ctx context.Context, in *QueryGetAnyRepositoryRequest, opts ...grpc.CallOption) (*QueryGetAnyRepositoryResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryGetAnyRepositoryResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryGetAnyRepositoryRequest, ...grpc.CallOption) *QueryGetAnyRepositoryResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryGetAnyRepositoryResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryGetAnyRepositoryRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AnyRepositoryAll provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) AnyRepositoryAll(ctx context.Context, in *QueryAllAnyRepositoryRequest, opts ...grpc.CallOption) (*QueryAllAnyRepositoryResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryAllAnyRepositoryResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryAllAnyRepositoryRequest, ...grpc.CallOption) *QueryAllAnyRepositoryResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryAllAnyRepositoryResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryAllAnyRepositoryRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BranchAll provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) BranchAll(ctx context.Context, in *QueryAllBranchRequest, opts ...grpc.CallOption) (*QueryAllBranchResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryAllBranchResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryAllBranchRequest, ...grpc.CallOption) *QueryAllBranchResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryAllBranchResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryAllBranchRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BranchSha provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) BranchSha(ctx context.Context, in *QueryGetRepositoryBranchShaRequest, opts ...grpc.CallOption) (*QueryGetRepositoryBranchShaResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryGetRepositoryBranchShaResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryGetRepositoryBranchShaRequest, ...grpc.CallOption) *QueryGetRepositoryBranchShaResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryGetRepositoryBranchShaResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryGetRepositoryBranchShaRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckGitServerAuthorization provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) CheckGitServerAuthorization(ctx context.Context, in *QueryCheckGitServerAuthorizationRequest, opts ...grpc.CallOption) (*QueryCheckGitServerAuthorizationResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryCheckGitServerAuthorizationResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryCheckGitServerAuthorizationRequest, ...grpc.CallOption) *QueryCheckGitServerAuthorizationResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryCheckGitServerAuthorizationResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryCheckGitServerAuthorizationRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Comment provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) Comment(ctx context.Context, in *QueryGetCommentRequest, opts ...grpc.CallOption) (*QueryGetCommentResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryGetCommentResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryGetCommentRequest, ...grpc.CallOption) *QueryGetCommentResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryGetCommentResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryGetCommentRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CommentAll provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) CommentAll(ctx context.Context, in *QueryAllCommentRequest, opts ...grpc.CallOption) (*QueryAllCommentResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryAllCommentResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryAllCommentRequest, ...grpc.CallOption) *QueryAllCommentResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryAllCommentResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryAllCommentRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ForkAll provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) ForkAll(ctx context.Context, in *QueryGetAllForkRequest, opts ...grpc.CallOption) (*QueryGetAllForkResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryGetAllForkResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryGetAllForkRequest, ...grpc.CallOption) *QueryGetAllForkResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryGetAllForkResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryGetAllForkRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Issue provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) Issue(ctx context.Context, in *QueryGetIssueRequest, opts ...grpc.CallOption) (*QueryGetIssueResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryGetIssueResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryGetIssueRequest, ...grpc.CallOption) *QueryGetIssueResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryGetIssueResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryGetIssueRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IssueAll provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) IssueAll(ctx context.Context, in *QueryAllIssueRequest, opts ...grpc.CallOption) (*QueryAllIssueResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryAllIssueResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryAllIssueRequest, ...grpc.CallOption) *QueryAllIssueResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryAllIssueResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryAllIssueRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Dao provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) Dao(ctx context.Context, in *QueryGetDaoRequest, opts ...grpc.CallOption) (*QueryGetDaoResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryGetDaoResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryGetDaoRequest, ...grpc.CallOption) *QueryGetDaoResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryGetDaoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryGetDaoRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DaoAll provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) DaoAll(ctx context.Context, in *QueryAllDaoRequest, opts ...grpc.CallOption) (*QueryAllDaoResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryAllDaoResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryAllDaoRequest, ...grpc.CallOption) *QueryAllDaoResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryAllDaoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryAllDaoRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PullRequest provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) PullRequest(ctx context.Context, in *QueryGetPullRequestRequest, opts ...grpc.CallOption) (*QueryGetPullRequestResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryGetPullRequestResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryGetPullRequestRequest, ...grpc.CallOption) *QueryGetPullRequestResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryGetPullRequestResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryGetPullRequestRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PullRequestAll provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) PullRequestAll(ctx context.Context, in *QueryAllPullRequestRequest, opts ...grpc.CallOption) (*QueryAllPullRequestResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryAllPullRequestResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryAllPullRequestRequest, ...grpc.CallOption) *QueryAllPullRequestResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryAllPullRequestResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryAllPullRequestRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PullRequestMergePermission provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) PullRequestMergePermission(ctx context.Context, in *QueryGetPullRequestMergePermissionRequest, opts ...grpc.CallOption) (*QueryGetPullRequestMergePermissionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryGetPullRequestMergePermissionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryGetPullRequestMergePermissionRequest, ...grpc.CallOption) *QueryGetPullRequestMergePermissionResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryGetPullRequestMergePermissionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryGetPullRequestMergePermissionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Release provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) Release(ctx context.Context, in *QueryGetReleaseRequest, opts ...grpc.CallOption) (*QueryGetReleaseResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryGetReleaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryGetReleaseRequest, ...grpc.CallOption) *QueryGetReleaseResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryGetReleaseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryGetReleaseRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReleaseAll provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) ReleaseAll(ctx context.Context, in *QueryAllReleaseRequest, opts ...grpc.CallOption) (*QueryAllReleaseResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryAllReleaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryAllReleaseRequest, ...grpc.CallOption) *QueryAllReleaseResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryAllReleaseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryAllReleaseRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) Repository(ctx context.Context, in *QueryGetRepositoryRequest, opts ...grpc.CallOption) (*QueryGetRepositoryResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryGetRepositoryResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryGetRepositoryRequest, ...grpc.CallOption) *QueryGetRepositoryResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryGetRepositoryResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryGetRepositoryRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RepositoryAll provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) RepositoryAll(ctx context.Context, in *QueryAllRepositoryRequest, opts ...grpc.CallOption) (*QueryAllRepositoryResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryAllRepositoryResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryAllRepositoryRequest, ...grpc.CallOption) *QueryAllRepositoryResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryAllRepositoryResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryAllRepositoryRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RepositoryIssue provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) RepositoryIssue(ctx context.Context, in *QueryGetRepositoryIssueRequest, opts ...grpc.CallOption) (*QueryGetRepositoryIssueResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryGetRepositoryIssueResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryGetRepositoryIssueRequest, ...grpc.CallOption) *QueryGetRepositoryIssueResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryGetRepositoryIssueResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryGetRepositoryIssueRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RepositoryIssueAll provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) RepositoryIssueAll(ctx context.Context, in *QueryAllRepositoryIssueRequest, opts ...grpc.CallOption) (*QueryAllRepositoryIssueResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryAllRepositoryIssueResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryAllRepositoryIssueRequest, ...grpc.CallOption) *QueryAllRepositoryIssueResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryAllRepositoryIssueResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryAllRepositoryIssueRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RepositoryPullRequest provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) RepositoryPullRequest(ctx context.Context, in *QueryGetRepositoryPullRequestRequest, opts ...grpc.CallOption) (*QueryGetRepositoryPullRequestResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryGetRepositoryPullRequestResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryGetRepositoryPullRequestRequest, ...grpc.CallOption) *QueryGetRepositoryPullRequestResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryGetRepositoryPullRequestResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryGetRepositoryPullRequestRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RepositoryPullRequestAll provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) RepositoryPullRequestAll(ctx context.Context, in *QueryAllRepositoryPullRequestRequest, opts ...grpc.CallOption) (*QueryAllRepositoryPullRequestResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryAllRepositoryPullRequestResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryAllRepositoryPullRequestRequest, ...grpc.CallOption) *QueryAllRepositoryPullRequestResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryAllRepositoryPullRequestResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryAllRepositoryPullRequestRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RepositoryRelease provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) RepositoryRelease(ctx context.Context, in *QueryGetRepositoryReleaseRequest, opts ...grpc.CallOption) (*QueryGetRepositoryReleaseResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryGetRepositoryReleaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryGetRepositoryReleaseRequest, ...grpc.CallOption) *QueryGetRepositoryReleaseResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryGetRepositoryReleaseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryGetRepositoryReleaseRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RepositoryReleaseAll provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) RepositoryReleaseAll(ctx context.Context, in *QueryAllRepositoryReleaseRequest, opts ...grpc.CallOption) (*QueryAllRepositoryReleaseResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryAllRepositoryReleaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryAllRepositoryReleaseRequest, ...grpc.CallOption) *QueryAllRepositoryReleaseResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryAllRepositoryReleaseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryAllRepositoryReleaseRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RepositoryReleaseLatest provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) RepositoryReleaseLatest(ctx context.Context, in *QueryGetLatestRepositoryReleaseRequest, opts ...grpc.CallOption) (*QueryGetLatestRepositoryReleaseResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryGetLatestRepositoryReleaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryGetLatestRepositoryReleaseRequest, ...grpc.CallOption) *QueryGetLatestRepositoryReleaseResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryGetLatestRepositoryReleaseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryGetLatestRepositoryReleaseRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagAll provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) TagAll(ctx context.Context, in *QueryAllTagRequest, opts ...grpc.CallOption) (*QueryAllTagResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryAllTagResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryAllTagRequest, ...grpc.CallOption) *QueryAllTagResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryAllTagResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryAllTagRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagSha provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) TagSha(ctx context.Context, in *QueryGetRepositoryTagShaRequest, opts ...grpc.CallOption) (*QueryGetRepositoryTagShaResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryGetRepositoryTagShaResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryGetRepositoryTagShaRequest, ...grpc.CallOption) *QueryGetRepositoryTagShaResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryGetRepositoryTagShaResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryGetRepositoryTagShaRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Task provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) Task(ctx context.Context, in *QueryGetTaskRequest, opts ...grpc.CallOption) (*QueryGetTaskResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryGetTaskResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryGetTaskRequest, ...grpc.CallOption) *QueryGetTaskResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryGetTaskResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryGetTaskRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TaskAll provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) TaskAll(ctx context.Context, in *QueryAllTaskRequest, opts ...grpc.CallOption) (*QueryAllTaskResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryAllTaskResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryAllTaskRequest, ...grpc.CallOption) *QueryAllTaskResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryAllTaskResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryAllTaskRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// User provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) User(ctx context.Context, in *QueryGetUserRequest, opts ...grpc.CallOption) (*QueryGetUserResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryGetUserResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryGetUserRequest, ...grpc.CallOption) *QueryGetUserResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryGetUserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryGetUserRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserAll provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) UserAll(ctx context.Context, in *QueryAllUserRequest, opts ...grpc.CallOption) (*QueryAllUserResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryAllUserResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryAllUserRequest, ...grpc.CallOption) *QueryAllUserResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryAllUserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryAllUserRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserDaoAll provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) UserDaoAll(ctx context.Context, in *QueryAllUserDaoRequest, opts ...grpc.CallOption) (*QueryAllUserDaoResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryAllUserDaoResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryAllUserDaoRequest, ...grpc.CallOption) *QueryAllUserDaoResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryAllUserDaoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryAllUserDaoRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Whois provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) Whois(ctx context.Context, in *QueryGetWhoisRequest, opts ...grpc.CallOption) (*QueryGetWhoisResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryGetWhoisResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryGetWhoisRequest, ...grpc.CallOption) *QueryGetWhoisResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryGetWhoisResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryGetWhoisRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WhoisAll provides a mock function with given fields: ctx, in, opts
func (_m *MockQueryClient) WhoisAll(ctx context.Context, in *QueryAllWhoisRequest, opts ...grpc.CallOption) (*QueryAllWhoisResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *QueryAllWhoisResponse
	if rf, ok := ret.Get(0).(func(context.Context, *QueryAllWhoisRequest, ...grpc.CallOption) *QueryAllWhoisResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryAllWhoisResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *QueryAllWhoisRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockQueryClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockQueryClient creates a new instance of MockQueryClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockQueryClient(t mockConstructorTestingTNewMockQueryClient) *MockQueryClient {
	mock := &MockQueryClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
