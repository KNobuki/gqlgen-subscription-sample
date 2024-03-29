package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"
	"database/sql"

	"gqlgen-subscription-sample/graph/model"

	_ "github.com/go-sql-driver/mysql"
)

// CreateSmartMat is the resolver for the createSmartMat field.
func (r *mutationResolver) CreateSmartMat(ctx context.Context, currentWeight float64) (*model.SmartMat, error) {
	dbDriver := "mysql"
	dsn := "root:root@tcp(db:3306)/smart_shopping"
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		return &model.SmartMat{}, err
	}
	defer db.Close()
	result, err := db.Exec("INSERT INTO smart_mats (current_weight) VALUES (?)", currentWeight)
	if err != nil {
		return &model.SmartMat{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return &model.SmartMat{}, err
	}
	return &model.SmartMat{
		ID:            id,
		CurrentWeight: currentWeight,
	}, err
}

// UpdateSmartMatWeight is the resolver for the updateSmartMatWeight field.
func (r *mutationResolver) UpdateSmartMatWeight(ctx context.Context, id int64, currentWeight float64) (*model.SmartMat, error) {
	// db更新処理
	dbDriver := "mysql"
	dsn := "root:root@tcp(db:3306)/smart_shopping"
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		return &model.SmartMat{}, err
	}
	defer db.Close()
	_, err = db.Exec("UPDATE smart_mats SET current_weight = ? WHERE id = ?", currentWeight, id)
	if err != nil {
		return &model.SmartMat{}, err
	}
	if err != nil {
		return &model.SmartMat{}, err
	}

	// サブスクリプション更新処理
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	// 対応するMatIDのチャネルにPublish
	for _, ch := range r.ChannelsByMatID[id] {
		ch <- &model.SmartMat{
			ID:            id,
			CurrentWeight: currentWeight,
		}
	}

	return &model.SmartMat{
		ID:            id,
		CurrentWeight: currentWeight,
	}, err
}

// SmartMats is the resolver for the smartMats field.
func (r *queryResolver) SmartMats(ctx context.Context) ([]*model.SmartMat, error) {
	dbDriver := "mysql"
	dsn := "root:root@tcp(db:3306)/smart_shopping"
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		return []*model.SmartMat{}, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM smart_mats")
	defer rows.Close()
	if err != nil {
		return []*model.SmartMat{}, err
	}
	smartMats := make([]*model.SmartMat, 0)
	for rows.Next() {
		sm := &model.SmartMat{}
		if err = rows.Scan(&sm.ID, &sm.CurrentWeight); err != nil {
			return []*model.SmartMat{}, err
		}
		smartMats = append(smartMats, sm)
	}
	return smartMats, nil
}

// SmartMatWeightUpdated is the resolver for the smartMatWeightUpdated field.
func (r *subscriptionResolver) SmartMatWeightUpdated(ctx context.Context, id int64) (<-chan *model.SmartMat, error) {
	// Mutex で ChannelsByMatID の操作を排他制御
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	// マットIDに対応するチャネルを追加
	ch := make(chan *model.SmartMat, 1)
	r.ChannelsByMatID[id] = append(r.ChannelsByMatID[id], ch)

	// コネクション終了時にチャネルを削除
	go func() {
		<-ctx.Done()
		r.Mutex.Lock()
		defer r.Mutex.Unlock()
		for i, c := range r.ChannelsByMatID[id] {
			if c == ch {
				r.ChannelsByMatID[id] = append(r.ChannelsByMatID[id][:i], r.ChannelsByMatID[id][i+1:]...)
				break
			}
		}
	}()

	return ch, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
