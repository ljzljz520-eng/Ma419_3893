package service

import "heritage/model"

func WorkflowReady(h model.Homepage) bool {
	return h.Featured.ID != "" && len(h.Articles) > 0 && len(h.Collections) > 0
}
