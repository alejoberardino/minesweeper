<template>
  <div class="home">
    <div>
      <button @click="$router.push('/')">Go home</button>
      <button @click="restart">Restart</button>
    </div>
    <GameMatrix :id="game.id" :matrix="game.cells" @cell-clicked="cellClicked" @cell-right-clicked="cellRightClicked" />
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator'
import { Watch } from 'vue-property-decorator'
import GameMatrix from '@/components/GameMatrix.vue'
import { Matrix, Game, State } from '@/utils/state'
import { gameService } from '@/services/game'

@Component({
  components: {
    GameMatrix,
  },
})
export default class GameView extends Vue {
  game: Game = {
    id: null,
    cells: [[]],
    columns: 0,
    rows: 0,
    mines: 0,
    startedAt: null,
    value: null,
  }

  async created() {
    this.game = await gameService.get(this.$route.params.id)
  }

  @Watch('$route.params')
  async onParamsChange(from: any, to: any) {
    this.game = await gameService.get(to.id)
  }

  async restart() {
    const id = await gameService.create({
      rows: this.game.rows,
      columns: this.game.columns,
      mines: this.game.mines,
    })
    this.$router.push(`/games/${id}`)
  }

  async cellClicked({ x, y }: { x: number; y: number }, state = State.CLICKED) {
    console.log(`Cell clicked at (${x};${y})`)
    const currentState = this.game.cells[y][x].state
    if (currentState === State.CLICKED) {
      console.log('Cell already clicked')
      return
    }
    // Assume always left click
    const updatedGame = await gameService.click(this.game.id!, x, y, state)
    if (updatedGame.cells[y][x].state !== state) {
      throw new Error(`Mismatched states, wanted ${state} got ${updatedGame.cells[y][x].state}`)
    }

    this.game = updatedGame
  }

  cellRightClicked({ x, y }: { x: number; y: number }) {
    const currentState = this.game.cells[y][x].state
    const nextState = this.getNextState(currentState)
    console.log(`Current state = ${currentState}, setting to ${nextState}`)
    return this.cellClicked({ x, y }, nextState)
  }

  getNextState(state: State): State {
    return state - 1 < State.FLAGGED ? State.UNKNOWN : state - 1
  }
}
</script>
