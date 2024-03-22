/*
繁體中文翻譯詞彙選擇主要依據下列來源：
The selection of Traditional Chinese translation vocabulary is mainly based on
the following sources:

 1. GitLab: 其介面有相當完整的繁體中文翻譯，但缺少一些本地端功能的對照，例如 stash。

 2. Pro Git: Git 的權威參考用書，可惜繁中部分翻譯僅約一半。
    https://git-scm.com/book/zh-tw/v2

 3. Microsoft 語言入口網站 (Visual Studio)
    https://www.microsoft.com/zh-tw/language/

### Glossary ###

	譯文中括號內文字會依語境添加或省略。

	Repository		版本庫
	Amend			修改
	Checkout		檢出
	Cherry-pick		揀選
	Diff			差異
	Discard			捨棄
	Drop [stash]	捨棄
	Fast-forward	快轉 (Fast-forward)
	Fetch			擷取
	Fixup			修復 (Fixup)
	Patch			補丁
	Pop [stash]		還原
	Rebase			變基 (Rebase)
	Reset			重設
	Revert			還原
	Reword			改寫
	Squash			壓縮 (Squash)
	Stage			預存 (Stage)
	Stash			收藏 (Stash)
*/
package i18n

const traditionalChineseIntroPopupMessage = `
感謝使用 lazygit！這裡有一些資源可供參考：

 1) 📺lazygit 教學📺：
      https://youtu.be/CPLdltN7wgE

 2) 📣釋出說明📣：
      https://github.com/jesseduffield/lazygit/releases

 3) 💖如果你想要貢獻一份心力你可以💖：
    改進 lazygit 原始碼：https://github.com/jesseduffield/lazygit
    按右下角的捐款斗內我們
    或單存添加 lazygit 到你的 star 清單內以增加曝光度都能大力的幫助我們！
`

const traditionalChineseDeprecatedEditConfigWarning = `
### Deprecated config warning ###

以下設定已被取代並將於未來版本中刪除：
{{configs}}

編輯器設定教學：

  https://github.com/jesseduffield/lazygit/blob/master/docs/Config.md#configuring-file-editing

`

