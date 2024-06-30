<template>
  <i class="fa-regular fa-circle-question tooltip">
    <span> {{ message }} </span>
  </i>
</template>

<script setup>
const props = defineProps({
  message: {
    required: true,
    type: String,
    default: 'This is a help message :)'
  },
  width: {
    required: false,
    type: String,
    default: '20ch'
  }
})
</script>

<style scoped>
.tooltip {
  color: var(--white);
  font-size: 1.2rem;
  transition: 0.5s;
  position: relative;
}
.tooltip:hover {
  color: var(--primary);
}

.tooltip span {
  --scale: 0;
  /* Center align vertically */
  position: absolute;
  top: 50%;
  width: v-bind(width);
  padding: 1ch;
  color: var(--background);
  background-color: var(--primary);
  transform: translateX(2ch) translateY(-50%) scale(var(--scale));
  border-radius: 3px;

  font-family: 'Outfit', sans-serif;
  font-size: 1rem;
  font-weight: light;

  transition: 0.2s;

  z-index: 100;
}

.tooltip:hover span {
  --scale: 1;
}
</style>
