import { BaseService } from './base'
import { Matrix } from '@/utils/state'

const ROWS = 10
const COLUMNS = 9

export const gameService = new (class GameService extends BaseService {
  path = '/games'

  create(): Matrix {
    // MOCK
    return Array.from(
      {
        length: ROWS,
      },
      () =>
        Array.from(
          {
            length: COLUMNS,
          },
          () => Math.floor(Math.random() * 12) - 4
        )
    )
  }
})()
