export enum SpecialStates {
  BLANK = 0,
  UNKNOWN = -1,
  FLAGGED = -2,
  BOMB = -3,
}
export type State = SpecialStates | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8
export type Matrix = State[][]