// exporting this so we can use it in tests
func traditionalChineseTranslationSet() TranslationSet {
	return TranslationSet{
		NotEnoughSpace:                   "無足夠空間顯示面板",
		DiffTitle:                        "差異",
		FilesTitle:                       "檔案",
		BranchesTitle:                    "分支",
		CommitsTitle:                     "提交",
		StashTitle:                       "收藏 (Stash)",
		SnakeTitle:                       "貪食蛇",
		EasterEgg:                        "彩蛋",
		UnstagedChanges:                  "未預存變更",
		StagedChanges:                    "已預存變更",
		MainTitle:                        "主要",
		MergeConfirmTitle:                "合併",
		StagingTitle:                     "主面板（預存）",
		MergingTitle:                     "主面板（合併）",
		NormalTitle:                      "主面板（一般）",
		LogTitle:                         "版本記錄",
		CommitSummary:                    "提交摘要",
		CredentialsUsername:              "使用者名稱",
		CredentialsPassword:              "密碼",
		CredentialsPassphrase:            "SSH 金鑰密語",
		CredentialsPIN:                   "SSH 金鑰 PIN 碼",
		PassUnameWrong:                   "密碼、密語或使用者名稱錯誤",
		Commit:                           "提交變更",
		AmendLastCommit:                  "修改上次提交",
		AmendLastCommitTitle:             "修改上次提交",
		SureToAmend:                      "是否確定要修改上次提交？之後你可以從提交面板中再次更改此次提交的訊息。",
		NoCommitToAmend:                  "沒有可以修改的提交。",
		CommitChangesWithEditor:          "使用 git 編輯器提交變更",
		StatusTitle:                      "狀態",
		Menu:                             "選單",
		Execute:                          "執行",
		Stage:                            "切換預存",
		ToggleStagedAll:                  "全部預存/取消預存",
		ToggleTreeView:                   "顯示檔案樹狀視圖",
		OpenMergeTool:                    "開啟外部合併工具 (git mergetool)",
		Refresh:                          "重新整理",
		Push:                             "推送",
		Pull:                             "拉取",
		Scroll:                           "捲動",
		MergeConflictsTitle:              "合併衝突",
		Checkout:                         "檢出",
		FileFilter:                       "篩選檔案 (預存/未預存)",
		FilterStagedFiles:                "僅顯示預存的檔案",
		FilterUnstagedFiles:              "僅顯示未預存的檔案",
		ResetFilter:                      "重設篩選",
		NoChangedFiles:                   "沒有變更的檔案",
		SoftReset:                        "軟重設",
		AlreadyCheckedOutBranch:          "你已經檢出這個分支了",
		SureForceCheckout:                "是否強制檢出？這將會使你失去本地的所有更改",
		ForceCheckoutBranch:              "強制檢出分支",
		BranchName:                       "分支名稱",
		NewBranchNameBranchOff:           "新的分支名稱 (根據 '{{.branchName}}' 分支創建)",
		CantDeleteCheckOutBranch:         "無法刪除已檢出的分支！",
		ForceDeleteBranchMessage:         "'{{.selectedBranchName}}' 分支尚未完全合併。是否刪除？",
		RebaseBranch:                     "將已檢出的分支變基至此分支",
		CantRebaseOntoSelf:               "無法將分支變基至自己",
		CantMergeBranchIntoItself:        "無法將一個分支合併至自己",
		ForceCheckout:                    "強制檢出",
		CheckoutByName:                   "根據名稱檢出",
		NewBranch:                        "新分支",
		NoBranchesThisRepo:               "這個版本庫中沒有分支",
		CommitWithoutMessageErr:          "沒有提交訊息，無法提交",
		Close:                            "關閉",
		CloseCancel:                      "關閉/取消",
		Confirm:                          "確認",
		Quit:                             "結束",
		NoCommitsThisBranch:              "這個分支沒有提交",
		UpdateRefHere:                    "在這裡更新 '{{.ref}}' 分支",
		CannotSquashOrFixupFirstCommit:   "沒有可以壓縮的提交",
		Fixup:                            "修復 (Fixup)",
		SureFixupThisCommit:              "是否對此提交進行 '修復' ？ 其將被合併於以下之提交中",
		SureSquashThisCommit:             "是否要把這個提交壓縮到下面的提交中？",
		Squash:                           "壓縮 (Squash)",
		PickCommitTooltip:                "挑選提交 (於變基過程中)",
		RevertCommit:                     "還原提交",
		Reword:                           "改寫提交",
		DropCommit:                       "刪除提交",
		MoveDownCommit:                   "向下移動提交",
		MoveUpCommit:                     "向上移動提交",
		EditCommitTooltip:                "編輯提交",
		AmendCommitTooltip:               "使用已預存的更改修正提交",
		ResetAuthor:                      "重設作者",
		SetAuthor:                        "設定作者",
		AmendCommitAttribute:             "設定/重設提交作者",
		SetAuthorPromptTitle:             "設定作者（格式：「姓名 <電子郵件>」）",
		SureResetCommitAuthor:            "為了符合已配置的使用者，此作者的提交欄位以及時間戳將被更新。是否繼續？",
		RewordCommitEditor:               "使用編輯器改寫提交",
		Error:                            "錯誤",
		PickHunk:                         "挑選程式碼片段",
		PickAllHunks:                     "挑選所有程式碼片段",
		Undo:                             "復原",
		UndoReflog:                       "復原",
		RedoReflog:                       "取消復原",
		UndoTooltip:                      "將使用 reflog 確任 git 指令以復原。這不包括工作區更改；只考慮提交。",
		RedoTooltip:                      "將使用 reflog 確任 git 指令以重作。這不包括工作區更改；只考慮提交。",
		DiscardAllTooltip:                "捨棄 '{{.path}}' 預存/未預存更改。",
		DiscardUnstagedTooltip:           "捨棄 '{{.path}}' 未預存更改。",
		Pop:                              "還原",
		Drop:                             "捨棄",
		Apply:                            "套用",
		NoStashEntries:                   "沒有收藏記錄",
		StashDrop:                        "放棄收藏記錄",
		SureDropStashEntry:               "是否捨棄這條收藏記錄？",
		StashPop:                         "還原收藏記錄",
		SurePopStashEntry:                "是否從收藏中還原這個記錄？",
		StashApply:                       "套用收藏記錄",
		SureApplyStashEntry:              "是否套用這個收藏記錄？",
		NoTrackedStagedFilesStash:        "你沒有被追蹤的、預存的檔案可進行收藏",
		NoFilesToStash:                   "沒有檔案可以進行收藏",
		StashChanges:                     "安置現有變更到收藏中",
		RenameStash:                      "重新命名收藏",
		RenameStashPrompt:                "重新命名收藏：{{.stashName}}",
		OpenConfig:                       "開啟設定檔案",
		EditConfig:                       "編輯設定檔案",
		ForcePush:                        "強制推送",
		ForcePushPrompt:                  "你的分支與遠端分支分岔。按 'ESC' 取消，或按 'Enter' 強制推送。",
		ForcePushDisabled:                "你的分支與遠端分支分岔，你已禁用強制推送",
		CheckForUpdate:                   "檢查更新",
		CheckingForUpdates:               "正在檢查更新...",
		UpdateAvailableTitle:             "有可用的更新！",
		UpdateAvailable:                  "下載並安裝版本 {{.newVersion}}？",
		UpdateInProgressWaitingStatus:    "更新中",
		UpdateCompletedTitle:             "更新已完成！",
		UpdateCompleted:                  "更新已成功安裝。為了使其生效，請重新啟動 lazygit。",
		FailedToRetrieveLatestVersionErr: "無法取得版本資訊",
		OnLatestVersionErr:               "已更新至最新版本",
		MajorVersionErr:                  "新版本（{{.newVersion}}）不支援當前版本（{{.currentVersion}}）更改",
		CouldNotFindBinaryErr:            "找不到 {{.url}} 執行檔",
		UpdateFailedErr:                  "更新失敗：{{.errMessage}}",
		ConfirmQuitDuringUpdateTitle:     "正在更新中",
		ConfirmQuitDuringUpdate:          "正在進行更新，是否結束？",
		MergeToolTitle:                   "合併工具",
		MergeToolPrompt:                  "是否開啟 'git mergetool'？",
		IntroPopupMessage:                traditionalChineseIntroPopupMessage,
		DeprecatedEditConfigWarning:      traditionalChineseDeprecatedEditConfigWarning,
		GitconfigParseErr:                `Gogit 無法解析你的 gitconfig 檔案，因為存在未引用的 '\' 字符，刪除它們應該可以解決這個問題。`,
		EditFile:                         `編輯檔案`,
		OpenFile:                         `開啟檔案`,
		IgnoreFile:                       `添加到 .gitignore`,
		ExcludeFile:                      `添加到 .git/info/exclude`,
		RefreshFiles:                     `重新整理檔案`,
		Merge:                            `合併到當前檢出的分支`,
		ConfirmQuit:                      `是否結束？`,
		SwitchRepo:                       `切換到最近使用的版本庫`,
		AllBranchesLogGraph:              `顯示所有分支日誌`,
		UnsupportedGitService:            `不支援的 git 服務`,
		CreatePullRequest:                `建立拉取請求`,
		CopyPullRequestURL:               `複製拉取請求的 URL 到剪貼板`,
		NoBranchOnRemote:                 `這個分支在遠端不存在。需要先將其推送至遠端。`,
		Fetch:                            `擷取`,
		NoAutomaticGitFetchTitle:         `手動 git 擷取`,
		NoAutomaticGitFetchBody:          `lazygit 無法在私有庫使用 "git 擷取"；在檔案面板中使用 'f' 手動執行 "git 擷取"`,
		FileEnter:                        `選擇檔案中的單個程式碼塊/行，或展開/折疊目錄`,
		FileStagingRequirements:          `只能選擇跟踪檔案中的單個行`,
		StageSelectionTooltip:            `切換現有行的狀態 (已預存/未預存)`,
		DiscardSelection:                 `刪除變更 (git reset)`,
		ToggleRangeSelect:                `切換拖曳選擇`,
		ToggleSelectHunk:                 `切換選擇程式碼塊`,
		ToggleSelectionForPatch:          `向 (或從) 補丁中添加/刪除行`,
		EditHunk:                         `編輯程式碼塊`,
		ToggleStagingView:                `切換至另一個面板 (已預存/未預存更改)`,
		ReturnToFilesPanel:               `返回檔案面板`,
		FastForward:                      `從上游快進此分支`,
		FastForwarding:                   "的擷取和快進中",
		FoundConflictsTitle:              "自動合併失敗",
		ViewMergeRebaseOptions:           "查看合併/變基選項",
		NotMergingOrRebasing:             "你當前既不在變基也不在合併中",
		AlreadyRebasing:                  "無法在變基期間執行此操作",
		RecentRepos:                      "最近的版本庫",
		MergeOptionsTitle:                "合併選項",
		RebaseOptionsTitle:               "變基選項",
		CommitSummaryTitle:               "提交摘要",
		CommitDescriptionTitle:           "提交描述",
		CommitDescriptionSubTitle:        "按 tab 鍵聚焦",
		LocalBranchesTitle:               "本地分支",
		SearchTitle:                      "搜尋",
		TagsTitle:                        "標籤",
		MenuTitle:                        "功能表",
		RemotesTitle:                     "遠端",
		RemoteBranchesTitle:              "遠端分支",
		PatchBuildingTitle:               "主面板 (補丁生成)",
		InformationTitle:                 "資訊",
		SecondaryTitle:                   "次要",
		ReflogCommitsTitle:               "日誌",
		GlobalTitle:                      "全域快捷鍵",
		ConflictsResolved:                "所有合併衝突都已解決。是否繼續？",
		Continue:                         "確認",
		Keybindings:                      "鍵盤快捷鍵",
		RebasingTitle:                    "將 '{{.checkedOutBranch}}' 變基至 '{{.ref}}'",
		SimpleRebase:                     "簡單變基",
		InteractiveRebase:                "互動變基",
		InteractiveRebaseTooltip:         "開始一個互動變基，以中斷開始，這樣你可以在繼續之前更新TODO提交",
		ConfirmMerge:                     "是否將 '{{.selectedBranch}}' 合併至 '{{.checkedOutBranch}}' ？",
		FwdNoUpstream:                    "無法快進無上游分支",
		FwdNoLocalUpstream:               "無法快進尚未在本地註冊的遠端分支",
		FwdCommitsToPush:                 "無法快進帶有尚未推送的提交的分支",
		ErrorOccurred:                    "發生錯誤！請在此詢問錯誤：",
		NoRoom:                           "無足夠的空間",
		YouAreHere:                       "你在這",
		YouDied:                          "你死了！",
		RewordNotSupported:               "在互動變基期間改寫提交目前不支援",
		ChangingThisActionIsNotAllowed:   "不允許更改此類變基待辦事項",
		CherryPickCopy:                   "複製提交 (揀選)",
		PasteCommits:                     "貼上提交 (揀選)",
		SureCherryPick:                   "是否將複製的提交揀選到此分支？",
		CherryPick:                       "揀選 (Cherry-pick)",
		Donate:                           "贊助",
		AskQuestion:                      "諮詢",
		PrevLine:                         "選擇上一行",
		NextLine:                         "選擇下一行",
		PrevHunk:                         "選擇上一段",
		NextHunk:                         "選擇下一段",
		PrevConflict:                     "選擇上一個衝突",
		NextConflict:                     "選擇下一個衝突",
		SelectPrevHunk:                   "選擇上一段",
		SelectNextHunk:                   "選擇下一段",
		ScrollDown:                       "向下捲動",
		ScrollUp:                         "向上捲動",
		ScrollUpMainWindow:               "向上捲動主面板",
		ScrollDownMainWindow:             "向下捲動主面板",
		AmendCommitTitle:                 "修改提交",
		AmendCommitPrompt:                "是否使用預存檔案修改提交？",
		DropCommitTitle:                  "刪除提交",
		DropCommitPrompt:                 "是否刪除此提交？",
		PullingStatus:                    "拉取",
		PushingStatus:                    "推送",
		FetchingStatus:                   "擷取",
		SquashingStatus:                  "壓縮中",
		FixingStatus:                     "修復中",
		DeletingStatus:                   "刪除中",
		MovingStatus:                     "移動中",
		RebasingStatus:                   "變基中",
		MergingStatus:                    "合併中",
		LowercaseRebasingStatus:          "變基", // lowercase because it shows up in parentheses
		LowercaseMergingStatus:           "合併", // lowercase because it shows up in parentheses
		AmendingStatus:                   "修改中",
		CherryPickingStatus:              "揀選中",
		UndoingStatus:                    "復原中",
		RedoingStatus:                    "重做中",
		CheckingOutStatus:                "檢出中",
		CommittingStatus:                 "提交中",
		RevertingStatus:                  "還原中",
		CommitFiles:                      "提交檔案",
		SubCommitsDynamicTitle:           "提交 (共 %s項)",
		CommitFilesDynamicTitle:          "差異檔案 (共 %s項)",
		RemoteBranchesDynamicTitle:       "遠端分支 (共 %s項)",
		ViewItemFiles:                    "檢視所選項目的檔案",
		CommitFilesTitle:                 "提交檔案",
		CheckoutCommitFileTooltip:        "檢出檔案",
		DiscardFileChangesTitle:          "捨棄檔案更改",
		DiscardFileChangesPrompt:         "是否捨棄此提交？如果這個檔案是在此提交中創建的，它將被刪除",
		DisabledForGPG:                   "此功能不適用於 GPG 加密",
		CreateRepo:                       "未在 git 版本庫中。是否建立新版本庫？ (y/n): ",
		BareRepo:                         "你嘗試在裸版本庫中開啟 Lazygit，但 Lazygit 尚未支援裸版本庫。是否開啟最新版本庫？ (y/n) ",
		InitialBranch:                    "分支名稱？（留空使用 git 的預設值）：",
		NoRecentRepositories:             "必須在 git 版本庫中開啟 lazygit。沒有有效的最近版本庫。退出。",
		IncorrectNotARepository:          "無效 `notARepository` 輸入。輸入應為「prompt」、「create」、「skip」、或「quit」。",
		AutoStashTitle:                   "是否自動收藏？",
		AutoStashPrompt:                  "必須收藏並拾起變更才得以繼續操作。是否自動執行？（Enter/Esc）",
		StashPrefix:                      "自動收藏 ",
		Cancel:                           "取消",
		DiscardAllChanges:                "刪除所有變更",
		DiscardUnstagedChanges:           "刪除未預存變更",
		DiscardAllChangesToAllFiles:      "刪除工作目錄",
		DiscardAnyUnstagedChanges:        "刪除未預存變更",
		DiscardUntrackedFiles:            "刪除未追蹤檔案",
		DiscardStagedChanges:             "刪除已預存變更",
		HardReset:                        "強制重設",
		ViewResetOptions:                 "檢視重設選項",
		CreateFixupCommitTooltip:         "為此提交建立修復提交",
		SquashAboveCommits:               "壓縮上方所有「fixup」提交（自動壓縮）",
		SquashAboveCommitsTooltip:        "是否壓縮上方 {{.commit}} 所有「fixup」提交？",
		CreateFixupCommit:                "建立修復提交",
		ExecuteCustomCommand:             "執行自訂命令",
		CustomCommand:                    "自訂命令：",
		CommitChangesWithoutHook:         "沒有預提交 hook 就提交更改",
		SkipHookPrefixNotConfigured:      "你尚未配置略過 hook 的提交訊息前綴，請在設定中設置 `git.skipHookPrefix = 'WIP'`",
		ResetTo:                          `重設至`,
		PressEnterToReturn:               "按 Enter 返回到 lazygit",
		ViewStashOptions:                 "檢視收藏選項",
		StashAllChanges:                  "收藏所有變更",
		StashStagedChanges:               "收藏已預存變更",
		StashAllChangesKeepIndex:         "收藏所有變更並保留預存區",
		StashUnstagedChanges:             "收藏未預存變更",
		StashIncludeUntrackedChanges:     "收藏所有變更，包括未追蹤檔案",
		StashOptions:                     "收藏選項",
		NotARepository:                   "錯誤：必須在 git 版本庫中執行",
		Jump:                             "跳轉至面板",
		ScrollLeftRight:                  "左右捲動",
		ScrollLeft:                       "向左捲動",
		ScrollRight:                      "向右捲動",
		DiscardPatch:                     "捨棄補丁",
		DiscardPatchConfirm:              "你只能從單一提交或收藏項目建立一個補丁。是否捨棄當前補丁？",
		CantPatchWhileRebasingError:      "在合併或變基狀態下，你不能建立或運行補丁命令",
		ToggleAddToPatch:                 "切換檔案是否包含在補丁中",
		ToggleAllInPatch:                 "切換所有檔案是否包含在補丁中",
		UpdatingPatch:                    "正在更新補丁",
		ViewPatchOptions:                 "檢視自訂補丁選項",
		PatchOptionsTitle:                "補丁選項",
		NoPatchError:                     "尚未建立補丁。要開始建立補丁，請在提交檔案上使用空格或輸入以添加特定行",
		EnterCommitFile:                  "輸入檔案以將選定的行添加至補丁（或切換目錄折疊）",
		ExitCustomPatchBuilder:           `退出自訂補丁建立器`,
		EnterUpstream:                    `輸入上游為 '<remote> <branchname>'`,
		InvalidUpstream:                  "無效的上游。必須符合 '<remote> <branchname>' 的格式",
		ReturnToRemotesList:              `返回遠端列表`,
		NewRemote:                        `新增遠端`,
		NewRemoteName:                    `新遠端名稱：`,
		NewRemoteUrl:                     `新遠端 URL：`,
		EditRemoteName:                   `輸入更新 {{.remoteName}} 遠端名稱：`,
		EditRemoteUrl:                    `輸入更新 {{.remoteName}} 遠端 URL：`,
		RemoveRemote:                     `移除遠端`,
		RemoveRemotePrompt:               "你確定要移除遠端？",
		DeleteRemoteBranch:               "刪除遠端分支",
		DeleteRemoteBranchMessage:        "你確定要刪除遠端分支？",
		SetAsUpstreamTooltip:             "將此分支設為當前分支之上游",
		SetUpstream:                      "設定所選分支之上游",
		UnsetUpstream:                    "取消設定選定分支之上游",
		SetUpstreamTitle:                 "設定上游分支",
		SetUpstreamMessage:               "你確定要將 '{{. selected}}' 設為 '{{.checkedOut}}' 的上游分支？",
		EditRemoteTooltip:                "編輯遠端",
		TagCommit:                        "打標籤到提交",
		TagMenuTitle:                     "建立標籤",
		TagNameTitle:                     "標籤名稱",
		TagMessageTitle:                  "標籤訊息",
		AnnotatedTag:                     "附註標籤",
		LightweightTag:                   "輕量標籤",
		PushTagTitle:                     "推送標籤 '{{.tagName}}' 至遠端：",
		PushTag:                          "推送標籤",
		NewTag:                           "建立標籤",
		FetchRemoteTooltip:               "擷取遠端",
		FetchingRemoteStatus:             "正在擷取遠端",
		CheckoutCommit:                   "檢出提交",
		SureCheckoutThisCommit:           "你確定要檢出這個提交？",
		GitFlowOptions:                   "顯示 git-flow 選項",
		NotAGitFlowBranch:                "這似乎不是一個 git flow 分支",
		NewGitFlowBranchPrompt:           "{{.branchType}} 名稱：",
		IgnoreTracked:                    "忽略已追蹤檔案",
		IgnoreTrackedPrompt:              "你確定要忽略一個已追蹤的檔案？",
		ExcludeTracked:                   "排除已追蹤檔案",
		ViewResetToUpstreamOptions:       "檢視上游重設選項",
		NextScreenMode:                   "下一個螢幕模式（常規/半螢幕/全螢幕）",
		PrevScreenMode:                   "上一個螢幕模式",
		StartSearch:                      "搜尋",
		StartFilter:                      "搜尋",
		Panel:                            "面板",
		KeybindingsLegend:                "說明：`<c-b>` 表示 Ctrl＋B、`<a-b>` 表示 Alt＋B，`B`表示 Shift＋B",
		RenameBranch:                     "重新命名分支",
		BranchUpstreamOptionsTitle:       "上游分支設定",
		ViewBranchUpstreamOptionsTooltip: "檢視有關上游分支的設定（例如重設至上游）",
		UpstreamNotSetError:              "目標分支沒有上游分支（或其上游分支未儲存於本地）",
		ViewBranchUpstreamOptions:        "檢視上游設定",
		NewBranchNamePrompt:              "為分支輸入新名稱",
		RenameBranchWarning:              "此分支正在追蹤遠端分支。此操作僅會重新命名本地分支名稱，而不是遠端分支的名稱。是否繼續？",
		OpenKeybindingsMenu:              "開啟選單",
		ResetCherryPick:                  "重設選定的揀選 (複製) 提交",
		NextTab:                          "下一個索引標籤",
		PrevTab:                          "上一個索引標籤",
		CantUndoWhileRebasing:            "在變基時無法復原",
		CantRedoWhileRebasing:            "在變基時無法取消復原",
		MustStashWarning:                 "將補丁提取到索引中需要收藏並取消收藏你的變更。如果出現問題，你可以從收藏中訪問你的檔案。是否繼續？",
		MustStashTitle:                   "必須收藏",
		ConfirmationTitle:                "確認面板",
		PrevPage:                         "上一頁",
		NextPage:                         "下一頁",
		GotoTop:                          "捲動到頂部",
		GotoBottom:                       "捲動到底部",
		FilteringBy:                      "篩選方式",
		ResetInParentheses:               "（已重設）",
		OpenFilteringMenu:                "檢視篩選路徑選項",
		FilterBy:                         "篩選路徑",
		ExitFilterMode:                   "停止按路徑篩選",
		FilterPathOption:                 "輸入要依路徑篩選的路徑",
		EnterFileName:                    "輸入路徑：",
		FilteringMenuTitle:               "篩選",
		MustExitFilterModeTitle:          "命令不可用",
		MustExitFilterModePrompt:         "在按路徑篩選的模式下，該命令不可用。是否退出按路徑篩選的模式？",
		Diff:                             "差異",
		EnterRefToDiff:                   "輸入欲比較之 Ref",
		EnterRefName:                     "輸入 Ref：",
		ExitDiffMode:                     "退出差異模式",
		DiffingMenuTitle:                 "差異比較",
		SwapDiff:                         "反轉差異方向",
		ViewDiffingOptions:               "開啟差異比較選單",
		// the actual view is the extras view which I intend to give more tabs in future but for now we'll only mention the command log part
		OpenCommandLogMenu:                  "開啟命令記錄選單",
		ShowingGitDiff:                      "顯示輸出：",
		CommitDiff:                          "提交差異",
		CopyCommitShaToClipboard:            "複製提交 SHA 到剪貼簿",
		CommitSha:                           "提交 SHA",
		CommitURL:                           "提交 URL",
		CopyCommitMessageToClipboard:        "複製提交訊息到剪貼簿",
		CommitMessage:                       "提交訊息",
		CommitAuthor:                        "提交者",
		CopyCommitAttributeToClipboard:      "複製提交屬性",
		CopyBranchNameToClipboard:           "複製分支名稱到剪貼簿",
		CopyPathToClipboard:                 "複製檔案名稱到剪貼簿",
		CopySelectedTextToClipboard:         "複製所選文本至剪貼簿",
		CommitPrefixPatternError:            "commitPrefix 模式錯誤",
		NoFilesStagedTitle:                  "沒有檔案預存",
		NoFilesStagedPrompt:                 "你沒有預存任何檔案。提交所有檔案？",
		BranchNotFoundTitle:                 "找不到分支",
		BranchNotFoundPrompt:                "找不到分支。新分支名稱",
		BranchUnknown:                       "分支未知",
		DiscardChangeTitle:                  "取消預存行",
		DiscardChangePrompt:                 "是否刪除所選行（git reset）？此操作不可逆。\n將「gui.skipDiscardChangeWarning」設為 true 可禁用此警告。",
		CreateNewBranchFromCommit:           "從提交建立新分支",
		BuildingPatch:                       "正在建立補丁",
		ViewCommits:                         "檢視提交",
		MinGitVersionError:                  "請升級 git 至新於 2.20（即從 2018 年起）之版本。或於 https://github.com/jesseduffield/lazygit/issues 上回報問題使 lazygit 能支援更舊的 git 版本。",
		RunningCustomCommandStatus:          "正在執行自訂命令",
		SubmoduleStashAndReset:              "收藏未提交的子模組變更並更新",
		AndResetSubmodules:                  "以及重設子模組",
		EnterSubmoduleTooltip:               "進入子模組",
		CopySubmoduleNameToClipboard:        "複製子模組名稱到剪貼簿",
		RemoveSubmodule:                     "移除子模組",
		RemoveSubmodulePrompt:               "是否確定要刪除子模組 '%s' 以及它相應的目錄？此操作是不可逆的。",
		ResettingSubmoduleStatus:            "重設子模型中",
		NewSubmoduleName:                    "子模組名稱：",
		NewSubmoduleUrl:                     "新子模組 URL：",
		NewSubmodulePath:                    "新子模組路徑：",
		NewSubmodule:                        "新增子模組",
		AddingSubmoduleStatus:               "正在新增子模組",
		UpdateSubmoduleUrl:                  "更新子模組 '%s' 的 URL",
		UpdatingSubmoduleUrlStatus:          "正在更新 URL",
		EditSubmoduleUrl:                    "更新子模組 URL",
		InitializingSubmoduleStatus:         "正在初始化子模組",
		InitSubmoduleTooltip:                "初始化子模組",
		SubmoduleUpdateTooltip:              "更新子模組",
		UpdatingSubmoduleStatus:             "正在更新子模組",
		BulkInitSubmodules:                  "批量初始化子模組",
		BulkUpdateSubmodules:                "批量更新子模組",
		BulkDeinitSubmodules:                "批量解除子模組初始化",
		ViewBulkSubmoduleOptions:            "查看批量子模組選項",
		BulkSubmoduleOptions:                "批量子模組選項",
		RunningCommand:                      "正在執行命令",
		SubCommitsTitle:                     "子提交",
		SubmodulesTitle:                     "子模組",
		NavigationTitle:                     "移動",
		SuggestionsCheatsheetTitle:          "提示",
		SuggestionsTitle:                    "提示（按 %s 進入焦點）",
		ExtrasTitle:                         "命令記錄",
		PushingTagStatus:                    "正在推送標籤",
		PullRequestURLCopiedToClipboard:     "複製拉取請求 URL 至剪貼簿",
		CommitDiffCopiedToClipboard:         "已複製提交差異至剪貼簿",
		CommitURLCopiedToClipboard:          "已複製提交 URL 至剪貼簿",
		CommitMessageCopiedToClipboard:      "已複製提交訊息至剪貼簿",
		CommitAuthorCopiedToClipboard:       "已複製提交者至剪貼簿",
		PatchCopiedToClipboard:              "已複製補丁至剪貼簿",
		CopiedToClipboard:                   "已複製至剪貼簿",
		ErrCannotEditDirectory:              "無法編輯目錄：你只能編輯單獨的檔案",
		ErrStageDirWithInlineMergeConflicts: "不能預存/取消預存包含具備內嵌合併衝突的檔案的目錄。請先解決合併衝突",
		ErrRepositoryMovedOrDeleted:         "找不到版本庫。可能已被移動或刪除",
		CommandLog:                          "命令記錄",
		ToggleShowCommandLog:                "切換顯示/隱藏命令記錄",
		FocusCommandLog:                     "聚焦命令記錄",
		CommandLogHeader:                    " '%s' 隱藏/聚焦此面板\n",
		RandomTip:                           "隨機提示",
		SelectParentCommitForMerge:          "選擇合併的父提交",
		ToggleWhitespaceInDiffView:          "切換是否在差異檢視中顯示空格變更",
		IgnoreWhitespaceDiffViewSubTitle:    "（忽略空格）",
		IgnoreWhitespaceNotSupportedHere:    "在此檢視中不支援忽略空格",
		IncreaseContextInDiffView:           "增加差異檢視中顯示變更周圍上下文的大小",
		DecreaseContextInDiffView:           "減小差異檢視中顯示變更周圍上下文的大小",
		CreatePullRequestOptions:            "建立拉取請求選項",
		DefaultBranch:                       "預設分支",
		SelectBranch:                        "選擇分支",
		SelectConfigFile:                    "選擇設定檔",
		NoConfigFileFoundErr:                "找不到設定檔",
		LoadingFileSuggestions:              "正在加載檔案建議",
		LoadingCommits:                      "正在加載提交",
		MustSpecifyOriginError:              "如果指定分支，必須指定遠端",
		GitOutput:                           "git 輸出：",
		GitCommandFailed:                    "git 命令失敗。請查看命令記錄以獲取詳細資訊（按 %s 開啟）",
		AbortTitle:                          "中止%s",
		AbortPrompt:                         "是否確定要中止當前的%s？",
		OpenLogMenu:                         "開啟記錄選單",
		LogMenuTitle:                        "提交記錄選項",
		ToggleShowGitGraphAll:               "切換顯示整個 git 圖表（將 `--all` 標誌傳遞給 `git log`）",
		ShowGitGraph:                        "顯示 git 圖表",
		SortCommits:                         "提交排序順序",
		CantChangeContextSizeError:          "在製作補丁期間無法更改上下文大小，因為當發布功能時我們太懒了以至於沒有支援它。如果你真的需要它，請告訴我們！",
		OpenCommitInBrowser:                 "在瀏覽器中開啟提交",
		ViewBisectOptions:                   "查看二分選項",
		ConfirmRevertCommit:                 "是否還原 {{.selectedCommit}} ？",
		RewordInEditorTitle:                 "在編輯器中改寫",
		RewordInEditorPrompt:                "是否在編輯器中改寫此提交？",
		HardResetAutostashPrompt:            "是否強制重設為 '%s' ？如果需要會進行自動存儲。",
		CheckoutPrompt:                      "是否檢出 '%s' ？",
		UpstreamGone:                        "(上游已經不存在)",
		NukeDescription:                     "如果你想讓所有工作樹上的變更消失，這就是要做的方式。如果有未提交的子模組變更，它將把這些變更藏在子模組中。",
		DiscardStagedChangesDescription:     "這將創建一個新的存儲條目，其中只包含預存檔案，然後如果存儲條目不需要，將其刪除，因此工作樹僅保留未預存的變更。",
		EmptyOutput:                         "<空輸出>",
		Patch:                               "補丁",
		CustomPatch:                         "自定義補丁",
		CommitsCopied:                       "提交已複製", // lowercase because it's used in a sentence
		CommitCopied:                        "提交已複製", // lowercase because it's used in a sentence
		ResetPatch:                          "重設補丁",
		ApplyPatch:                          "套用補丁",
		ApplyPatchInReverse:                 "反向套用補丁",
		RemovePatchFromOriginalCommit:       "從原始提交中刪除補丁（%s）",
		MovePatchOutIntoIndex:               "將補丁移到預存區",
		MovePatchIntoNewCommit:              "將補丁移到新的提交",
		MovePatchToSelectedCommit:           "將補丁移到選定的提交（%s）",
		CopyPatchToClipboard:                "將補丁複製到剪貼簿",
		NoMatchesFor:                        "沒有找到符合 '%s' %s 的結果",
		ExitSearchMode:                      "%s：退出搜尋模式",
		MatchesFor:                          "符合 '%s' 的結果（%d/%d）%s", // lowercase because it's after other text
		SearchKeybindings:                   "%s：下一個結果，%s：上一個結果，%s：退出搜尋模式",
		SearchPrefix:                        "搜尋：",
		FilterPrefix:                        "篩選：",
		WorktreesTitle:                      "工作目錄",
		WorktreeTitle:                       "工作目錄",
		SwitchToWorktree:                    "切換至工作目錄面板",
		AlreadyCheckedOutByWorktree:         "此分支已被檢出到 {{.worktreeName}} 是否切換到此工作目錄？",
		BranchCheckedOutByWorktree:          "分支 {{.branchName}} 已被 {{.worktreeName}} 檢出",
		DetachWorktreeTooltip:               "此將在工作目錄中執行 `git checkout --detach` 以解開分支與它的連結，但工作目錄本身將不被更動",
		Switching:                           "切換中",
		RemoveWorktree:                      "刪除工作目錄",
		RemoveWorktreeTitle:                 "刪除工作目錄",
		RemoveWorktreePrompt:                "是否刪除 {{.worktreeName}} 工作目錄？",
		ForceRemoveWorktreePrompt:           "'{{.worktreeName}}' 包括已更動或未追蹤的檔案。是否繼續刪除工作目錄？",
		RemovingWorktree:                    "正在刪除工作目錄",
		DetachWorktree:                      "解開工作目錄連結",
		DetachingWorktree:                   "正在解除工作目錄連結",
		AddingWorktree:                      "正在建立工作目錄",
		CantDeleteCurrentWorktree:           "無法刪除當前工作目錄！",
		AlreadyInWorktree:                   "已經在目標工作目錄內",
		CantDeleteMainWorktree:              "無法刪除主要工作目錄！",
		NoWorktreesThisRepo:                 "無工作目錄",
		MissingWorktree:                     "（失蹤）",
		MainWorktree:                        "（主要）",
		NewWorktreePath:                     "工作目錄路徑",
		NewWorktreeBase:                     "工作目錄來源",
		BranchNameCannotBeBlank:             "分支名稱不能為空",
		NewBranchName:                       "分支名稱",
		NewBranchNameLeaveBlank:             "分支名稱（留空將檢出 {{.default}}）",
		ViewWorktreeOptions:                 "檢視工作目錄選項",
		CreateWorktreeFrom:                  "從 {{.ref}} 建立工作目錄",
		CreateWorktreeFromDetached:          "從 {{.ref}} 建立工作目錄（未連結）",
		LcWorktree:                          "工作目錄",
		ChangingDirectoryTo:                 "切換至 {{.path}}",
		Name:                                "名稱",
		Branch:                              "分支",
		Path:                                "路徑",
		MarkedBaseCommitStatus:              "為了變基已標注基準提交",
		MarkAsBaseCommit:                    "為了變基已標注提交為基準提交",
		MarkAsBaseCommitTooltip:             "請為了下一次變基選擇一項基準提交；此將執行 `git rebase --onto`。",
		MarkedCommitMarker:                  "↑↑↑ 將由此變基 ↑↑↑",
		PleaseGoToURL:                       "請開啟 URL：{{.url}}",
		DisabledMenuItemPrefix:              "已停用：",
		NoCopiedCommits:                     "未複製提交",
		Actions: Actions{
			// TODO: combine this with the original keybinding descriptions (those are all in lowercase atm)
			CheckoutCommit:                    "檢出提交",
			CheckoutTag:                       "檢出標籤",
			CheckoutBranch:                    "檢出分支",
			ForceCheckoutBranch:               "強制檢出分支",
			DeleteBranch:                      "刪除分支",
			Merge:                             "合併",
			RebaseBranch:                      "變基分支",
			RenameBranch:                      "重新命名分支",
			CreateBranch:                      "建立分支",
			CherryPick:                        "（Cherry-pick）複製提交",
			CheckoutFile:                      "檢出檔案",
			DiscardOldFileChange:              "放棄舊檔案更改",
			SquashCommitDown:                  "下列次方執行 Squash",
			FixupCommit:                       "修復提交",
			RewordCommit:                      "改寫提交",
			DropCommit:                        "捨棄提交",
			EditCommit:                        "編輯提交",
			AmendCommit:                       "修改提交",
			ResetCommitAuthor:                 "重設提交作者",
			SetCommitAuthor:                   "設置提交作者",
			RevertCommit:                      "還原提交",
			CreateFixupCommit:                 "建立修改提交",
			SquashAllAboveFixupCommits:        "Squash 所有上面的修改提交",
			CreateLightweightTag:              "建立輕量標籤",
			CreateAnnotatedTag:                "建立附註標籤",
			CopyCommitMessageToClipboard:      "將提交訊息複製到剪貼簿",
			CopyCommitDiffToClipboard:         "將提交差異複製到剪貼簿",
			CopyCommitSHAToClipboard:          "將提交 SHA 複製到剪貼簿",
			CopyCommitURLToClipboard:          "將提交 URL 複製到剪貼簿",
			CopyCommitAuthorToClipboard:       "將提交作者複製到剪貼簿",
			CopyCommitAttributeToClipboard:    "複製到剪貼簿",
			CopyPatchToClipboard:              "將補丁複製到剪貼簿",
			MoveCommitUp:                      "上移提交",
			MoveCommitDown:                    "下移提交",
			CustomCommand:                     "自定義命令",
			DiscardAllChangesInDirectory:      "捨棄目錄中的所有更改",
			DiscardUnstagedChangesInDirectory: "捨棄目錄中未預存的更改",
			DiscardAllChangesInFile:           "捨棄檔案中的所有更改",
			DiscardAllUnstagedChangesInFile:   "捨棄檔案中未預存的所有更改",
			StageFile:                         "預存檔案",
			StageResolvedFiles:                "預存已解決合併衝突的檔案",
			UnstageFile:                       "取消預存檔案",
			UnstageAllFiles:                   "取消預存所有檔案",
			StageAllFiles:                     "預存所有檔案",
			IgnoreExcludeFile:                 "忽略或排除檔案",
			IgnoreFileErr:                     "無法忽略 .gitignore 檔案",
			ExcludeFile:                       "排除檔案",
			ExcludeFileErr:                    "無法排除 .git/info/exclude 檔案",
			ExcludeGitIgnoreErr:               "無法排除 .gitignore 檔案",
			Commit:                            "提交",
			EditFile:                          "編輯檔案",
			Push:                              "推送",
			Pull:                              "拉取",
			OpenFile:                          "開啟檔案",
			StashAllChanges:                   "收藏所有更改",
			StashAllChangesKeepIndex:          "收藏所有更改並保留索引",
			StashStagedChanges:                "收藏已預存的更改",
			StashUnstagedChanges:              "收藏未預存的更改",
			StashIncludeUntrackedChanges:      "收藏所有更改，包括未追蹤的檔案",
			GitFlowFinish:                     "`git flow` 完成",
			GitFlowStart:                      "`git flow` 開始",
			CopyToClipboard:                   "複製到剪貼簿",
			CopySelectedTextToClipboard:       "複製所選文本到剪貼簿",
			RemovePatchFromCommit:             "從提交中刪除補丁",
			MovePatchToSelectedCommit:         "將補丁移動到所選提交",
			MovePatchIntoIndex:                "將補丁移動到索引中",
			MovePatchIntoNewCommit:            "將補丁移動到新提交中",
			DeleteRemoteBranch:                "刪除遠端分支",
			SetBranchUpstream:                 "設置分支上游",
			AddRemote:                         "添加遠端",
			RemoveRemote:                      "移除遠端",
			UpdateRemote:                      "更新遠端",
			ApplyPatch:                        "套用補丁",
			Stash:                             "收藏 (Stash)",
			RenameStash:                       "重命名暫存",
			RemoveSubmodule:                   "移除子模塊",
			ResetSubmodule:                    "重設子模塊",
			AddSubmodule:                      "添加子模塊",
			UpdateSubmoduleUrl:                "更新子模塊 URL",
			InitialiseSubmodule:               "初始化子模塊",
			BulkInitialiseSubmodules:          "批量初始化子模塊",
			BulkUpdateSubmodules:              "批量更新子模塊",
			BulkDeinitialiseSubmodules:        "批量取消初始化子模塊",
			UpdateSubmodule:                   "更新子模塊",
			PushTag:                           "推送標籤",
			NukeWorkingTree:                   "清空工作樹",
			DiscardUnstagedFileChanges:        "放棄未預存的檔案更改",
			RemoveUntrackedFiles:              "移除未追蹤的檔案",
			RemoveStagedFiles:                 "移除已預存的檔案",
			SoftReset:                         "軟重設",
			MixedReset:                        "混合重設",
			HardReset:                         "強制重設",
			FastForwardBranch:                 "快進分支",
			Undo:                              "復原",
			Redo:                              "重做",
			CopyPullRequestURL:                "複製拉取請求的 URL",
			OpenMergeTool:                     "開啟合併工具",
			OpenCommitInBrowser:               "在瀏覽器中開啟提交",
			OpenPullRequest:                   "在瀏覽器中開啟拉取請求",
			StartBisect:                       "開始二分查找",
			ResetBisect:                       "重設二分查找",
			BisectSkip:                        "二分查找跳過",
			BisectMark:                        "二分查找標記",
		},
		Bisect: Bisect{
			Mark:                        "將 %s 標記為 %s",
			MarkStart:                   "將 %s 標記為 %s（開始二分查找）",
			SkipCurrent:                 "跳過 %s",
			ResetTitle:                  "重設 `git bisect`",
			ResetPrompt:                 "是否重設 `git bisect`？",
			ResetOption:                 "重設二分查找",
			BisectMenuTitle:             "二分查找",
			CompleteTitle:               "二分查找完成",
			CompletePrompt:              "二分查找完成！以下提交引入了更改：\n\n%s\n\n是否重設 `git bisect` ？",
			CompletePromptIndeterminate: "二分查找完成！有一些提交被跳過，因此以下任何提交皆可能引進更改：\n\n%s\n\n是否重設 `git bisect`？",
			Bisecting:                   "二分查找中",
		},
	}
}
