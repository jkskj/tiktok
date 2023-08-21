package service

import (
	"errors"

	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func (s *InteractionService) Dislike(req *interaction.FavoriteActionRequest, userId int64) error {
	exist, err := cache.IsVideoLikeExist(s.ctx, req.VideoId, userId)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("you did not like this video")
	}

	if err := cache.ReduceVideoLikeCount(s.ctx, req.VideoId, userId); err != nil {
		return err
	}

	// write into mysql
	exist, err = db.IsFavoriteExist(s.ctx, userId, req.VideoId)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("you did not like this video")
	}

	if err := db.UpdateFavoriteStatus(s.ctx, userId, req.VideoId, 0); err != nil {
		return err
	}
	return nil
}
