package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/shoplineos/shopline-sdk-go/client"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/test"
	"github.com/shoplineos/shopline-sdk-go/rest/admin/v20260301/onlinestore"
	"github.com/stretchr/testify/assert"
)

// ── Helpers ──────────────────────────────────────────────────────────────────

func onlineStoreURL(cli *client.Client, path string) string {
	return fmt.Sprintf("https://%s.myshopline.com/%s/%s/%s",
		cli.StoreHandle, cli.PathPrefix, cli.ApiVersion, path)
}

// ══════════════════════════════════════════════════════════════════════════════
// Comment APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestApproveTheComment(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	commentId := "638040c6a0a9ba4487960880"
	mockResp := `{"comment":{"id":"638040c6a0a9ba4487960880","status":"approved","author":"demo-author","email":"demo@qq.com","body":"demo-content","body_html":"<p>demo-content</p>","blog_id":"63a2afa527797d5cedcecb52","blog_collection_id":"639bfda0ee877c4a9bd600b6","create_at":"2022-08-03T10:45:00+08:00","published_at":"2022-08-03T10:45:00+08:00","updated_at":"2022-08-03T10:45:00+08:00"}}`

	httpmock.RegisterResponder("POST",
		onlineStoreURL(cli, fmt.Sprintf("store/comments/%s/approve.json", commentId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.ApproveTheCommentAPIReq{CommentId: commentId}
	apiResp := &onlinestore.ApproveTheCommentAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, commentId, apiResp.Comment.Id)
	assert.Equal(t, "approved", apiResp.Comment.Status)
	assert.Equal(t, "demo-author", apiResp.Comment.Author)
	assert.Equal(t, "demo@qq.com", apiResp.Comment.Email)
}

func TestApproveTheComment_MissingCommentId(t *testing.T) {
	err := (&onlinestore.ApproveTheCommentAPIReq{}).Verify()
	assert.EqualError(t, err, "CommentId is required")
}

func TestByReviewingIdQueryComments(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	commentId := "63aaa20f6b7b014b289c0765"
	mockResp := `{"comment":{"id":"63aaa20f6b7b014b289c0765","status":"unapproved","author":"demo-author","email":"demo@qq.com"}}`

	httpmock.RegisterResponder("GET",
		onlineStoreURL(cli, fmt.Sprintf("store/comments/%s.json", commentId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.ByReviewingIdQueryCommentsAPIReq{CommentId: commentId}
	apiResp := &onlinestore.ByReviewingIdQueryCommentsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, commentId, apiResp.Comment.Id)
	assert.Equal(t, "unapproved", apiResp.Comment.Status)
}

func TestByReviewingIdQueryComments_MissingCommentId(t *testing.T) {
	err := (&onlinestore.ByReviewingIdQueryCommentsAPIReq{}).Verify()
	assert.EqualError(t, err, "CommentId is required")
}

func TestCommentForRecoveryDelete(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	commentId := "638040c6a0a9ba4487960880"
	mockResp := `{"comment":{"id":"638040c6a0a9ba4487960880","status":"unapproved"}}`

	httpmock.RegisterResponder("POST",
		onlineStoreURL(cli, fmt.Sprintf("store/comments/%s/restore.json", commentId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.CommentForRecoveryDeleteAPIReq{CommentId: commentId}
	apiResp := &onlinestore.CommentForRecoveryDeleteAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, commentId, apiResp.Comment.Id)
}

func TestCommentForRecoveryDelete_MissingCommentId(t *testing.T) {
	err := (&onlinestore.CommentForRecoveryDeleteAPIReq{}).Verify()
	assert.EqualError(t, err, "CommentId is required")
}

func TestCreateAComment(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"comment":{"id":"63aaa20f6b7b014b289c0765","status":"unapproved","author":"demo-author","email":"demo@qq.com","body":"demo-content","blog_id":"63a2afa527797d5cedcecb52"}}`

	httpmock.RegisterResponder("POST",
		onlineStoreURL(cli, "store/comments.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.CreateACommentAPIReq{
		Author: "demo-author",
		BlogId: "63a2afa527797d5cedcecb52",
		Body:   "demo-content",
		Email:  "demo@qq.com",
		Ip:     "172.28.173.136",
	}
	apiResp := &onlinestore.CreateACommentAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "63aaa20f6b7b014b289c0765", apiResp.Comment.Id)
	assert.Equal(t, "demo-author", apiResp.Comment.Author)
}

func TestCreateAComment_MissingRequired(t *testing.T) {
	assert.EqualError(t, (&onlinestore.CreateACommentAPIReq{}).Verify(), "Author is required")
	assert.EqualError(t, (&onlinestore.CreateACommentAPIReq{Author: "a"}).Verify(), "BlogId is required")
	assert.EqualError(t, (&onlinestore.CreateACommentAPIReq{Author: "a", BlogId: "b"}).Verify(), "Body is required")
	assert.EqualError(t, (&onlinestore.CreateACommentAPIReq{Author: "a", BlogId: "b", Body: "c"}).Verify(), "Email is required")
	assert.EqualError(t, (&onlinestore.CreateACommentAPIReq{Author: "a", BlogId: "b", Body: "c", Email: "e"}).Verify(), "Ip is required")
}

func TestDeleteAComment(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	commentId := "638040c6a0a9ba4487960880"
	mockResp := `{"comment":{"id":"638040c6a0a9ba4487960880","status":"removed"}}`

	httpmock.RegisterResponder("POST",
		onlineStoreURL(cli, fmt.Sprintf("store/comments/%s/remove.json", commentId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.DeleteACommentAPIReq{CommentId: commentId}
	apiResp := &onlinestore.DeleteACommentAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, commentId, apiResp.Comment.Id)
}

func TestDeleteAComment_MissingCommentId(t *testing.T) {
	err := (&onlinestore.DeleteACommentAPIReq{}).Verify()
	assert.EqualError(t, err, "CommentId is required")
}

func TestLabeledCommentsAsSpam(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	commentId := "638040c6a0a9ba4487960880"
	mockResp := `{"comment":{"id":"638040c6a0a9ba4487960880","status":"spam"}}`

	httpmock.RegisterResponder("POST",
		onlineStoreURL(cli, fmt.Sprintf("store/comments/%s/spam.json", commentId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.LabeledCommentsAsSpamAPIReq{CommentId: commentId}
	apiResp := &onlinestore.LabeledCommentsAsSpamAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, commentId, apiResp.Comment.Id)
	assert.Equal(t, "spam", apiResp.Comment.Status)
}

func TestLabeledCommentsAsSpam_MissingCommentId(t *testing.T) {
	err := (&onlinestore.LabeledCommentsAsSpamAPIReq{}).Verify()
	assert.EqualError(t, err, "CommentId is required")
}

func TestLabeledCommentsAsNonSpam(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	commentId := "638040c6a0a9ba4487960880"
	mockResp := `{"comment":{"id":"638040c6a0a9ba4487960880","status":"published"}}`

	httpmock.RegisterResponder("POST",
		onlineStoreURL(cli, fmt.Sprintf("store/comments/%s/not_spam.json", commentId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.LabeledCommentsAsNonSpamAPIReq{CommentId: commentId}
	apiResp := &onlinestore.LabeledCommentsAsNonSpamAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, commentId, apiResp.Comment.Id)
}

func TestLabeledCommentsAsNonSpam_MissingCommentId(t *testing.T) {
	err := (&onlinestore.LabeledCommentsAsNonSpamAPIReq{}).Verify()
	assert.EqualError(t, err, "CommentId is required")
}

func TestNumberOfGetTheComments(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"count":42}`

	httpmock.RegisterResponder("GET",
		onlineStoreURL(cli, "store/comments/count.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.NumberOfGetTheCommentsAPIReq{}
	apiResp := &onlinestore.NumberOfGetTheCommentsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, int64(42), apiResp.Count)
}

func TestQueryListOfComments(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"comments":[{"id":"63aaa20f6b7b014b289c0765","author":"demo-author","status":"approved"}]}`

	httpmock.RegisterResponder("GET",
		onlineStoreURL(cli, "store/comments.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.QueryListOfCommentsAPIReq{}
	apiResp := &onlinestore.QueryListOfCommentsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Comments, 1)
	assert.Equal(t, "63aaa20f6b7b014b289c0765", apiResp.Comments[0].Id)
}

func TestUpdateComments(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	commentId := "63aaa20f6b7b014b289c0765"
	mockResp := `{"comment":{"id":"63aaa20f6b7b014b289c0765","body":"updated-content","author":"demo-author"}}`

	httpmock.RegisterResponder("PUT",
		onlineStoreURL(cli, fmt.Sprintf("store/comments/%s.json", commentId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.UpdateCommentsAPIReq{CommentId: commentId, Body: "updated-content"}
	apiResp := &onlinestore.UpdateCommentsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, commentId, apiResp.Comment.Id)
	assert.Equal(t, "updated-content", apiResp.Comment.Body)
}

func TestUpdateComments_MissingCommentId(t *testing.T) {
	err := (&onlinestore.UpdateCommentsAPIReq{}).Verify()
	assert.EqualError(t, err, "CommentId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Blog Collection APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestCreateABlogCollection(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"blog":{"id":"63a9a344609552261f00f23e","title":"demo blogs","handle":"demo-blogs","commentable":"moderate"}}`

	httpmock.RegisterResponder("POST",
		onlineStoreURL(cli, "store/blogs.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.CreateABlogCollectionAPIReq{
		Blog: onlinestore.CreateABlogCollectionAPIReqBlog{
			Title:       "demo blogs",
			Handle:      "demo-blogs",
			Commentable: "moderate",
		},
	}
	apiResp := &onlinestore.CreateABlogCollectionAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "63a9a344609552261f00f23e", apiResp.Blog.Id)
	assert.Equal(t, "demo blogs", apiResp.Blog.Title)
}

func TestQueryBlogCollectionList(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"blogs":[{"id":"63a9a344609552261f00f23e","title":"demo blogs","handle":"demo-blogs"}]}`

	httpmock.RegisterResponder("GET",
		onlineStoreURL(cli, "store/blogs.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.QueryBlogCollectionListAPIReq{}
	apiResp := &onlinestore.QueryBlogCollectionListAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Blogs, 1)
	assert.Equal(t, "63a9a344609552261f00f23e", apiResp.Blogs[0].Id)
}

func TestQueryBlogCollectionQuantity(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"count":5}`

	httpmock.RegisterResponder("GET",
		onlineStoreURL(cli, "store/blogs/count.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.QueryBlogCollectionQuantityAPIReq{}
	apiResp := &onlinestore.QueryBlogCollectionQuantityAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, int64(5), apiResp.Count)
}

func TestQuerySingleBlogCollection(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	blogCollectionId := "63a9a344609552261f00f23e"
	mockResp := `{"blog":{"id":"63a9a344609552261f00f23e","title":"demo blogs","handle":"demo-blogs","commentable":"moderate"}}`

	httpmock.RegisterResponder("GET",
		onlineStoreURL(cli, fmt.Sprintf("store/blogs/%s.json", blogCollectionId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.QuerySingleBlogCollectionAPIReq{BlogCollectionId: blogCollectionId}
	apiResp := &onlinestore.QuerySingleBlogCollectionAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, blogCollectionId, apiResp.Blog.Id)
	assert.Equal(t, "demo blogs", apiResp.Blog.Title)
}

func TestQuerySingleBlogCollection_MissingBlogCollectionId(t *testing.T) {
	err := (&onlinestore.QuerySingleBlogCollectionAPIReq{}).Verify()
	assert.EqualError(t, err, "BlogCollectionId is required")
}

func TestUpdateBlogCollection(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	blogCollectionId := "63a9a344609552261f00f23e"
	mockResp := `{"blog":{"id":"63a9a344609552261f00f23e","title":"updated blogs","handle":"demo-blogs"}}`

	httpmock.RegisterResponder("PUT",
		onlineStoreURL(cli, fmt.Sprintf("store/blogs/%s.json", blogCollectionId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.UpdateBlogCollectionAPIReq{
		BlogCollectionId: blogCollectionId,
		Blog:             onlinestore.UpdateBlogCollectionAPIReqBlog{Title: "updated blogs"},
	}
	apiResp := &onlinestore.UpdateBlogCollectionAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, blogCollectionId, apiResp.Blog.Id)
	assert.Equal(t, "updated blogs", apiResp.Blog.Title)
}

func TestUpdateBlogCollection_MissingBlogCollectionId(t *testing.T) {
	err := (&onlinestore.UpdateBlogCollectionAPIReq{}).Verify()
	assert.EqualError(t, err, "BlogCollectionId is required")
}

func TestDeleteBlogCollection(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	blogCollectionId := "63a9a344609552261f00f23e"

	httpmock.RegisterResponder("DELETE",
		onlineStoreURL(cli, fmt.Sprintf("store/blogs/%s.json", blogCollectionId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &onlinestore.DeleteBlogCollectionAPIReq{BlogCollectionId: blogCollectionId}
	apiResp := &onlinestore.DeleteBlogCollectionAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteBlogCollection_MissingBlogCollectionId(t *testing.T) {
	err := (&onlinestore.DeleteBlogCollectionAPIReq{}).Verify()
	assert.EqualError(t, err, "BlogCollectionId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Blog Post APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestCreateBlogArticlesForTheCollection(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	blogCollectionId := "64e313c4cd5956279e61d150"
	mockResp := `{"blog":{"id":"66718d010588d64ef7d15c96","title":"demo title","author":"Alvin","blog_collection_id":"64e313c4cd5956279e61d150"}}`

	httpmock.RegisterResponder("POST",
		onlineStoreURL(cli, fmt.Sprintf("store/blogs/%s/articles.json", blogCollectionId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.CreateBlogArticlesForTheCollectionAPIReq{
		BlogCollectionId: blogCollectionId,
		Blog:             onlinestore.CreateBlogArticlesForTheCollectionAPIReqBlog{Title: "demo title", Author: "Alvin"},
	}
	apiResp := &onlinestore.CreateBlogArticlesForTheCollectionAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "66718d010588d64ef7d15c96", apiResp.Blog.Id)
	assert.Equal(t, blogCollectionId, apiResp.Blog.BlogCollectionId)
}

func TestCreateBlogArticlesForTheCollection_MissingBlogCollectionId(t *testing.T) {
	err := (&onlinestore.CreateBlogArticlesForTheCollectionAPIReq{}).Verify()
	assert.EqualError(t, err, "BlogCollectionId is required")
}

func TestQueryBlogPost(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	blogCollectionId := "64e313c4cd5956279e61d150"
	mockResp := `{"count":10}`

	httpmock.RegisterResponder("GET",
		onlineStoreURL(cli, fmt.Sprintf("store/blogs/%s/articles/count.json", blogCollectionId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.QueryBlogPostAPIReq{BlogCollectionId: blogCollectionId}
	apiResp := &onlinestore.QueryBlogPostAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, int64(10), apiResp.Count)
}

func TestQueryBlogPost_MissingBlogCollectionId(t *testing.T) {
	err := (&onlinestore.QueryBlogPostAPIReq{}).Verify()
	assert.EqualError(t, err, "BlogCollectionId is required")
}

func TestQueryBlogPostDetails(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	blogCollectionId := "64e313c4cd5956279e61d150"
	blogId := "66718d010588d64ef7d15c96"
	mockResp := `{"blog":{"id":"66718d010588d64ef7d15c96","title":"demo title","author":"Alvin","blog_collection_id":"64e313c4cd5956279e61d150"}}`

	httpmock.RegisterResponder("GET",
		onlineStoreURL(cli, fmt.Sprintf("store/blogs/%s/articles/%s.json", blogCollectionId, blogId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.QueryBlogPostDetailsAPIReq{BlogCollectionId: blogCollectionId, BlogId: blogId}
	apiResp := &onlinestore.QueryBlogPostDetailsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, blogId, apiResp.Blog.Id)
	assert.Equal(t, "Alvin", apiResp.Blog.Author)
}

func TestQueryBlogPostDetails_MissingRequired(t *testing.T) {
	assert.EqualError(t, (&onlinestore.QueryBlogPostDetailsAPIReq{}).Verify(), "BlogCollectionId is required")
	assert.EqualError(t, (&onlinestore.QueryBlogPostDetailsAPIReq{BlogCollectionId: "x"}).Verify(), "BlogId is required")
}

func TestQueryBlogPostList(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	blogCollectionId := "64e313c4cd5956279e61d150"
	mockResp := `{"blogs":[{"id":"66718d010588d64ef7d15c96","title":"demo title","author":"Alvin"}]}`

	httpmock.RegisterResponder("GET",
		onlineStoreURL(cli, fmt.Sprintf("store/blogs/%s/articles.json", blogCollectionId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.QueryBlogPostListAPIReq{BlogCollectionId: blogCollectionId}
	apiResp := &onlinestore.QueryBlogPostListAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Blogs, 1)
	assert.Equal(t, "66718d010588d64ef7d15c96", apiResp.Blogs[0].Id)
}

func TestQueryBlogPostList_MissingBlogCollectionId(t *testing.T) {
	err := (&onlinestore.QueryBlogPostListAPIReq{}).Verify()
	assert.EqualError(t, err, "BlogCollectionId is required")
}

func TestQueryListOfAllArticlesAuthors(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"authors":["Alvin","Bob"]}`

	httpmock.RegisterResponder("GET",
		onlineStoreURL(cli, "store/blogs/authors.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.QueryListOfAllArticlesAuthorsAPIReq{}
	apiResp := &onlinestore.QueryListOfAllArticlesAuthorsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, []string{"Alvin", "Bob"}, apiResp.Authors)
}

func TestUpdateBlogPost(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	blogCollectionId := "64e313c4cd5956279e61d150"
	blogId := "66718d010588d64ef7d15c96"
	mockResp := `{"blog":{"id":"66718d010588d64ef7d15c96","title":"updated title","author":"Alvin"}}`

	httpmock.RegisterResponder("PUT",
		onlineStoreURL(cli, fmt.Sprintf("store/blogs/%s/articles/%s.json", blogCollectionId, blogId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.UpdateBlogPostAPIReq{
		BlogCollectionId: blogCollectionId,
		BlogId:           blogId,
		Blog:             onlinestore.UpdateBlogPostAPIReqBlog{Title: "updated title"},
	}
	apiResp := &onlinestore.UpdateBlogPostAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, blogId, apiResp.Blog.Id)
	assert.Equal(t, "updated title", apiResp.Blog.Title)
}

func TestUpdateBlogPost_MissingRequired(t *testing.T) {
	assert.EqualError(t, (&onlinestore.UpdateBlogPostAPIReq{}).Verify(), "BlogCollectionId is required")
	assert.EqualError(t, (&onlinestore.UpdateBlogPostAPIReq{BlogCollectionId: "x"}).Verify(), "BlogId is required")
}

func TestDeleteBlogPost(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	blogCollectionId := "64e313c4cd5956279e61d150"
	blogId := "66718d010588d64ef7d15c96"

	httpmock.RegisterResponder("DELETE",
		onlineStoreURL(cli, fmt.Sprintf("store/blogs/%s/articles/%s.json", blogCollectionId, blogId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &onlinestore.DeleteBlogPostAPIReq{BlogCollectionId: blogCollectionId, BlogId: blogId}
	apiResp := &onlinestore.DeleteBlogPostAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteBlogPost_MissingRequired(t *testing.T) {
	assert.EqualError(t, (&onlinestore.DeleteBlogPostAPIReq{}).Verify(), "BlogCollectionId is required")
	assert.EqualError(t, (&onlinestore.DeleteBlogPostAPIReq{BlogCollectionId: "x"}).Verify(), "BlogId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Script Tag APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestCreateAScriptTag(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"script_tag":{"id":"61a75c4e9f8c201e6e9473e0","event":"onload","src":"https://djavaskripped.org/fancy.js","display_scope":"all"}}`

	httpmock.RegisterResponder("POST",
		onlineStoreURL(cli, "store/script_tags.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.CreateAScriptTagAPIReq{
		ScriptTag: onlinestore.CreateAScriptTagAPIReqScriptTag{
			Event:        "onload",
			Src:          "https://djavaskripped.org/fancy.js",
			DisplayScope: "all",
		},
	}
	apiResp := &onlinestore.CreateAScriptTagAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "61a75c4e9f8c201e6e9473e0", apiResp.ScriptTag.Id)
	assert.Equal(t, "onload", apiResp.ScriptTag.Event)
}

func TestQueryASingleScriptTag(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	scriptTagId := "61a75c4e9f8c201e6e9473e0"
	mockResp := `{"script_tag":{"id":"61a75c4e9f8c201e6e9473e0","event":"onload","src":"https://djavaskripped.org/fancy.js"}}`

	httpmock.RegisterResponder("GET",
		onlineStoreURL(cli, fmt.Sprintf("store/%s.json", scriptTagId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.QueryASingleScriptTagAPIReq{ScriptTagId: scriptTagId}
	apiResp := &onlinestore.QueryASingleScriptTagAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, scriptTagId, apiResp.ScriptTag.Id)
}

func TestQueryASingleScriptTag_MissingScriptTagId(t *testing.T) {
	err := (&onlinestore.QueryASingleScriptTagAPIReq{}).Verify()
	assert.EqualError(t, err, "ScriptTagId is required")
}

func TestQueryScriptTagList(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"script_tags":[{"id":"61a75c4e9f8c201e6e9473e0","event":"onload","src":"https://djavaskripped.org/fancy.js"}]}`

	httpmock.RegisterResponder("GET",
		onlineStoreURL(cli, "store/script_tags.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.QueryScriptTagListAPIReq{}
	apiResp := &onlinestore.QueryScriptTagListAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.ScriptTags, 1)
	assert.Equal(t, "61a75c4e9f8c201e6e9473e0", apiResp.ScriptTags[0].Id)
}

func TestQueryTheNumberOfScriptTags(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"count":3}`

	httpmock.RegisterResponder("GET",
		onlineStoreURL(cli, "store/script_tags/count.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.QueryTheNumberOfScriptTagsAPIReq{}
	apiResp := &onlinestore.QueryTheNumberOfScriptTagsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, int64(3), apiResp.Count)
}

func TestUpdateScriptTag(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	scriptTagId := "61a75c4e9f8c201e6e9473e0"
	mockResp := `{"script_tag":{"id":"61a75c4e9f8c201e6e9473e0","src":"https://updated.org/script.js","event":"onload"}}`

	httpmock.RegisterResponder("PUT",
		onlineStoreURL(cli, fmt.Sprintf("store/script_tags/%s.json", scriptTagId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.UpdateScriptTagAPIReq{
		ScriptTagId: scriptTagId,
		ScriptTag:   onlinestore.UpdateScriptTagAPIReqScriptTag{Src: "https://updated.org/script.js"},
	}
	apiResp := &onlinestore.UpdateScriptTagAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, scriptTagId, apiResp.ScriptTag.Id)
}

func TestUpdateScriptTag_MissingScriptTagId(t *testing.T) {
	err := (&onlinestore.UpdateScriptTagAPIReq{}).Verify()
	assert.EqualError(t, err, "ScriptTagId is required")
}

func TestDeleteScriptTag(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	scriptTagId := "61a75c4e9f8c201e6e9473e0"

	httpmock.RegisterResponder("DELETE",
		onlineStoreURL(cli, fmt.Sprintf("store/script_tags/%s.json", scriptTagId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &onlinestore.DeleteScriptTagAPIReq{ScriptTagId: scriptTagId}
	apiResp := &onlinestore.DeleteScriptTagAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteScriptTag_MissingScriptTagId(t *testing.T) {
	err := (&onlinestore.DeleteScriptTagAPIReq{}).Verify()
	assert.EqualError(t, err, "ScriptTagId is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Theme APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestCreateTheme(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"theme":{"id":"627d1003d4baa549bf5b83f9","name":"demo_theme","role":0}}`

	httpmock.RegisterResponder("POST",
		onlineStoreURL(cli, "themes.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.CreateThemeAPIReq{Name: "demo_theme"}
	apiResp := &onlinestore.CreateThemeAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "627d1003d4baa549bf5b83f9", apiResp.Theme.Id)
	assert.Equal(t, "demo_theme", apiResp.Theme.Name)
}

func TestCreateTheme_MissingName(t *testing.T) {
	err := (&onlinestore.CreateThemeAPIReq{}).Verify()
	assert.EqualError(t, err, "Name is required")
}

func TestGetAllThemes(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"themes":[{"id":"671f4c96d2682a1f5407536e","name":"Modern1","role":"published"}]}`

	httpmock.RegisterResponder("GET",
		onlineStoreURL(cli, "themes.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.GetAllThemesAPIReq{}
	apiResp := &onlinestore.GetAllThemesAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Themes, 1)
	assert.Equal(t, "671f4c96d2682a1f5407536e", apiResp.Themes[0].Id)
	assert.Equal(t, "Modern1", apiResp.Themes[0].Name)
}

func TestGetASingleTheme(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	themeId := "671f4c96d2682a1f5407536e"
	mockResp := `{"theme":{"id":"671f4c96d2682a1f5407536e","name":"Modern1","role":"published"}}`

	httpmock.RegisterResponder("GET",
		onlineStoreURL(cli, fmt.Sprintf("themes/%s.json", themeId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.GetASingleThemeAPIReq{ThemeId: themeId}
	apiResp := &onlinestore.GetASingleThemeAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, themeId, apiResp.Theme.Id)
	assert.Equal(t, "Modern1", apiResp.Theme.Name)
}

func TestGetASingleTheme_MissingThemeId(t *testing.T) {
	err := (&onlinestore.GetASingleThemeAPIReq{}).Verify()
	assert.EqualError(t, err, "ThemeId is required")
}

func TestUpdateTheme(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	themeId := "627d1003d4baa549bf5b83f9"
	mockResp := `{"theme":{"id":"627d1003d4baa549bf5b83f9","name":"updated_theme","role":"published"}}`

	httpmock.RegisterResponder("PUT",
		onlineStoreURL(cli, fmt.Sprintf("themes/%s.json", themeId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.UpdateThemeAPIReq{ThemeId: themeId, Name: "updated_theme"}
	apiResp := &onlinestore.UpdateThemeAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, themeId, apiResp.Theme.Id)
	assert.Equal(t, "updated_theme", apiResp.Theme.Name)
}

func TestUpdateTheme_MissingThemeId(t *testing.T) {
	err := (&onlinestore.UpdateThemeAPIReq{}).Verify()
	assert.EqualError(t, err, "ThemeId is required")
}

func TestDeleteTheme(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	themeId := "627d1003d4baa549bf5b83f9"
	mockResp := `{"theme":{"id":"627d1003d4baa549bf5b83f9"}}`

	httpmock.RegisterResponder("DELETE",
		onlineStoreURL(cli, fmt.Sprintf("themes/%s.json", themeId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.DeleteThemeAPIReq{ThemeId: themeId}
	apiResp := &onlinestore.DeleteThemeAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteTheme_MissingThemeId(t *testing.T) {
	err := (&onlinestore.DeleteThemeAPIReq{}).Verify()
	assert.EqualError(t, err, "ThemeId is required")
}

func TestGetAllFilesForATheme(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	themeId := "671f4c96d2682a1f5407536e"
	mockResp := `{"assets":[{"key":"config/settings_data.json","theme_id":"671f4c96d2682a1f5407536e","content_type":"application/json","size":1024}]}`

	httpmock.RegisterResponder("GET",
		onlineStoreURL(cli, fmt.Sprintf("themes/%s/assets/list.json", themeId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.GetAllFilesForAThemeAPIReq{ThemeId: themeId}
	apiResp := &onlinestore.GetAllFilesForAThemeAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Assets, 1)
	assert.Equal(t, "config/settings_data.json", apiResp.Assets[0].Key)
	assert.Equal(t, themeId, apiResp.Assets[0].ThemeId)
}

func TestGetAllFilesForATheme_MissingThemeId(t *testing.T) {
	err := (&onlinestore.GetAllFilesForAThemeAPIReq{}).Verify()
	assert.EqualError(t, err, "ThemeId is required")
}

func TestBatchCreateModifyAssetContent(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	themeId := "66012436ec4cd35a973324c4"
	mockResp := `{"assets":[{"key":"assets/blog.css","theme_id":"66012436ec4cd35a973324c4","content_type":"text/css","size":500}]}`

	httpmock.RegisterResponder("POST",
		onlineStoreURL(cli, fmt.Sprintf("themes/%s/batch-assets.json", themeId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.BatchCreateModifyAssetContentAPIReq{
		ThemeId: themeId,
		Assets:  []onlinestore.BatchCreateModifyAssetContentAPIReqAsset{{Key: "assets/blog.css", Value: "body{}"}},
	}
	apiResp := &onlinestore.BatchCreateModifyAssetContentAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Assets, 1)
	assert.Equal(t, "assets/blog.css", apiResp.Assets[0].Key)
}

func TestBatchCreateModifyAssetContent_MissingThemeId(t *testing.T) {
	err := (&onlinestore.BatchCreateModifyAssetContentAPIReq{}).Verify()
	assert.EqualError(t, err, "ThemeId is required")
}

func TestDeleteThemeAsset(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	themeId := "671f4c96d2682a1f5407536e"

	httpmock.RegisterResponder("DELETE",
		onlineStoreURL(cli, fmt.Sprintf("themes/%s/assets.json", themeId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &onlinestore.DeleteThemeAssetAPIReq{ThemeId: themeId, AssetKey: "assets/blog.css"}
	apiResp := &onlinestore.DeleteThemeAssetAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestDeleteThemeAsset_MissingRequired(t *testing.T) {
	assert.EqualError(t, (&onlinestore.DeleteThemeAssetAPIReq{}).Verify(), "ThemeId is required")
	assert.EqualError(t, (&onlinestore.DeleteThemeAssetAPIReq{ThemeId: "x"}).Verify(), "AssetKey is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Redirect APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestCreateARedirect(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"redirect":{"id":"63a95473609552261f00942c","path":"/path/A","target":"/path/B"}}`

	httpmock.RegisterResponder("POST",
		onlineStoreURL(cli, "redirects/redirect.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.CreateARedirectAPIReq{Path: "/path/A", Target: "/path/B"}
	apiResp := &onlinestore.CreateARedirectAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "63a95473609552261f00942c", apiResp.Redirect.Id)
	assert.Equal(t, "/path/A", apiResp.Redirect.Path)
}

func TestCreateARedirect_MissingRequired(t *testing.T) {
	assert.EqualError(t, (&onlinestore.CreateARedirectAPIReq{}).Verify(), "Path is required")
	assert.EqualError(t, (&onlinestore.CreateARedirectAPIReq{Path: "/a"}).Verify(), "Target is required")
}

func TestGetAListOfRedirect(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"redirects":[{"id":"63a95473609552261f00942c","path":"/path/A","target":"/path/B"}]}`

	httpmock.RegisterResponder("GET",
		onlineStoreURL(cli, "redirects/redirect.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.GetAListOfRedirectAPIReq{}
	apiResp := &onlinestore.GetAListOfRedirectAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Redirects, 1)
	assert.Equal(t, "63a95473609552261f00942c", apiResp.Redirects[0].Id)
}

func TestQuerySingleRedirectDetails(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	redirectId := "62f372efcb98825e786f4196"
	mockResp := `{"redirect":{"id":"62f372efcb98825e786f4196","path":"/products/A","target":"/products/B"}}`

	httpmock.RegisterResponder("GET",
		onlineStoreURL(cli, fmt.Sprintf("redirects/redirect/%s.json", redirectId)),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.QuerySingleRedirectDetailsAPIReq{RedirectId: redirectId}
	apiResp := &onlinestore.QuerySingleRedirectDetailsAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, redirectId, apiResp.Redirect.Id)
	assert.Equal(t, "/products/A", apiResp.Redirect.Path)
}

func TestQuerySingleRedirectDetails_MissingRedirectId(t *testing.T) {
	err := (&onlinestore.QuerySingleRedirectDetailsAPIReq{}).Verify()
	assert.EqualError(t, err, "RedirectId is required")
}

func TestUpdateRedirect(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"redirect":{"id":"63a95473609552261f00942c","path":"/products/A","target":"/products/B"}}`

	httpmock.RegisterResponder("PUT",
		onlineStoreURL(cli, "redirects/redirect.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.UpdateRedirectAPIReq{Id: "63a95473609552261f00942c", Path: "/products/A", Target: "/products/B"}
	apiResp := &onlinestore.UpdateRedirectAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "63a95473609552261f00942c", apiResp.Redirect.Id)
}

func TestUpdateRedirect_MissingId(t *testing.T) {
	err := (&onlinestore.UpdateRedirectAPIReq{}).Verify()
	assert.EqualError(t, err, "Id is required")
}

func TestRemoveARedirect(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	redirectId := "62f372efcb98825e786f4196"

	httpmock.RegisterResponder("DELETE",
		onlineStoreURL(cli, fmt.Sprintf("redirects/%s.json", redirectId)),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &onlinestore.RemoveARedirectAPIReq{RedirectId: redirectId}
	apiResp := &onlinestore.RemoveARedirectAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}

func TestRemoveARedirect_MissingRedirectId(t *testing.T) {
	err := (&onlinestore.RemoveARedirectAPIReq{}).Verify()
	assert.EqualError(t, err, "RedirectId is required")
}

func TestStatisticalRedirectQuantity(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"count":7}`

	httpmock.RegisterResponder("GET",
		onlineStoreURL(cli, "redirects/count.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.StatisticalRedirectQuantityAPIReq{}
	apiResp := &onlinestore.StatisticalRedirectQuantityAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, int64(7), apiResp.Count)
}

// ══════════════════════════════════════════════════════════════════════════════
// Policy APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestQueryPolicyPageInformation(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"policies":[{"id":"6497413009837170726","handle":"privacy-policy","title":"PRIVACY POLICY","body_html":"<p>Free refund in 7 days</p>"}]}`

	httpmock.RegisterResponder("GET",
		onlineStoreURL(cli, "store/policy/policies.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.QueryPolicyPageInformationAPIReq{}
	apiResp := &onlinestore.QueryPolicyPageInformationAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Len(t, apiResp.Policies, 1)
	assert.Equal(t, "privacy-policy", apiResp.Policies[0].Handle)
	assert.Equal(t, "PRIVACY POLICY", apiResp.Policies[0].Title)
}

func TestUpdateAStorePolicyPage(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"policy":{"handle":"refund-policy","title":"Refund Policy","body_html":"<p>Free refund in 7 days</p>"}}`

	httpmock.RegisterResponder("POST",
		onlineStoreURL(cli, "store/policy/policy.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.UpdateAStorePolicyPageAPIReq{Handle: "refund-policy", Title: "Refund Policy", BodyHtml: "<p>Free refund in 7 days</p>"}
	apiResp := &onlinestore.UpdateAStorePolicyPageAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "refund-policy", apiResp.Policy.Handle)
	assert.Equal(t, "Refund Policy", apiResp.Policy.Title)
}

func TestUpdateAStorePolicyPage_MissingHandle(t *testing.T) {
	err := (&onlinestore.UpdateAStorePolicyPageAPIReq{}).Verify()
	assert.EqualError(t, err, "Handle is required")
}

// ══════════════════════════════════════════════════════════════════════════════
// Custom Page APIs
// ══════════════════════════════════════════════════════════════════════════════

func TestCreateACustomPage(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"data":{"id":"6743314731912662424","title":"Contact us","handle":"contact-us","author":"Alvin"},"msg":"success"}`

	httpmock.RegisterResponder("POST",
		onlineStoreURL(cli, "store/page/customize.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.CreateACustomPageAPIReq{
		Req: onlinestore.CreateACustomPageAPIReqReq{Title: "Contact us", Handle: "contact-us", Author: "Alvin"},
	}
	apiResp := &onlinestore.CreateACustomPageAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "6743314731912662424", apiResp.Data.Id)
	assert.Equal(t, "Contact us", apiResp.Data.Title)
}

func TestCountTheNumberOfCustomPages(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"data":{"count":12},"msg":"success"}`

	httpmock.RegisterResponder("POST",
		onlineStoreURL(cli, "store/page/customize/count.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.CountTheNumberOfCustomPagesAPIReq{}
	apiResp := &onlinestore.CountTheNumberOfCustomPagesAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "success", apiResp.Msg)
}

func TestGetCustomPages(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"data":{"list":[{"id":"6743314731912662424","title":"Contact us"}],"total":1},"msg":"success"}`

	httpmock.RegisterResponder("POST",
		onlineStoreURL(cli, "store/page/customize/list.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.GetCustomPagesAPIReq{}
	apiResp := &onlinestore.GetCustomPagesAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "success", apiResp.Msg)
}

func TestQueryPageSDetailInformation(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"data":{"id":"6743314731912662424","title":"Contact us","handle":"contact-us"},"msg":"success"}`

	httpmock.RegisterResponder("POST",
		onlineStoreURL(cli, "store/page/customize/id.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.QueryPageSDetailInformationAPIReq{}
	apiResp := &onlinestore.QueryPageSDetailInformationAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "success", apiResp.Msg)
}

func TestUpdateACustomPage(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	mockResp := `{"data":{"id":"6743314731912662424","title":"Updated Page"},"msg":"success"}`

	httpmock.RegisterResponder("PUT",
		onlineStoreURL(cli, "store/page/customize.json"),
		httpmock.NewStringResponder(200, mockResp))

	apiReq := &onlinestore.UpdateACustomPageAPIReq{}
	apiResp := &onlinestore.UpdateACustomPageAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
	assert.Equal(t, "success", apiResp.Msg)
}

func TestDeleteACustomPage(t *testing.T) {
	test.SetupWithVersion(ApiVersion)
	defer test.Teardown()
	cli := test.GetClient()

	httpmock.RegisterResponder("DELETE",
		onlineStoreURL(cli, "store/page/customize.json"),
		httpmock.NewStringResponder(200, `{}`))

	apiReq := &onlinestore.DeleteACustomPageAPIReq{}
	apiResp := &onlinestore.DeleteACustomPageAPIResp{}
	err := cli.Call(context.Background(), apiReq, apiResp)

	assert.NoError(t, err)
}
