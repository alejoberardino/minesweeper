import http, { BaseService } from './base'
import { Matrix } from '@/utils/state'

export const gameService = new (class GameService extends BaseService {
  path = '/games/'

  async create(dto: { rows: number; columns: number; mines: number }): Promise<string> {
    const response = (await http.post(this.path, dto)).data
    return response.id
  }

  async get(id: string) {
    return (await http.get(this.path + id)).data
  }

  private makeEmptyMatrix(rows: number, columns: number): Matrix {
    return Array.from(
      {
        length: rows,
      },
      () =>
        Array.from(
          {
            length: columns,
          },
          () => ({
            state: 0,
            value: 0,
          })
        )
    )
  }
})()
