package gobang

func NewCellReachedMatcher(stone Stone, count int) CellReachedMatcher {
  selectors := make([]ReachedMatcher, 0)
  selectors = append(selectors, NewVerticalCellMatcher(stone, count))
  selectors = append(selectors, NewHorizontalCellMatcher(stone, count))
  selectors = append(selectors, NewTopLeftDiagonalCellMatcher(stone, count))
  selectors = append(selectors, NewTopRightDiagonalCellMatcher(stone, count))

  return CellReachedMatcher { selectors: selectors }
}

type CellReachedMatcher struct {
  selectors []ReachedMatcher
}

func (s *CellReachedMatcher) Matches(board *Board) *MatchedResult {
  result := &MatchedResult {}
  ch := make(chan *MatchedResult, len(s.selectors))
  wn := len(s.selectors)

  defer close(ch)

  for _, selector := range s.selectors {
    go func(selector ReachedMatcher, board *Board) {
      ch <- selector.Matches(board)
    }(selector, board)
  }

  for res := range ch {
    wn--;
    result.Merge(res)
    if wn <= 0 {
      break;
    }
  }

  return result
}
