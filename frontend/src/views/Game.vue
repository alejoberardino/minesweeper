<template>
  <div class="home">
    <GameMatrix :id="id" />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import GameMatrix from '@/components/GameMatrix.vue'
import { Matrix } from '@/utils/state'
import { gameService } from '@/services/game'

@Component({
  components: {
    GameMatrix,
  },
})
export default class Game extends Vue {
  id = null
}
</script>
