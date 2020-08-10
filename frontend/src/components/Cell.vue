<template>
  <div class="cell" @click="$emit('click')">
    <img v-if="state === State.FLAGGED" src="@/assets/flag-24px.svg" />
    <img v-if="state === State.UNKNOWN" src="@/assets/help-24px.svg" />
    {{ state > 0 ? state : '' }}
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { State } from '../utils/state'

@Component
export default class Cell extends Vue {
  @Prop() private state!: State
  @Prop() private value!: number
  State = State

  get imageSrc() {
    switch (this.state) {
      case State.UNKNOWN:
        return 'help'
      case State.FLAGGED:
        return 'flag'
      default:
        return null
    }
  }
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
