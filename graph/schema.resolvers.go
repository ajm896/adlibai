package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"fmt"

	"github.com/ajm896/adlibai/db"
	"github.com/ajm896/adlibai/graph/model"
)

// CreateJournalEntry is the resolver for the createJournalEntry field.
func (r *mutationResolver) CreateJournalEntry(ctx context.Context, userID string, content string) (*model.JournalEntry, error) {
	panic(fmt.Errorf("not implemented: CreateJournalEntry - createJournalEntry"))
}

// UpdateJournalEntry is the resolver for the updateJournalEntry field.
func (r *mutationResolver) UpdateJournalEntry(ctx context.Context, id string, content string) (*model.JournalEntry, error) {
	panic(fmt.Errorf("not implemented: UpdateJournalEntry - updateJournalEntry"))
}

// DeleteJournalEntry is the resolver for the deleteJournalEntry field.
func (r *mutationResolver) DeleteJournalEntry(ctx context.Context, id string) (*bool, error) {
	panic(fmt.Errorf("not implemented: DeleteJournalEntry - deleteJournalEntry"))
}

// SetAIMode is the resolver for the setAIMode field.
func (r *mutationResolver) SetAIMode(ctx context.Context, userID string, mode model.AIMode) (*model.UserSettings, error) {
	panic(fmt.Errorf("not implemented: SetAIMode - setAIMode"))
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, username string, email string) (*model.User, error) {
	var user db.User
	user.Username = username
	user.Email = email
	r.DB.WithContext(ctx).Create(&user)
	return user.ToQLUser(), nil
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Me - me"))
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	var user db.User
	r.DB.WithContext(ctx).Find(&user, id)
	return user.ToQLUser(), nil
}

// GetJournalEntries is the resolver for the getJournalEntries field.
func (r *queryResolver) GetJournalEntries(ctx context.Context, userID string) ([]*model.JournalEntry, error) {
	panic(fmt.Errorf("not implemented: GetJournalEntries - getJournalEntries"))
}

// GetStreak is the resolver for the getStreak field.
func (r *queryResolver) GetStreak(ctx context.Context, userID string) (*model.Streak, error) {
	panic(fmt.Errorf("not implemented: GetStreak - getStreak"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
