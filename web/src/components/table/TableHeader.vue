<template>
  <th class="table-cell" :style="style">
    <span class="text" :style="textStyle">{{ text }}</span>
    <i v-if="canCopy" class="copy fa-regular fa-clone" @click="copyToClipboard"></i>
  </th>
</template>
<script setup>
import { computed } from 'vue'

const props = defineProps({
  text: {
    type: String,
    default: ''
  },
  canCopy: {
    type: Boolean,
    default: false
  },
  width: {
    type: String,
    default: '10ch'
  },
  bold: {
    type: Boolean,
    default: false
  }
})

const style = computed(() => {
  let currentWidth = props.width
  return `width: ${currentWidth}; `
})
const textStyle = computed(() => {
  let currentBold = props.bold == true ? 'bolder' : 'normal'
  return `font-weight: ${currentBold};`
})

function copyToClipboard() {
  navigator.clipboard.writeText(props.text)
}
</script>
<style scoped>
.table-cell {
  padding-top: 1rem;
  padding-bottom: 1rem;
}
.copy {
  width: 3ch;
  color: var(--grey-03);
  transition: color ease-in 0.2s;
}
.copy:hover {
  color: var(--sky-blue);
}
</style>
