export enum State {
  CLICKED = 1,
  UNKNOWN = 0,
  FLAGGED = -1,
  POSSIBLE = -2,
}

export interface Cell {
  state: State
  value: number
}
export type Matrix = Cell[][]

export interface Game {
  id: string | null
  cells: Matrix
  columns: number
  rows: number
  mines: number
  startedAt: Date | null
  value: string | null
}
