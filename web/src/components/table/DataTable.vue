<template>
  <div class="table-container">
    <div class="table-controls">
      <div class="pag-controls">
        <i class="fa-solid fa-chevron-left" @click="decreasePage"></i>
        <span>{{ page }}/{{ total }}</span>
        <i class="fa-solid fa-chevron-right" @click="increasePage"></i>
      </div>
      <i class="reload fa-solid fa-rotate-right" @click="reloadData"></i>
    </div>
    <table class="table">
      <thead>
        <tr>
          <TableHeader
            v-for="header in headers"
            :key="header.field"
            :text="header.label"
            :width="header.width"
            :bold="header.bold"
          />
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in rows" :key="item.id">
          <TableCell
            v-for="header in headers"
            :key="header.field"
            :text="item[header.field]"
            :can-copy="header.copy"
          />
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script setup>
import { onMounted, ref } from 'vue'
import TableCell from './TableCell.vue'
import TableHeader from './TableHeader.vue'

const rows = ref([])
const page = ref(1)
const total = ref(0)

const props = defineProps({
  headers: {
    type: Array,
    required: true,
    default: () => []
  },
  fetchData: {
    type: Function,
    required: true
  },
  limit: {
    type: Number,
    required: false,
    default: 10
  }
})

async function increasePage() {
  let newPage = page.value + 1
  if (newPage > total.value) return
  const data = await props.fetchData(newPage, props.limit)
  console.log(data)
  page.value = newPage
  rows.value = data.items
  total.value = data.total_pages
}

async function decreasePage() {
  let newPage = page.value - 1
  if (newPage < 1) return
  const data = await props.fetchData(newPage, props.limit)
  page.value = newPage
  rows.value = data.items
  total.value = data.total_pages
}

async function reloadData() {
  const data = await props.fetchData(page.value, props.limit)
  console.log(data)
  rows.value = data.items
  total.value = data.total_pages
}

onMounted(async () => {
  reloadData()
})
</script>
<style scoped>
.table-container {
  overflow-x: scroll;
  width: 100%;
}
.table {
  table-layout: fixed;
  border: 1px solid var(--grey-01);
  border-collapse: collapse;
  width: fit-content;
}
.table-controls {
  display: flex;
  justify-content: center;
}
.pag-controls {
  margin-bottom: 0.5rem;
}
.pag-controls i,
.pag-controls span,
.reload {
  color: var(--white);
  font-size: 1.5rem;
  margin-left: 1ch;
}
.pag-controls i,
.reload {
  cursor: pointer;
}
.pag-controls i:hover,
.reload:hover {
  cursor: pointer;
  color: var(--primary);
}
</style>
