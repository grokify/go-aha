package ideas

import "github.com/grokify/gocharts/v2/data/table"

const (
	IdeaAdminResponse                  = "Idea admin response"
	IdeaAssignedTo                     = "Idea assigned to"
	IdeaAssignedToEmail                = "Idea assigned to email"
	IdeaCategories                     = "Idea categories"
	IdeaCreatedBy                      = "Idea created by"
	IdeaCreatedByEmailDomain           = "Idea created by email domain"
	IdeaCreatedDate                    = "Idea created date"
	IdeaCreatedMonth                   = "Idea created month"
	IdeaCreatedYear                    = "Idea created year"
	IdeaCreatedDescription             = "Idea description"
	IdeaID                             = "Idea ID"
	IdeaInternalComments               = "Idea internal comments"
	IdeaInternalCommentsCount          = "Idea internal comments count"
	IdeaLastActiveDate                 = "Idea last active date"
	IdeaLastInternalComment            = "Idea last internal comment"
	IdeaLastPinnedDate                 = "Idea last pinned date"
	IdeaLastPortalComment              = "Idea last portal comment"
	IdeaLastPortalCommentDate          = "Idea last portal comment date"
	IdeaLastStatusChangeDate           = "Idea last status change date"
	IdeaLastStatusChangeMonth          = "Idea last status change month"
	IdeaLastStatusChangeTime           = "Idea last status change time"
	IdeaLastStatusChangeYear           = "Idea last status change year"
	IdeaLastVoteDate                   = "Idea last vote date"
	IdeaName                           = "Idea name"
	IdeaNameWithPortalURL              = "Idea name and URL"
	IdeaPinned                         = "Idea pinned"
	IdeaPortalComments                 = "Idea portal comments"
	IdeaPortalCommentsCount            = "Idea portal comments count"
	IdeaReference                      = "Idea reference"
	IdeaStatus                         = "Idea status"
	IdeaStatusCategory                 = "Idea status category"
	IdeaSubmittedPortal                = "Idea submitted portal"
	IdeaSubmittedPortalURL             = "Idea submitted portal URL"
	IdeaTags                           = "Idea tags"
	IdeaTimeInCurrentStatus            = "Idea time in current status"
	IdeaTimeinStatusCategoryDone       = "Idea time in satatus category Done"
	IdeaTimeinStatusCategoryInProgress = "Idea time in satatus category In progress"
	IdeaTimeinStatusCategoryNotStarted = "Idea time in satatus category Not started"
	IdeaURL                            = "Idea URL"
	IdeaURLInSubmitedPortal            = "Idea URL in submitted portal"
	IdeaVisibility                     = "Idea visibility"
	IdeaVoterEmailDomains              = "Idea voter email domains"
	IdeaVotes                          = "Idea votes"
	IdeaVotesOpportnityValue           = "Idea votes opportunity value"
	MergedIdeas                        = "Merged ideas"
	MergedIntoIdea                     = "Merged into idea"
	PromotedTo                         = "Promoted to"
	RecordLinks                        = "Record links"
	Watchers                           = "Watchers"

	RowNumber = "No."
)

func Columns() []string {
	return []string{
		RowNumber,
		IdeaID,
		IdeaReference,
		IdeaCategories,
		IdeaName,
		IdeaStatus,
		IdeaStatusCategory,
		IdeaVotes,

		IdeaCreatedBy,
		IdeaCreatedByEmailDomain,
		IdeaVoterEmailDomains,

		IdeaLastActiveDate,
		IdeaLastStatusChangeDate,
		IdeaLastVoteDate,
	}
}

func ColumnsShort() ([]string, map[int]string) {
	return []string{
			RowNumber,
			IdeaCategories,
			IdeaNameWithPortalURL,
			IdeaStatus,
			IdeaVotes},
		map[int]string{
			0: table.FormatInt,
			2: table.FormatURL}
}
