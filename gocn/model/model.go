package model

import "time"

type GetGoCNDirResp struct {
	Payload struct {
		AllShortcutsEnabled bool   `json:"allShortcutsEnabled"`
		Path                string `json:"path"`
		Repo                struct {
			Id                 int       `json:"id"`
			DefaultBranch      string    `json:"defaultBranch"`
			Name               string    `json:"name"`
			OwnerLogin         string    `json:"ownerLogin"`
			CurrentUserCanPush bool      `json:"currentUserCanPush"`
			IsFork             bool      `json:"isFork"`
			IsEmpty            bool      `json:"isEmpty"`
			CreatedAt          time.Time `json:"createdAt"`
			OwnerAvatar        string    `json:"ownerAvatar"`
			Public             bool      `json:"public"`
			Private            bool      `json:"private"`
			IsOrgOwned         bool      `json:"isOrgOwned"`
		} `json:"repo"`
		CurrentUser interface{} `json:"currentUser"`
		RefInfo     struct {
			Name         string `json:"name"`
			ListCacheKey string `json:"listCacheKey"`
			CanEdit      bool   `json:"canEdit"`
			RefType      string `json:"refType"`
			CurrentOid   string `json:"currentOid"`
		} `json:"refInfo"`
		Tree struct {
			Items                          []Item      `json:"items"`
			TemplateDirectorySuggestionUrl interface{} `json:"templateDirectorySuggestionUrl"`
			Readme                         interface{} `json:"readme"`
			TotalCount                     int         `json:"totalCount"`
			ShowBranchInfobar              bool        `json:"showBranchInfobar"`
		} `json:"tree"`
		FileTree struct {
			Field1 struct {
				Items []struct {
					Name        string `json:"name"`
					Path        string `json:"path"`
					ContentType string `json:"contentType"`
				} `json:"items"`
				TotalCount int `json:"totalCount"`
			} `json:""`
		} `json:"fileTree"`
		FileTreeProcessingTime float64       `json:"fileTreeProcessingTime"`
		FoldersToFetch         []interface{} `json:"foldersToFetch"`
		ShowSurveyBanner       bool          `json:"showSurveyBanner"`
		ShowCodeNavSurvey      bool          `json:"showCodeNavSurvey"`
		CsrfTokens             struct {
			GocnNewsBranches struct {
				Post string `json:"post"`
			} `json:"/gocn/news/branches"`
			GocnNewsBranchesFetchAndMergeMaster struct {
				Post string `json:"post"`
			} `json:"/gocn/news/branches/fetch_and_merge/master"`
			GocnNewsBranchesFetchAndMergeMasterDiscardChangesTrue struct {
				Post string `json:"post"`
			} `json:"/gocn/news/branches/fetch_and_merge/master?discard_changes=true"`
		} `json:"csrf_tokens"`
	} `json:"payload"`
	Title  string `json:"title"`
	Locale string `json:"locale"`
}

type Item struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	ContentType string `json:"contentType"`
}
