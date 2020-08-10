<template>
  <div class="container" :style="calculateMatrixStyle()">
    <template v-for="(rows, y) in matrix">
      <Cell
        v-for="(state, x) in rows"
        :key="`${x};${y}`"
        :style="calculateCellStyle(x, y)"
        :state="state"
        @click="$emit('cell-clicked', { x, y })"
      />
    </template>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import Cell from '@/components/Cell.vue'
import { State } from '@/utils/state'
import { gameService } from '@/services/game'

@Component({
  components: {
    Cell,
  },
})
export default class GameMatrix extends Vue {
  @Prop() private id!: number
  @Prop() private matrix!: State[][]

  calculateMatrixStyle() {
    return {
      'grid-template-rows': `repeat(${this.matrix.length}, 7vw)`,
      'grid-template-columns': `repeat(${this.matrix[0].length}, 7vw)`,
    }
  }

  calculateCellStyle(x: number, y: number) {
    return {
      'border-top-width': y === 0 ? '4px' : '1px',
      'border-left-width': x === 0 ? '4px' : '1px',
      'border-right-width': x === this.matrix[0].length - 1 ? '4px' : '1px',
      'border-bottom-width': y === this.matrix.length - 1 ? '4px' : '1px',
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
div.container {
  display: grid;
  justify-content: center;
  align-content: center;
  grid-gap: 0rem;
  list-style: none;
  margin: 0 0 2vw;
  padding: 0;
  font-size: calc(2vw + 10px);
}
</style>
