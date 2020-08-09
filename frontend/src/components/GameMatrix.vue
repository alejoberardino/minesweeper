<template>
  <div class="container" :style="calculateMatrixStyle()">
    <template v-for="(rows, i) in matrix">
      <Cell v-for="(state, j) in rows" :key="`${i};${j}`" :style="calculateCellStyle(i, j)" />
    </template>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import Cell from '@/components/Cell.vue'
import { State } from '@/utils/state'

@Component({
  components: {
    Cell,
  },
})
export default class GameMatrix extends Vue {
  @Prop() private matrix!: State[][]

  calculateMatrixStyle() {
    return {
      'grid-template-rows': `repeat(${this.matrix.length}, 5vw)`,
      'grid-template-columns': `repeat(${this.matrix[0].length}, 5vw)`,
    }
  }

  calculateCellStyle(i: number, j: number) {
    return {
      'border-top-width': i === 0 ? '4px' : '1px',
      'border-left-width': j === 0 ? '4px' : '1px',
      'border-right-width': j === this.matrix[0].length - 1 ? '4px' : '1px',
      'border-bottom-width': i === this.matrix.length - 1 ? '4px' : '1px',
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
