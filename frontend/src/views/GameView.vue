<template>
  <div class="home">
    <GameMatrix :id="game.id" :matrix="game.cells" @cell-clicked="cellClicked" @cell-right-clicked="cellRightClicked" />
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator'
import GameMatrix from '@/components/GameMatrix.vue'
import { Matrix, Game, State } from '@/utils/state'
import { MIN_STATE } from '@/utils/state'
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

  async cellClicked({ x, y }: { x: number; y: number }, state = State.CLICKED) {
    console.log(`Cell clicked at (${x};${y})`)
    const currentState = this.game.cells[y][x].state
    if (currentState === State.CLICKED) {
      console.log('Cell already clicked')
      return
    }
    // Assume always left click
    const updatedCell = await gameService.click(this.game.id!, x, y, state)
    if (updatedCell.state !== state) {
      throw new Error(`Mismatched states, wanted ${state} got ${updatedCell.state}`)
    }

    this.$set(this.game.cells[y], x, updatedCell)
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
