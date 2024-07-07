<template>
  <div class="table-container">
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
          <TableCell v-for="header in headers" 
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
import TableHeader from './TableHeader.vue';

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
    default: 20
  }
})

onMounted(async () => {
  const data = await props.fetchData(page.value, props.limit)
  rows.value = data.items
  total.value = data.total_pages
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
</style>
