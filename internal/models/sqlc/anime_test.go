package models

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAddAnime(t *testing.T) {
	anime := AddAnimeParams{
		AnimeID:     100001,
		GenreID:     100001,
		StudioID:    100001,
		Title:       "怪兽8号",
		Evaluate:    "在日本这一“怪兽大国”，人们的日常生活遭受着怪兽的无情破坏。主人公日比野卡夫卡，小时候住的城市遭到破坏，他与青梅竹马亚白米娜做出了“将怪兽全部消灭”的约定。然而，32岁的卡夫卡梦想破灭，就职于一家怪兽尸体解体公司“怪兽清洁公司”。相比米娜作为日本防卫队队员而日渐活跃，卡夫卡每天都过着忧郁的生活。过来打工的市川莱诺告诉他，防卫队的年龄限制将会提高，并劝他去参加入队考试。于是重新下定决心的卡夫卡，却突然受到神秘生物的侵蚀，身体开始怪兽化，被称为“怪兽8号”。",
		ReleaseDate: sql.NullTime{Time: time.Now(), Valid: true},
		AnimeStatus: NullAnimationsAnimeStatus{AnimationsAnimeStatus: AnimationsAnimeStatus("completed"), Valid: true},
		Rating:      sql.NullFloat64{Float64: 9.2, Valid: true},
	}
	err := testQueries.AddAnime(context.Background(), anime)

	require.NoError(t, err)

}
