<template>
  <div class="home">
    <GameMatrix :id="game.id" :matrix="game.cells" />
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
    console.log(this.game)
    this.game = await gameService.get(this.$route.params.id)
  }

  // beforeRouteUpdate(to: any, from: any, next: any) {
  //   console.log(this.$route.params)
  //   next()
  // }
}
</script>
