<template>
  <div class="cell" />
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { State } from '../utils/state'

@Component
export default class Cell extends Vue {
  @Prop() private state!: State | number
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
div.cell {
  margin: 0;
  padding: 0;
  text-align: center;
  border: 1px solid black;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
