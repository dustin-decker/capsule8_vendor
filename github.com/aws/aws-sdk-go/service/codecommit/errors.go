// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

const (

	// ErrCodeActorDoesNotExistException for service response error code
	// "ActorDoesNotExistException".
	//
	// The specified Amazon Resource Name (ARN) does not exist in the AWS account.
	ErrCodeActorDoesNotExistException = "ActorDoesNotExistException"

	// ErrCodeAuthorDoesNotExistException for service response error code
	// "AuthorDoesNotExistException".
	//
	// The specified Amazon Resource Name (ARN) does not exist in the AWS account.
	ErrCodeAuthorDoesNotExistException = "AuthorDoesNotExistException"

	// ErrCodeBeforeCommitIdAndAfterCommitIdAreSameException for service response error code
	// "BeforeCommitIdAndAfterCommitIdAreSameException".
	//
	// The before commit ID and the after commit ID are the same, which is not valid.
	// The before commit ID and the after commit ID must be different commit IDs.
	ErrCodeBeforeCommitIdAndAfterCommitIdAreSameException = "BeforeCommitIdAndAfterCommitIdAreSameException"

	// ErrCodeBlobIdDoesNotExistException for service response error code
	// "BlobIdDoesNotExistException".
	//
	// The specified blob does not exist.
	ErrCodeBlobIdDoesNotExistException = "BlobIdDoesNotExistException"

	// ErrCodeBlobIdRequiredException for service response error code
	// "BlobIdRequiredException".
	//
	// A blob ID is required but was not specified.
	ErrCodeBlobIdRequiredException = "BlobIdRequiredException"

	// ErrCodeBranchDoesNotExistException for service response error code
	// "BranchDoesNotExistException".
	//
	// The specified branch does not exist.
	ErrCodeBranchDoesNotExistException = "BranchDoesNotExistException"

	// ErrCodeBranchNameExistsException for service response error code
	// "BranchNameExistsException".
	//
	// The specified branch name already exists.
	ErrCodeBranchNameExistsException = "BranchNameExistsException"

	// ErrCodeBranchNameRequiredException for service response error code
	// "BranchNameRequiredException".
	//
	// A branch name is required but was not specified.
	ErrCodeBranchNameRequiredException = "BranchNameRequiredException"

	// ErrCodeClientRequestTokenRequiredException for service response error code
	// "ClientRequestTokenRequiredException".
	//
	// A client request token is required. A client request token is an unique,
	// client-generated idempotency token that when provided in a request, ensures
	// the request cannot be repeated with a changed parameter. If a request is
	// received with the same parameters and a token is included, the request will
	// return information about the initial request that used that token.
	ErrCodeClientRequestTokenRequiredException = "ClientRequestTokenRequiredException"

	// ErrCodeCommentContentRequiredException for service response error code
	// "CommentContentRequiredException".
	//
	// The comment is empty. You must provide some content for a comment. The content
	// cannot be null.
	ErrCodeCommentContentRequiredException = "CommentContentRequiredException"

	// ErrCodeCommentContentSizeLimitExceededException for service response error code
	// "CommentContentSizeLimitExceededException".
	//
	// The comment is too large. Comments are limited to 1,000 characters.
	ErrCodeCommentContentSizeLimitExceededException = "CommentContentSizeLimitExceededException"

	// ErrCodeCommentDeletedException for service response error code
	// "CommentDeletedException".
	//
	// This comment has already been deleted. You cannot edit or delete a deleted
	// comment.
	ErrCodeCommentDeletedException = "CommentDeletedException"

	// ErrCodeCommentDoesNotExistException for service response error code
	// "CommentDoesNotExistException".
	//
	// No comment exists with the provided ID. Verify that you have provided the
	// correct ID, and then try again.
	ErrCodeCommentDoesNotExistException = "CommentDoesNotExistException"

	// ErrCodeCommentIdRequiredException for service response error code
	// "CommentIdRequiredException".
	//
	// The comment ID is missing or null. A comment ID is required.
	ErrCodeCommentIdRequiredException = "CommentIdRequiredException"

	// ErrCodeCommentNotCreatedByCallerException for service response error code
	// "CommentNotCreatedByCallerException".
	//
	// You cannot modify or delete this comment. Only comment authors can modify
	// or delete their comments.
	ErrCodeCommentNotCreatedByCallerException = "CommentNotCreatedByCallerException"

	// ErrCodeCommitDoesNotExistException for service response error code
	// "CommitDoesNotExistException".
	//
	// The specified commit does not exist or no commit was specified, and the specified
	// repository has no default branch.
	ErrCodeCommitDoesNotExistException = "CommitDoesNotExistException"

	// ErrCodeCommitIdDoesNotExistException for service response error code
	// "CommitIdDoesNotExistException".
	//
	// The specified commit ID does not exist.
	ErrCodeCommitIdDoesNotExistException = "CommitIdDoesNotExistException"

	// ErrCodeCommitIdRequiredException for service response error code
	// "CommitIdRequiredException".
	//
	// A commit ID was not specified.
	ErrCodeCommitIdRequiredException = "CommitIdRequiredException"

	// ErrCodeCommitRequiredException for service response error code
	// "CommitRequiredException".
	//
	// A commit was not specified.
	ErrCodeCommitRequiredException = "CommitRequiredException"

	// ErrCodeDefaultBranchCannotBeDeletedException for service response error code
	// "DefaultBranchCannotBeDeletedException".
	//
	// The specified branch is the default branch for the repository, and cannot
	// be deleted. To delete this branch, you must first set another branch as the
	// default branch.
	ErrCodeDefaultBranchCannotBeDeletedException = "DefaultBranchCannotBeDeletedException"

	// ErrCodeEncryptionIntegrityChecksFailedException for service response error code
	// "EncryptionIntegrityChecksFailedException".
	//
	// An encryption integrity check failed.
	ErrCodeEncryptionIntegrityChecksFailedException = "EncryptionIntegrityChecksFailedException"

	// ErrCodeEncryptionKeyAccessDeniedException for service response error code
	// "EncryptionKeyAccessDeniedException".
	//
	// An encryption key could not be accessed.
	ErrCodeEncryptionKeyAccessDeniedException = "EncryptionKeyAccessDeniedException"

	// ErrCodeEncryptionKeyDisabledException for service response error code
	// "EncryptionKeyDisabledException".
	//
	// The encryption key is disabled.
	ErrCodeEncryptionKeyDisabledException = "EncryptionKeyDisabledException"

	// ErrCodeEncryptionKeyNotFoundException for service response error code
	// "EncryptionKeyNotFoundException".
	//
	// No encryption key was found.
	ErrCodeEncryptionKeyNotFoundException = "EncryptionKeyNotFoundException"

	// ErrCodeEncryptionKeyUnavailableException for service response error code
	// "EncryptionKeyUnavailableException".
	//
	// The encryption key is not available.
	ErrCodeEncryptionKeyUnavailableException = "EncryptionKeyUnavailableException"

	// ErrCodeFileTooLargeException for service response error code
	// "FileTooLargeException".
	//
	// The specified file exceeds the file size limit for AWS CodeCommit. For more
	// information about limits in AWS CodeCommit, see AWS CodeCommit User Guide
	// (http://docs.aws.amazon.com/codecommit/latest/userguide/limits.html).
	ErrCodeFileTooLargeException = "FileTooLargeException"

	// ErrCodeIdempotencyParameterMismatchException for service response error code
	// "IdempotencyParameterMismatchException".
	//
	// The client request token is not valid. Either the token is not in a valid
	// format, or the token has been used in a previous request and cannot be re-used.
	ErrCodeIdempotencyParameterMismatchException = "IdempotencyParameterMismatchException"

	// ErrCodeInvalidActorArnException for service response error code
	// "InvalidActorArnException".
	//
	// The Amazon Resource Name (ARN) is not valid. Make sure that you have provided
	// the full ARN for the user who initiated the change for the pull request,
	// and then try again.
	ErrCodeInvalidActorArnException = "InvalidActorArnException"

	// ErrCodeInvalidAuthorArnException for service response error code
	// "InvalidAuthorArnException".
	//
	// The Amazon Resource Name (ARN) is not valid. Make sure that you have provided
	// the full ARN for the author of the pull request, and then try again.
	ErrCodeInvalidAuthorArnException = "InvalidAuthorArnException"

	// ErrCodeInvalidBlobIdException for service response error code
	// "InvalidBlobIdException".
	//
	// The specified blob is not valid.
	ErrCodeInvalidBlobIdException = "InvalidBlobIdException"

	// ErrCodeInvalidBranchNameException for service response error code
	// "InvalidBranchNameException".
	//
	// The specified reference name is not valid.
	ErrCodeInvalidBranchNameException = "InvalidBranchNameException"

	// ErrCodeInvalidClientRequestTokenException for service response error code
	// "InvalidClientRequestTokenException".
	//
	// The client request token is not valid.
	ErrCodeInvalidClientRequestTokenException = "InvalidClientRequestTokenException"

	// ErrCodeInvalidCommentIdException for service response error code
	// "InvalidCommentIdException".
	//
	// The comment ID is not in a valid format. Make sure that you have provided
	// the full comment ID.
	ErrCodeInvalidCommentIdException = "InvalidCommentIdException"

	// ErrCodeInvalidCommitException for service response error code
	// "InvalidCommitException".
	//
	// The specified commit is not valid.
	ErrCodeInvalidCommitException = "InvalidCommitException"

	// ErrCodeInvalidCommitIdException for service response error code
	// "InvalidCommitIdException".
	//
	// The specified commit ID is not valid.
	ErrCodeInvalidCommitIdException = "InvalidCommitIdException"

	// ErrCodeInvalidContinuationTokenException for service response error code
	// "InvalidContinuationTokenException".
	//
	// The specified continuation token is not valid.
	ErrCodeInvalidContinuationTokenException = "InvalidContinuationTokenException"

	// ErrCodeInvalidDescriptionException for service response error code
	// "InvalidDescriptionException".
	//
	// The pull request description is not valid. Descriptions are limited to 1,000
	// characters in length.
	ErrCodeInvalidDescriptionException = "InvalidDescriptionException"

	// ErrCodeInvalidDestinationCommitSpecifierException for service response error code
	// "InvalidDestinationCommitSpecifierException".
	//
	// The destination commit specifier is not valid. You must provide a valid branch
	// name, tag, or full commit ID.
	ErrCodeInvalidDestinationCommitSpecifierException = "InvalidDestinationCommitSpecifierException"

	// ErrCodeInvalidFileLocationException for service response error code
	// "InvalidFileLocationException".
	//
	// The location of the file is not valid. Make sure that you include the extension
	// of the file as well as the file name.
	ErrCodeInvalidFileLocationException = "InvalidFileLocationException"

	// ErrCodeInvalidFilePositionException for service response error code
	// "InvalidFilePositionException".
	//
	// The position is not valid. Make sure that the line number exists in the version
	// of the file you want to comment on.
	ErrCodeInvalidFilePositionException = "InvalidFilePositionException"

	// ErrCodeInvalidMaxResultsException for service response error code
	// "InvalidMaxResultsException".
	//
	// The specified number of maximum results is not valid.
	ErrCodeInvalidMaxResultsException = "InvalidMaxResultsException"

	// ErrCodeInvalidMergeOptionException for service response error code
	// "InvalidMergeOptionException".
	//
	// The specified merge option is not valid. The only valid value is FAST_FORWARD_MERGE.
	ErrCodeInvalidMergeOptionException = "InvalidMergeOptionException"

	// ErrCodeInvalidOrderException for service response error code
	// "InvalidOrderException".
	//
	// The specified sort order is not valid.
	ErrCodeInvalidOrderException = "InvalidOrderException"

	// ErrCodeInvalidPathException for service response error code
	// "InvalidPathException".
	//
	// The specified path is not valid.
	ErrCodeInvalidPathException = "InvalidPathException"

	// ErrCodeInvalidPullRequestEventTypeException for service response error code
	// "InvalidPullRequestEventTypeException".
	//
	// The pull request event type is not valid.
	ErrCodeInvalidPullRequestEventTypeException = "InvalidPullRequestEventTypeException"

	// ErrCodeInvalidPullRequestIdException for service response error code
	// "InvalidPullRequestIdException".
	//
	// The pull request ID is not valid. Make sure that you have provided the full
	// ID and that the pull request is in the specified repository, and then try
	// again.
	ErrCodeInvalidPullRequestIdException = "InvalidPullRequestIdException"

	// ErrCodeInvalidPullRequestStatusException for service response error code
	// "InvalidPullRequestStatusException".
	//
	// The pull request status is not valid. The only valid values are OPEN and
	// CLOSED.
	ErrCodeInvalidPullRequestStatusException = "InvalidPullRequestStatusException"

	// ErrCodeInvalidPullRequestStatusUpdateException for service response error code
	// "InvalidPullRequestStatusUpdateException".
	//
	// The pull request status update is not valid. The only valid update is from
	// OPEN to CLOSED.
	ErrCodeInvalidPullRequestStatusUpdateException = "InvalidPullRequestStatusUpdateException"

	// ErrCodeInvalidReferenceNameException for service response error code
	// "InvalidReferenceNameException".
	//
	// The specified reference name format is not valid. Reference names must conform
	// to the Git references format, for example refs/heads/master. For more information,
	// see Git Internals - Git References (https://git-scm.com/book/en/v2/Git-Internals-Git-References)
	// or consult your Git documentation.
	ErrCodeInvalidReferenceNameException = "InvalidReferenceNameException"

	// ErrCodeInvalidRelativeFileVersionEnumException for service response error code
	// "InvalidRelativeFileVersionEnumException".
	//
	// Either the enum is not in a valid format, or the specified file version enum
	// is not valid in respect to the current file version.
	ErrCodeInvalidRelativeFileVersionEnumException = "InvalidRelativeFileVersionEnumException"

	// ErrCodeInvalidRepositoryDescriptionException for service response error code
	// "InvalidRepositoryDescriptionException".
	//
	// The specified repository description is not valid.
	ErrCodeInvalidRepositoryDescriptionException = "InvalidRepositoryDescriptionException"

	// ErrCodeInvalidRepositoryNameException for service response error code
	// "InvalidRepositoryNameException".
	//
	// At least one specified repository name is not valid.
	//
	// This exception only occurs when a specified repository name is not valid.
	// Other exceptions occur when a required repository parameter is missing, or
	// when a specified repository does not exist.
	ErrCodeInvalidRepositoryNameException = "InvalidRepositoryNameException"

	// ErrCodeInvalidRepositoryTriggerBranchNameException for service response error code
	// "InvalidRepositoryTriggerBranchNameException".
	//
	// One or more branch names specified for the trigger is not valid.
	ErrCodeInvalidRepositoryTriggerBranchNameException = "InvalidRepositoryTriggerBranchNameException"

	// ErrCodeInvalidRepositoryTriggerCustomDataException for service response error code
	// "InvalidRepositoryTriggerCustomDataException".
	//
	// The custom data provided for the trigger is not valid.
	ErrCodeInvalidRepositoryTriggerCustomDataException = "InvalidRepositoryTriggerCustomDataException"

	// ErrCodeInvalidRepositoryTriggerDestinationArnException for service response error code
	// "InvalidRepositoryTriggerDestinationArnException".
	//
	// The Amazon Resource Name (ARN) for the trigger is not valid for the specified
	// destination. The most common reason for this error is that the ARN does not
	// meet the requirements for the service type.
	ErrCodeInvalidRepositoryTriggerDestinationArnException = "InvalidRepositoryTriggerDestinationArnException"

	// ErrCodeInvalidRepositoryTriggerEventsException for service response error code
	// "InvalidRepositoryTriggerEventsException".
	//
	// One or more events specified for the trigger is not valid. Check to make
	// sure that all events specified match the requirements for allowed events.
	ErrCodeInvalidRepositoryTriggerEventsException = "InvalidRepositoryTriggerEventsException"

	// ErrCodeInvalidRepositoryTriggerNameException for service response error code
	// "InvalidRepositoryTriggerNameException".
	//
	// The name of the trigger is not valid.
	ErrCodeInvalidRepositoryTriggerNameException = "InvalidRepositoryTriggerNameException"

	// ErrCodeInvalidRepositoryTriggerRegionException for service response error code
	// "InvalidRepositoryTriggerRegionException".
	//
	// The region for the trigger target does not match the region for the repository.
	// Triggers must be created in the same region as the target for the trigger.
	ErrCodeInvalidRepositoryTriggerRegionException = "InvalidRepositoryTriggerRegionException"

	// ErrCodeInvalidSortByException for service response error code
	// "InvalidSortByException".
	//
	// The specified sort by value is not valid.
	ErrCodeInvalidSortByException = "InvalidSortByException"

	// ErrCodeInvalidSourceCommitSpecifierException for service response error code
	// "InvalidSourceCommitSpecifierException".
	//
	// The source commit specifier is not valid. You must provide a valid branch
	// name, tag, or full commit ID.
	ErrCodeInvalidSourceCommitSpecifierException = "InvalidSourceCommitSpecifierException"

	// ErrCodeInvalidTargetException for service response error code
	// "InvalidTargetException".
	//
	// The target for the pull request is not valid. A target must contain the full
	// values for the repository name, source branch, and destination branch for
	// the pull request.
	ErrCodeInvalidTargetException = "InvalidTargetException"

	// ErrCodeInvalidTargetsException for service response error code
	// "InvalidTargetsException".
	//
	// The targets for the pull request is not valid or not in a valid format. Targets
	// are a list of target objects. Each target object must contain the full values
	// for the repository name, source branch, and destination branch for a pull
	// request.
	ErrCodeInvalidTargetsException = "InvalidTargetsException"

	// ErrCodeInvalidTitleException for service response error code
	// "InvalidTitleException".
	//
	// The title of the pull request is not valid. Pull request titles cannot exceed
	// 100 characters in length.
	ErrCodeInvalidTitleException = "InvalidTitleException"

	// ErrCodeManualMergeRequiredException for service response error code
	// "ManualMergeRequiredException".
	//
	// The pull request cannot be merged automatically into the destination branch.
	// You must manually merge the branches and resolve any conflicts.
	ErrCodeManualMergeRequiredException = "ManualMergeRequiredException"

	// ErrCodeMaximumBranchesExceededException for service response error code
	// "MaximumBranchesExceededException".
	//
	// The number of branches for the trigger was exceeded.
	ErrCodeMaximumBranchesExceededException = "MaximumBranchesExceededException"

	// ErrCodeMaximumOpenPullRequestsExceededException for service response error code
	// "MaximumOpenPullRequestsExceededException".
	//
	// You cannot create the pull request because the repository has too many open
	// pull requests. The maximum number of open pull requests for a repository
	// is 1,000. Close one or more open pull requests, and then try again.
	ErrCodeMaximumOpenPullRequestsExceededException = "MaximumOpenPullRequestsExceededException"

	// ErrCodeMaximumRepositoryNamesExceededException for service response error code
	// "MaximumRepositoryNamesExceededException".
	//
	// The maximum number of allowed repository names was exceeded. Currently, this
	// number is 25.
	ErrCodeMaximumRepositoryNamesExceededException = "MaximumRepositoryNamesExceededException"

	// ErrCodeMaximumRepositoryTriggersExceededException for service response error code
	// "MaximumRepositoryTriggersExceededException".
	//
	// The number of triggers allowed for the repository was exceeded.
	ErrCodeMaximumRepositoryTriggersExceededException = "MaximumRepositoryTriggersExceededException"

	// ErrCodeMergeOptionRequiredException for service response error code
	// "MergeOptionRequiredException".
	//
	// A merge option or stategy is required, and none was provided.
	ErrCodeMergeOptionRequiredException = "MergeOptionRequiredException"

	// ErrCodeMultipleRepositoriesInPullRequestException for service response error code
	// "MultipleRepositoriesInPullRequestException".
	//
	// You cannot include more than one repository in a pull request. Make sure
	// you have specified only one repository name in your request, and then try
	// again.
	ErrCodeMultipleRepositoriesInPullRequestException = "MultipleRepositoriesInPullRequestException"

	// ErrCodePathDoesNotExistException for service response error code
	// "PathDoesNotExistException".
	//
	// The specified path does not exist.
	ErrCodePathDoesNotExistException = "PathDoesNotExistException"

	// ErrCodePathRequiredException for service response error code
	// "PathRequiredException".
	//
	// The filePath for a location cannot be empty or null.
	ErrCodePathRequiredException = "PathRequiredException"

	// ErrCodePullRequestAlreadyClosedException for service response error code
	// "PullRequestAlreadyClosedException".
	//
	// The pull request status cannot be updated because it is already closed.
	ErrCodePullRequestAlreadyClosedException = "PullRequestAlreadyClosedException"

	// ErrCodePullRequestDoesNotExistException for service response error code
	// "PullRequestDoesNotExistException".
	//
	// The pull request ID could not be found. Make sure that you have specified
	// the correct repository name and pull request ID, and then try again.
	ErrCodePullRequestDoesNotExistException = "PullRequestDoesNotExistException"

	// ErrCodePullRequestIdRequiredException for service response error code
	// "PullRequestIdRequiredException".
	//
	// A pull request ID is required, but none was provided.
	ErrCodePullRequestIdRequiredException = "PullRequestIdRequiredException"

	// ErrCodePullRequestStatusRequiredException for service response error code
	// "PullRequestStatusRequiredException".
	//
	// A pull request status is required, but none was provided.
	ErrCodePullRequestStatusRequiredException = "PullRequestStatusRequiredException"

	// ErrCodeReferenceDoesNotExistException for service response error code
	// "ReferenceDoesNotExistException".
	//
	// The specified reference does not exist. You must provide a full commit ID.
	ErrCodeReferenceDoesNotExistException = "ReferenceDoesNotExistException"

	// ErrCodeReferenceNameRequiredException for service response error code
	// "ReferenceNameRequiredException".
	//
	// A reference name is required, but none was provided.
	ErrCodeReferenceNameRequiredException = "ReferenceNameRequiredException"

	// ErrCodeReferenceTypeNotSupportedException for service response error code
	// "ReferenceTypeNotSupportedException".
	//
	// The specified reference is not a supported type.
	ErrCodeReferenceTypeNotSupportedException = "ReferenceTypeNotSupportedException"

	// ErrCodeRepositoryDoesNotExistException for service response error code
	// "RepositoryDoesNotExistException".
	//
	// The specified repository does not exist.
	ErrCodeRepositoryDoesNotExistException = "RepositoryDoesNotExistException"

	// ErrCodeRepositoryLimitExceededException for service response error code
	// "RepositoryLimitExceededException".
	//
	// A repository resource limit was exceeded.
	ErrCodeRepositoryLimitExceededException = "RepositoryLimitExceededException"

	// ErrCodeRepositoryNameExistsException for service response error code
	// "RepositoryNameExistsException".
	//
	// The specified repository name already exists.
	ErrCodeRepositoryNameExistsException = "RepositoryNameExistsException"

	// ErrCodeRepositoryNameRequiredException for service response error code
	// "RepositoryNameRequiredException".
	//
	// A repository name is required but was not specified.
	ErrCodeRepositoryNameRequiredException = "RepositoryNameRequiredException"

	// ErrCodeRepositoryNamesRequiredException for service response error code
	// "RepositoryNamesRequiredException".
	//
	// A repository names object is required but was not specified.
	ErrCodeRepositoryNamesRequiredException = "RepositoryNamesRequiredException"

	// ErrCodeRepositoryNotAssociatedWithPullRequestException for service response error code
	// "RepositoryNotAssociatedWithPullRequestException".
	//
	// The repository does not contain any pull requests with that pull request
	// ID. Check to make sure you have provided the correct repository name for
	// the pull request.
	ErrCodeRepositoryNotAssociatedWithPullRequestException = "RepositoryNotAssociatedWithPullRequestException"

	// ErrCodeRepositoryTriggerBranchNameListRequiredException for service response error code
	// "RepositoryTriggerBranchNameListRequiredException".
	//
	// At least one branch name is required but was not specified in the trigger
	// configuration.
	ErrCodeRepositoryTriggerBranchNameListRequiredException = "RepositoryTriggerBranchNameListRequiredException"

	// ErrCodeRepositoryTriggerDestinationArnRequiredException for service response error code
	// "RepositoryTriggerDestinationArnRequiredException".
	//
	// A destination ARN for the target service for the trigger is required but
	// was not specified.
	ErrCodeRepositoryTriggerDestinationArnRequiredException = "RepositoryTriggerDestinationArnRequiredException"

	// ErrCodeRepositoryTriggerEventsListRequiredException for service response error code
	// "RepositoryTriggerEventsListRequiredException".
	//
	// At least one event for the trigger is required but was not specified.
	ErrCodeRepositoryTriggerEventsListRequiredException = "RepositoryTriggerEventsListRequiredException"

	// ErrCodeRepositoryTriggerNameRequiredException for service response error code
	// "RepositoryTriggerNameRequiredException".
	//
	// A name for the trigger is required but was not specified.
	ErrCodeRepositoryTriggerNameRequiredException = "RepositoryTriggerNameRequiredException"

	// ErrCodeRepositoryTriggersListRequiredException for service response error code
	// "RepositoryTriggersListRequiredException".
	//
	// The list of triggers for the repository is required but was not specified.
	ErrCodeRepositoryTriggersListRequiredException = "RepositoryTriggersListRequiredException"

	// ErrCodeSourceAndDestinationAreSameException for service response error code
	// "SourceAndDestinationAreSameException".
	//
	// The source branch and the destination branch for the pull request are the
	// same. You must specify different branches for the source and destination.
	ErrCodeSourceAndDestinationAreSameException = "SourceAndDestinationAreSameException"

	// ErrCodeTargetRequiredException for service response error code
	// "TargetRequiredException".
	//
	// A pull request target is required. It cannot be empty or null. A pull request
	// target must contain the full values for the repository name, source branch,
	// and destination branch for the pull request.
	ErrCodeTargetRequiredException = "TargetRequiredException"

	// ErrCodeTargetsRequiredException for service response error code
	// "TargetsRequiredException".
	//
	// An array of target objects is required. It cannot be empty or null.
	ErrCodeTargetsRequiredException = "TargetsRequiredException"

	// ErrCodeTipOfSourceReferenceIsDifferentException for service response error code
	// "TipOfSourceReferenceIsDifferentException".
	//
	// The tip of the source branch in the destination repository does not match
	// the tip of the source branch specified in your request. The pull request
	// might have been updated. Make sure that you have the latest changes.
	ErrCodeTipOfSourceReferenceIsDifferentException = "TipOfSourceReferenceIsDifferentException"

	// ErrCodeTipsDivergenceExceededException for service response error code
	// "TipsDivergenceExceededException".
	//
	// The divergence between the tips of the provided commit specifiers is too
	// great to determine whether there might be any merge conflicts. Locally compare
	// the specifiers using git diff or a diff tool.
	ErrCodeTipsDivergenceExceededException = "TipsDivergenceExceededException"

	// ErrCodeTitleRequiredException for service response error code
	// "TitleRequiredException".
	//
	// A pull request title is required. It cannot be empty or null.
	ErrCodeTitleRequiredException = "TitleRequiredException"
)
