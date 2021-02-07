package tags

import (
	"bridgecrewio/yor/common/git_service"
	"fmt"
	"reflect"
)

type GitLastModifiedAtTag struct {
	Tag
}

func (t *GitLastModifiedAtTag) Init() ITag {
	t.Key = "git_last_modified_at"

	return t
}

func (t *GitLastModifiedAtTag) CalculateValue(data interface{}) error {
	gitBlame, ok := data.(*git_service.GitBlame)
	if !ok {
		return fmt.Errorf("failed to convert data to *GitBlame, which is required to calculte tag value. Type of data: %s", reflect.TypeOf(data))
	}

	t.Value = GetLatestCommit(gitBlame).Date.String()
	return nil
}