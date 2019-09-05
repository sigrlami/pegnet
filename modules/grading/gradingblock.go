package grading

import (
	"fmt"

	"github.com/pegnet/pegnet/modules/opr"
	log "github.com/sirupsen/logrus"
)

type GradingBlock struct {
	BlockHeight int32
	OPRVersion  uint8 // OPR version to use

	// graded indicates if the current set is graded.
	graded bool

	// OPRSet
	OPRs []*opr.OPR

	GradedOPRs []*opr.OPR

	// Grading variables
	PreviousWinners []string

	// Will output logs during the grading process to this logger if set
	Logger *log.Logger
}

func NewGradingBlock(height int32, version uint8) (*GradingBlock, error) {
	g := new(GradingBlock)
	g.BlockHeight = height
	switch version {
	case 1, 2:
		g.OPRVersion = version
	default:
		return nil, fmt.Errorf("%d is not a supported grading version", version)
	}
	// Silence all logs by default
	g.Logger = log.New()
	g.Logger.SetLevel(log.PanicLevel)

	return g, nil
}

func (g *GradingBlock) SetLogger(logger *log.Logger) {
	g.Logger = logger
}

func (g *GradingBlock) Height() int32 {
	return g.BlockHeight
}

func (g *GradingBlock) Version() uint8 {
	return g.OPRVersion
}

func (g *GradingBlock) AddOPR(entryhash []byte, extids [][]byte, content []byte) (added bool, err error) {
	// Unset the graded
	g.graded = false

	opr, err := opr.ParseOPR(entryhash, extids, content)
	if err != nil {
		// All errors are parse errors. We silence them here
		return false, nil
	}

	g.OPRs = append(g.OPRs, opr)
	return true, nil
}

func (g *GradingBlock) SetPreviousWinners(previousWinners []string) error {
	g.graded = false // Even if we error, we should unset this. A failed attempt will still reset

	switch g.Version() {
	case 1:
		if len(previousWinners) != 10 {
			return fmt.Errorf("exp 10 winners, found %d", len(previousWinners))
		}
	case 2:
		if !(len(previousWinners) == 10 || len(previousWinners) == 25) {
			return fmt.Errorf("exp 10 or 25 winners, found %d", len(previousWinners))
		}
	default:
		return fmt.Errorf("%d is not a supported grading version", g.Version())
	}

	// Verify they are all the right length
	for _, win := range previousWinners {
		if len(win) != 8 {
			return fmt.Errorf("exp winners to be of length 8, found %d", len(win))
		}
	}

	g.PreviousWinners = previousWinners

	return nil
}

func (g *GradingBlock) Grade() {
	g.GradedOPRs = g.GradeMinimum()
	g.graded = true
}

// WinnersShortHashes
//
// Requires: graded state
func (g *GradingBlock) WinnersShortHashes() ([]string, error) {
	winners, err := g.Winners()
	if err != nil {
		return nil, err
	}

	shorthashes := make([]string, g.winnerAmount(), g.winnerAmount())

	// A nil set is an empty set
	if winners == nil {
		return []string{}, nil
	}

	// This shouldn't ever happen...
	// TODO: Should this return an error?
	if len(winners) != len(shorthashes) {
		return shorthashes, nil
	}

	for i := range shorthashes {
		shorthashes[i] = winners[i].ShortEntryHash()
	}

	return shorthashes, nil
}

// Winners
//
// Requires: graded state
func (g *GradingBlock) Winners() (winners []*opr.OPR, err error) {
	return g.gradedUpTo(g.winnerAmount())
}

// Graded
//
// Requires: graded state
func (g *GradingBlock) Graded() (graded []*opr.OPR, err error) {
	return g.gradedUpTo(50)
}

func (g *GradingBlock) IsGraded() bool {
	return g.graded
}

func (g *GradingBlock) TotalOPRs() int {
	return len(g.OPRs)
}

func (g *GradingBlock) GetPreviousWinners() []string {
	return g.PreviousWinners
}

// gradedUpTo will return the set up to the maximum `pos`. So if `pos` is 50, but only 25 records exist,
// then graded[:25] is returned
func (g *GradingBlock) gradedUpTo(pos int) (graded []*opr.OPR, err error) {
	if !g.graded {
		return nil, fmt.Errorf("opr set is not graded yet")
	}

	if g.GradedOPRs == nil {
		return nil, nil
	}

	if len(g.GradedOPRs) < g.winnerAmount() {
		// This should never happen
		return nil, fmt.Errorf("something is wrong with the graded set, not enough winners")
	}

	// If
	if len(g.GradedOPRs) < pos {
		pos = len(g.GradedOPRs)
	}

	return g.GradedOPRs[:pos], nil
}

func (g *GradingBlock) winnerAmount() int {
	switch g.Version() {
	case 1:
		return 10
	case 2:
		return 25
	default:
		return 0
	}
}