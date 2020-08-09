<template>
  <div class="home">
    <div>
      <label for="rows">Rows</label>
      <input type="number" name="rows" v-model="dto.rows" />
    </div>
    <div>
      <label for="columns">Columns</label>
      <input type="number" name="columns" v-model="dto.columns" />
    </div>
    <div>
      <label for="mines">Mines</label>
      <input type="number" name="mines" v-model="dto.mines" />
    </div>
    <button @click="start">Start game</button>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { Matrix } from '@/utils/state'
import { gameService } from '@/services/game'

@Component
export default class Home extends Vue {
  dto = {
    rows: 10,
    columns: 11,
    mines: 12,
  }

  async start() {
    const id = await gameService.create(this.dto)
    this.$router.push(`/games/${id}`)
  }
}
</script>

<style scoped>
div.home {
  width: 300px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-content: center;
}

div.home > div {
  display: flex;
  justify-content: space-between;
  margin-bottom: 1rem;
}
</style>
