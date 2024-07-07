<template>
  <input class="input" :type="type" @input="update" :placeholder="placeholder" />
</template>
<script setup>
const emit = defineEmits(['input'])

const model = defineModel()

function update(event) {
  model.value = event.target.value
}

const props = defineProps({
  type: {
    type: String,
    default: 'text'
  },
  placeholder: {
    type: String,
    default: ''
  }
})
</script>
<style scoped>
.input {
  border-color: #fdfcfc1a;
  background-color: #1d191673;
  color: white;

  border-radius: 0.75rem;
  border-width: 1px;
  border-style: solid;

  padding: 0.375rem 0.725rem;
  margin-top: 0.5ch;
  appearance: none;
  line-height: 1.5rem;

  font-size: 1rem;

  outline: 2px solid transparent;

  outline-offset: 2px;
}
.input:hover {
  border-color: var(--grey-03);
}
.input:focus {
  border-color: var(--primary);
}
</style>
