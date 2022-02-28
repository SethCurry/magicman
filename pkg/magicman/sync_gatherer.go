package magicman

import (
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/SethCurry/magicman/pkg/ent"
	"github.com/SethCurry/magicman/pkg/ent/card"
	"github.com/SethCurry/magicman/pkg/ent/cardtype"
	"github.com/SethCurry/magicman/pkg/ent/color"
	"github.com/SethCurry/magicman/pkg/ent/set"
	"github.com/SethCurry/magicman/pkg/ent/subtype"
	"github.com/SethCurry/magicman/pkg/ent/supertype"
	"github.com/SethCurry/magicman/pkg/mtg"
	"go.uber.org/zap"
)

func CacheImages(ctx context.Context, db *ent.Client, intoDirectory string, logger *zap.Logger) error {
	for {
		cardToCache, err := db.Card.Query().Where(card.And(card.CachedImagePathIsNil(), card.ImageURLNEQ(""))).First(ctx)
		if ent.IsNotFound(err) {
			break
		}

		logger.Info("caching image for card",
			zap.String("name", cardToCache.Name),
			zap.String("multiverse_id", cardToCache.MultiverseID),
		)
		imageURL := cardToCache.ImageURL

		resp, err := http.Get(imageURL)
		if err != nil {
			return err
		}

		imgPath := filepath.Join(intoDirectory, cardToCache.MultiverseID+".png")

		fd, err := os.Create(imgPath)
		if err != nil {
			return err
		}

		_, err = io.Copy(fd, resp.Body)
		if err != nil {
			return err
		}

		resp.Body.Close()
		fd.Close()

		_, err = db.Card.UpdateOne(cardToCache).SetCachedImagePath(imgPath).Save(ctx)
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 2)
	}
	return nil
}

// TODO add rate limiter
// TODO add logging
func SyncGatherer(ctx context.Context, db *ent.Client, mtgClient *mtg.Client, logger *zap.Logger) error {
	page := 1

	for _, v := range mtg.Colors {
		_, err := db.Color.Query().Where(color.NameEQ(v.Full)).Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				_, err = db.Color.Create().SetName(v.Full).SetCode(v.Short).Save(ctx)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	sets, err := mtgClient.ListSets()
	if err != nil {
		return err
	}

	for _, s := range sets {
		_, err := db.Set.Query().Where(set.NameEQ(s.Name)).Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				_, err := db.Set.Create().
					SetCode(s.Code).
					SetName(s.Name).
					SetType(s.Type).
					SetBorder(s.Border).
					SetReleaseDate(s.ReleaseDate).
					Save(ctx)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	mtgTypes, err := mtgClient.ListTypes()
	if err != nil {
		return err
	}

	for _, t := range mtgTypes {
		_, err := db.CardType.Query().Where(cardtype.NameEQ(t)).Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				_, err = db.CardType.Create().SetName(t).Save(ctx)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	subTypes, err := mtgClient.ListSubTypes()
	if err != nil {
		return err
	}

	for _, t := range subTypes {
		_, err = db.SubType.Query().Where(subtype.NameEQ(t)).Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				_, err = db.SubType.Create().SetName(t).Save(ctx)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	superTypes, err := mtgClient.ListSuperTypes()
	if err != nil {
		return err
	}

	for _, t := range superTypes {
		_, err = db.SuperType.Query().Where(supertype.NameEQ(t)).Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				_, err = db.SuperType.Create().SetName(t).Save(ctx)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	for {
		cards, err := mtgClient.SearchCards(mtg.WithPage(page))
		if err != nil {
			return err
		}

		for _, v := range cards {
			_, err = db.Card.Query().Where(card.MultiverseIDEQ(v.MultiverseID)).Only(ctx)
			if err != nil {
				if ent.IsNotFound(err) {
					logger.Debug("creating new card",
						zap.String("name", v.Name),
					)
					// TODO set the rest of the fields
					newCardQuery := db.Card.Create().
						SetGathererID(v.ID).
						SetMultiverseID(v.MultiverseID).
						SetName(v.Name).
						SetType(v.Type).
						SetText(v.Text).
						SetManaCost(v.ManaCost).
						SetArtist(v.Artist).
						SetPower(v.Power).
						SetCmc(int(v.CMC)).
						SetToughness(v.Toughness).
						SetImageURL(v.ImageURL).
						SetRarity(v.Rarity).
						SetOriginalText(v.OriginalText).
						SetOriginalType(v.OriginalType)

					foundSet, err := db.Set.Query().Where(set.CodeEQ(v.Set)).Only(ctx)
					if err != nil {
						logger.Error("failed to find set",
							zap.String("set_code", v.Set),
						)
					} else {
						newCardQuery.SetSet(foundSet)
					}

					if len(v.Types) > 0 {
						cardTypes, err := db.CardType.Query().Where(cardtype.NameIn(v.Types...)).All(ctx)
						if err != nil {
							return err
						}
						newCardQuery = newCardQuery.AddTypes(cardTypes...)
					}

					if len(v.Supertypes) > 0 {
						cardSuperTypes, err := db.SuperType.Query().Where(supertype.NameIn(v.Supertypes...)).All(ctx)
						if err != nil {
							return err
						}
						newCardQuery = newCardQuery.AddSupertypes(cardSuperTypes...)
					}

					if len(v.Subtypes) > 0 {
						cardSubTypes, err := db.SubType.Query().Where(subtype.NameIn(v.Subtypes...)).All(ctx)
						if err != nil {
							return err
						}
						newCardQuery = newCardQuery.AddSubtypes(cardSubTypes...)
					}

					if len(v.ColorIdentity) > 0 {
						colors := make([]string, len(v.ColorIdentity))
						for i, c := range v.ColorIdentity {
							colors[i] = c.Full
						}

						cardColors, err := db.Color.Query().Where(color.NameIn(colors...)).All(ctx)
						if err != nil {
							return err
						}
						newCardQuery = newCardQuery.AddColors(cardColors...)
					}

					newCard, err := newCardQuery.Save(ctx)
					if err != nil {
						return err
					}

					for _, r := range v.Rulings {
						rulingDate, err := time.Parse("2006-01-02", r.Date)
						if err != nil {
							return err
						}
						_, err = db.Ruling.Create().SetText(r.Text).SetCard(newCard).SetDate(rulingDate).Save(ctx)
						if err != nil {
							return err
						}
					}
				} else {
					return err
				}
			}

			// TODO handle updating if it already exists
		}

		page += 1
		time.Sleep(time.Second * 2)
	}

	return nil
}
