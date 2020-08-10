<template>
  <div class="cell" @click="$emit('click')" @contextmenu.prevent="$emit('contextmenu')" :style="cellStyle">
    <img v-if="state === State.FLAGGED" src="@/assets/flag-24px.svg" />
    <img v-if="state === State.POSSIBLE" src="@/assets/help-24px.svg" />
    {{ state > 0 ? value : '' }}
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

  get cellStyle() {
    return {
      border: `3px ${this.state === State.CLICKED ? 'inset' : 'outset'}`,
    }
  }

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
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
