<template>
  <div class="home">
    <GameMatrix :matrix="matrix" />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import GameMatrix from '@/components/GameMatrix.vue'
import { Matrix } from '@/utils/state'

@Component({
  components: {
    GameMatrix,
  },
})
export default class Game extends Vue {
  private matrix: Matrix = [
    [0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0],
  ]
}
</script>
