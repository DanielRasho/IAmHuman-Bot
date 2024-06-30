<template>
  <button class="button" @click="handleClick">
    <i v-if="preIcon" :class="preIcon"></i>
    {{ text }}
    <i v-if="postIcon" :class="postIcon"></i>
  </button>
</template>
<script setup>
const emit = defineEmits(['clicked'])

const props = defineProps({
  text: {
    type: String,
    required: true,
    default: 'Button'
  },
  preIcon: {
    type: String,
    required: false,
    default: undefined
  },
  postIcon: {
    type: String,
    required: false,
    default: undefined
  }
})

function handleClick() {
  emit('clicked')
}
</script>
<style scoped>
.button {
  background-color: var(--primary);
  color: var(--dark-blue);
  font-size: 1.2rem;
  padding: 0.5em 0.7em;
  border-radius: 0.4rem;
  transition: background-color ease-in-out 0.3s;
}
.button i {
  font-family: 'Font Awesome 6 Free';
  color: var(--dark-blue);
}
.button:hover {
  cursor: pointer;
  background-color: var(--primary-light);
}
</style>
