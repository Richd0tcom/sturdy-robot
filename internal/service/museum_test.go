package service

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"

	db "github.com/Richd0tcom/sturdy-robot/internal/db/sqlc"
)


func SeedMuseumDatabase(ctx context.Context, q *db.Queries) error {
	// Seed random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create art categories
	artCategories := make([]db.ArtCategory, 0)
	for i := 1; i <= 5; i++ {
		category, err := q.CreateArtCategory(ctx, db.CreateArtCategoryParams{
			Name:        fmt.Sprintf("Art Category %d", i),
			Description: pgtype.Text{String: fmt.Sprintf("Description for Art Category %d", i), Valid: true},
		})
		if err != nil {
			return fmt.Errorf("error creating art category: %v", err)
		}
		artCategories = append(artCategories, category)
	}

	// Create staff roles
	staffRoles := make([]db.StaffRole, 0)
	roleTitles := []string{"Curator", "Night Watchman", "Admin", "Janitor"}
	for _, title := range roleTitles {
		role, err := q.CreateStaffRole(ctx, db.CreateStaffRoleParams{
			Title:       title,
			Description: pgtype.Text{String: fmt.Sprintf("Description for %s", title), Valid: true},
		})
		if err != nil {
			return fmt.Errorf("error creating staff role: %v", err)
		}
		staffRoles = append(staffRoles, role)
	}

	// Create artists
	artists := make([]db.Artist, 0)
	for i := 1; i <= 5; i++ {
		artist, err := q.CreateArtist(ctx, db.CreateArtistParams{
			Name:        fmt.Sprintf("Artist %d", i),
			Biography:   pgtype.Text{String: fmt.Sprintf("Biography for Artist %d", i), Valid: true},
			BirthDate:   pgtype.Date{Time: time.Now().AddDate(-50, 0, 0), Valid: true},
			DeathDate:   pgtype.Date{Time: time.Now().AddDate(-10, 0, 0), Valid: true},
			Nationality:  pgtype.Text{String: "Unknown", Valid: true},
		})
		if err != nil {
			return fmt.Errorf("error creating artist: %v", err)
		}
		artists = append(artists, artist)
	}

	// Create artworks
	artworks := make([]db.Artwork, 0)
	for i := 1; i <= 10; i++ {
		artwork, err := q.CreateArtwork(ctx, db.CreateArtworkParams{
			Title:            fmt.Sprintf("Artwork %d", i),
			ArtistID:         pgtype.UUID{Bytes: artists[rand.Intn(len(artists))].ID.Bytes, Valid: true},
			CategoryID:       pgtype.UUID{Bytes: artCategories[rand.Intn(len(artCategories))].ID.Bytes, Valid: true},
			YearCreated:      pgtype.Int4{Int32: int32(rand.Intn(2023-1900) + 1900), Valid: true, },
			Medium:           pgtype.Text{String:"Oil on Canvas", Valid: true},
			Dimensions:       pgtype.Text{String:"100x100 cm", Valid: true},
			Description:      pgtype.Text{String: fmt.Sprintf("Description for Artwork %d", i), Valid: true},
			AcquisitionDate:  pgtype.Date{Time: time.Now().AddDate(-1, 0, 0), Valid: true},
			ConditionStatus:  pgtype.Text{String:"Good", Valid: true},
			LocationInMuseum: pgtype.Text{String:"Gallery A", Valid: true},
			ImageUrl:         pgtype.Text{String:fmt.Sprintf("https://example.com/artwork-%d.jpg", i), Valid: true},
		})
		if err != nil {
			return fmt.Errorf("error creating artwork: %v", err)
		}
		artworks = append(artworks, artwork)
	}

	// Create staff members
	staffMembers := make([]db.Staff, 0)
	for i := 1; i <= 10; i++ {
		staff, err := q.CreateStaff(ctx, db.CreateStaffParams{
			FirstName: fmt.Sprintf("First%d", i),
			LastName:  fmt.Sprintf("Last%d", i),
			RoleID:    pgtype.UUID{Bytes: staffRoles[rand.Intn(len(staffRoles))].ID.Bytes, Valid: true},
			Email:     fmt.Sprintf("staff%d@example.com", i),
			Phone:     pgtype.Text{String: fmt.Sprintf("+123456789%d", i), Valid: true},
			HireDate:  pgtype.Date{Time: time.Now().AddDate(-1, 0, 0), Valid: true},
			Status:    pgtype.Text{String:"Active", Valid: true},
		})
		if err != nil {
			return fmt.Errorf("error creating staff: %v", err)
		}
		staffMembers = append(staffMembers, staff)
	}
	

	// Create shifts
	for _, staff := range staffMembers {
		for j := 1; j <= 5; j++ {
			_, err := q.CreateShift(ctx, db.CreateShiftParams{
				StaffID:   pgtype.UUID{Bytes: staff.ID.Bytes, Valid: true},
				ShiftDate: pgtype.Date{Time: time.Now().AddDate(0, 0, j), Valid: true},
				StartTime: pgtype.Time{Microseconds: time.Now().UnixMicro(), Valid: true},
				EndTime:   pgtype.Time{Microseconds: time.Now().Add(time.Hour * 3).UnixMicro(), Valid: true},
				Status:    pgtype.Text{String:"Scheduled", Valid: true},
				Notes:     pgtype.Text{String: fmt.Sprintf("Notes for shift %d", j), Valid: true},
			})
			if err != nil {
				return fmt.Errorf("error creating shift: %v", err)
			}
		}
	}

	return nil
}

func TestMuseumDatabaseSeed(t *testing.T) {
	ctx := context.Background()
	err := SeedMuseumDatabase(ctx, testQueries)
	require.NoError(t, err)
}