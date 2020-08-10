<template>
  <div class="home">
    <GameMatrix :id="game.id" :matrix="game.cells" @cell-clicked="cellClicked" />
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator'
import GameMatrix from '@/components/GameMatrix.vue'
import { Matrix, Game } from '@/utils/state'
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

  cellClicked({ x, y }: { x: number; y: number }) {
    console.log(`Cell clicked at (${x};${y})`)
  }
}
</script>
