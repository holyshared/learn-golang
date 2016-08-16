package gobang

func NewGameFacilitator(ctx *GameContext) *GameFacilitator {
  return &GameFacilitator {
    ctx: ctx,
  }
}

type GameFacilitator struct {
  ctx *GameContext
}

func (f *GameFacilitator) PlayerSelectCell(point *Point) (*Cell, error) {
  player := f.ctx.GamePlayer()
  cell, err := player.SelectBoardCell(point)

  return cell, err
}

func (f *GameFacilitator) PlayerPutStoneTo(cell *Cell) (GameProgressResult, error) {
  player := f.ctx.GamePlayer()
  err := player.PutStoneTo(cell)

  if err != nil {
    return Retry, err
  }
  result := f.ctx.CheckBoard()

  if result == Reached {
    return Win, nil
  } else if result == Filled {
    return Draw, nil
  } else {
    return Next, nil // FIXME GameResult
  }
}

func (f *GameFacilitator) ChangeToNextPlayer() {
  f.ctx.ChangeToNextPlayer()
}

func (f *GameFacilitator) NpcPlayerPutStone() (GameProgressResult, error) {
  player := f.ctx.NpcPlayer()
  cell := player.SelectTargetCell()
  err := player.PutStoneTo(cell)

  if err != nil {
    return Retry, err
  }
  result := f.ctx.CheckBoard()

  if result == Reached {
    return Lose, nil
  } else if result == Filled {
    return Draw, nil
  } else {
    return Next, nil // FIXME GameResult
  }
}
